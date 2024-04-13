package proxmox

import (
	"fmt"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/proxmox"
)

func CreateVmHandler(argsMap map[string]string) {

	credentials := entity.NewProxmoxCredential(argsMap["endpoint"], argsMap["user"], argsMap["password"])

	ticketUseCase := proxmox.NewGetTicketUseCase(credentials)
	ticket := ticketUseCase.Execute()

	var ipconfig string
	if argsMap["vm_ip"] != "" && argsMap["gateway"] != "" && argsMap["vm_netmask"] != "" {
		ipconfig = fmt.Sprintf("ip=%s/%s,gw=%s", argsMap["vm_ip"], argsMap["vm_netmask"], argsMap["gateway"])
	} else {
		ipconfig = ""
	}
	ciConfig := proxmox.CiConfig{
		CiUser:      argsMap["vm_user"],
		CiPassword:  argsMap["vm_pass"],
		SshKeysFile: argsMap["vm_pub_keys"],
		Ipconfig:    ipconfig,
		Cores:  argsMap["cores"],
		Sockets: argsMap["sockets"],
	}

	vmClone := proxmox.VmClone{
		TmplId:  argsMap["vm_template_id"],
		NewVmId: argsMap["vm_id"],
		Name:    argsMap["vm_name"],
		Node:    argsMap["node"],
		Pool:    argsMap["pool"],
	}

	createVmUseCase := proxmox.NewCreateVmUseCase(ticket, ciConfig, vmClone)
	createVmUseCase.Execute()

}
