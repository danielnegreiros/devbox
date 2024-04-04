package proxmox

import (
	"strconv"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/proxmox"
)

func CreateSnapshotHandler(argsMap map[string]string) {

	credentials := entity.NewProxmoxCredential(argsMap["endpoint"], argsMap["user"], argsMap["password"])

	ticketUseCase := proxmox.NewGetTicketUseCase(credentials)
	ticket := ticketUseCase.Execute()

	listVmUSeCase := proxmox.NewListVMsFromPoolUseCase(ticket)
	vms := listVmUSeCase.Execute(argsMap["pool"])
	vmsIdList := []string{}

	for _, vm := range vms.Data.Members {
		vmsIdList = append(vmsIdList, strconv.Itoa(vm.VMID))
	}

	snapPoolUseCase := proxmox.NewCreateVmPoolSnapUseCase(ticket, vmsIdList)
	snapPoolUseCase.Execute(argsMap["node"])

	// days, _ := strconv.Atoi(argsMap["days"])
	// snapCleanUpUseCase := proxmox.NewSnapCleanUpUseCase(ticket, days, argsMap["include"], argsMap["exclude"])
	// executeSnapshotCleanUpCommand(argsMap, ticketUseCase, listVmUSeCase, snapListUseCase, snapCleanUpUseCase)
}
