package proxmox

import (
	"os"
	"testing"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
)

func TestCreateVmFromTemplate(t *testing.T) {

	f, err := os.CreateTemp("", "id_rsa.pub")
	f.Write([]byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCwACjkWg++UdGxMM49B6Q91WNIhD5veRapk8oXnnjn01SZKvPryj/UqITYemlJ82mSDQwvLf8Fquy/r05Jx0au/JYl7PcTdDVD9wuDVAcg73JDFUZH8esteGK+ta6xjxMJ0XdSGUEN3Wp19NOI7Pq6N/UQl30Y1SsPoCjApajNK+3l6Bs45TDtnsJk1cm5DqN4vhIhzPMSC2cS/OePiZm035baOt1vqRGB/+Hm4GbwyYAuKDfXNcpDREfdJ7iOrSHhUZHRDTHNkoDKhv5ZAnf9/EHBEijRIlFk/lXfjmGesD8q/ULvf+lk7OhPqtpye/p/vA7zAxLo1FmldOw99CLqF5AkojcMapGOqvOOtWF59YPqPi3P/VaI321y03lNO8jR5YxW6ppShfeSV0teUQsqkWbnwHhwOPcGC4xPvA8jE1AgdVxDqMBOiFvZviKvcMS4V0rdw43OJibFyAqvqLWNAIUD9htBsjpqWA8i2/bcfD//VDO1ZtwuT0eHOOGsFaE= daniel@proxmox"))
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(f.Name())

	credentials := entity.NewProxmoxCredential(os.Getenv("PROXMOX_ENDPOINT"), os.Getenv("PROXMOX_USERNAME"), os.Getenv("PROXMOX_PASSWORD"))

	ticketUseCase := NewGetTicketUseCase(credentials)
	ticket := ticketUseCase.Execute()

	ciConfig := CiConfig{
		CiUser:      "someuser",
		CiPassword:  "somekey",
		SshKeysFile: "",
		Ipconfig:    "",
	}

	vmClone := VmClone{
		TmplId:  "900",
		NewVmId: "115",
		Name:    "ansible-1",
		Node:    "proxmox",
		Pool:    "ansible",
	}

	createVmUseCase := NewCreateVmUseCase(ticket, ciConfig, vmClone)
	createVmUseCase.Execute()

}
