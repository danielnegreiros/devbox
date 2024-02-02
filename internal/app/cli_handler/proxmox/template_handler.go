package proxmox

import (
	"strconv"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/errs"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/proxmox"
	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/adapters"
)

func TemplateHandler(argsMap map[string]string) {
	port, err := strconv.Atoi(argsMap["port"])
	errs.PanicIfErr(err)

	infra := adapters.NewLinuxAdapter(argsMap["host"], port, argsMap["user"], argsMap["password"])
	proxmoxUseCase := proxmox.NewcloudInitTemplate(infra,
		argsMap["image"], argsMap["id"], argsMap["name"], argsMap["storage"], argsMap["net"])

	proxmoxUseCase.Execute()
}
