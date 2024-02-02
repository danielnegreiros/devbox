package proxmox

import (
	"fmt"
	"log"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/errs"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/conf"
	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/ports"
)

type cloudInitTemplateUseCase struct {
	infra   ports.InfraRepo
	image   string
	id      string
	name    string
	storage string
	net     string
}

func NewcloudInitTemplate(infra ports.InfraRepo, image string, id string, name string, storage string, net string) *cloudInitTemplateUseCase {
	return &cloudInitTemplateUseCase{
		infra:   infra,
		image:   image,
		id:      id,
		storage: storage,
		name:    name,
		net:     net,
	}
}

func (u cloudInitTemplateUseCase) Execute() {
	result := u.infra.ExecuteCommand("mkdir -pv /tmp/cloudinit")
	failIfResultErr(result)

	image, err := conf.GetImageUrl(u.image)
	errs.PanicIfErr(err)

	result = u.infra.ExecuteCommand("wget -P /tmp/cloudinit/ " + image)
	failIfResultErr(result)

	result = u.infra.ExecuteCommand("qm stop " + u.id)
	failIfResultErr(result)

	result = u.infra.ExecuteCommand("qm destroy " + u.id)
	failIfResultErr(result)

	makeCmd := fmt.Sprintf("qm create %s --memory 2048 --net0 virtio,bridge=%s --scsihw virtio-scsi-pci", u.id, u.net)
	result = u.infra.ExecuteCommand(makeCmd)
	failIfResultErr(result)

	makeCmd = fmt.Sprintf("qm set %s --scsi0 %s:0,import-from=/tmp/cloudinit/%s", u.id, u.storage, getImageName(image))
	result = u.infra.ExecuteCommand(makeCmd)
	failIfResultErr(result)

	makeCmd = fmt.Sprintf("qm set %s --ide2 %v:cloudinit", u.id, u.storage)
	result = u.infra.ExecuteCommand(makeCmd)
	failIfResultErr(result)

	makeCmd = fmt.Sprintf("qm set %s --boot order=scsi0", u.id)
	result = u.infra.ExecuteCommand(makeCmd)
	failIfResultErr(result)

	makeCmd = fmt.Sprintf("qm set %s --name %s", u.id, u.name)
	result = u.infra.ExecuteCommand(makeCmd)
	failIfResultErr(result)

	makeCmd = fmt.Sprintf("qm template %s", u.id)
	result = u.infra.ExecuteCommand(makeCmd)
	failIfResultErr(result)

	// qm resize 100 virtio0 +5G
	u.close()

}

func getImageName(imageUrl string) string {
	splitImage := strings.Split(imageUrl, "/")
	return splitImage[len(splitImage)-1]
}

func (u cloudInitTemplateUseCase) close() {
	u.infra.Close()
}

func failIfResultErr(output ports.CommandOutputDTO) {

	if strings.Contains(fmt.Sprint(output.Error), "unable to find configuration file for") {
		return
	}

	if strings.Contains(fmt.Sprint(output.Error), "does not exist") {
		return
	}

	if !output.Success {
		log.Fatal(output.Error)
	}
}
