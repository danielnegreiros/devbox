package proxmox

import (
	"strconv"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/proxmox"
)

func CleanUpSnapshotHandler(argsMap map[string]string) {

	credentials := entity.NewProxmoxCredential(argsMap["endpoint"], argsMap["user"], argsMap["password"])

	ticketUseCase := proxmox.NewGetTicketUseCase(credentials)
	ticket := ticketUseCase.Execute()

	listVmUSeCase := proxmox.NewListVMsUseCase(ticket)
	snapListUseCase := proxmox.NewSnapListUseCase(ticket)

	days, _ := strconv.Atoi(argsMap["days"])
	snapCleanUpUseCase := proxmox.NewSnapCleanUpUseCase(ticket, days, argsMap["include"], argsMap["exclude"])
	executeSnapshotCleanUpCommand(argsMap, ticketUseCase, listVmUSeCase, snapListUseCase, snapCleanUpUseCase)
}

func executeSnapshotCleanUpCommand(argsMap map[string]string,
	ticketUseCase *proxmox.GetTicketUseCase, listVmUSeCase *proxmox.ListVMsUseCase,
	snapListUseCase *proxmox.SnapListUseCase, snapCleanUpUseCase *proxmox.SnapCleanUpUseCase) {

	vmRresult := listVmUSeCase.Execute(argsMap["node"])
	snapshots := getAllSnapsByVm(*vmRresult, argsMap["node"], snapListUseCase)

	snapCleanUpUseCase.Execute(argsMap["node"], snapshots)
}

func getAllSnapsByVm(vmRresult proxmox.VMListResponse, node string, snapListUseCase *proxmox.SnapListUseCase) []proxmox.SnapShotData {

	snapshots := []proxmox.SnapShotData{}
	for _, vm := range vmRresult.VmData {

		res := snapListUseCase.Execute(node, vm.Vmid)
		for _, snap := range res.VmData {
			snap.Vmid = vm.Vmid
			snap.VmName = vm.Name
			snapshots = append(snapshots, snap)
		}
	}

	return snapshots
}
