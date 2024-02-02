package proxmox

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
)

func TestListSnaps(t *testing.T) {

	credentials := entity.NewProxmoxCredential(os.Getenv("PROXMOX_ENDPOINT"), os.Getenv("PROXMOX_USERNAME"), os.Getenv("PROXMOX_PASSWORD"))

	ticketUseCase := NewGetTicketUseCase(credentials)
	ticket := ticketUseCase.Execute()

	vmUseCase := NewListVMsUseCase(ticket)
	vmRresult := vmUseCase.Execute(os.Getenv("PROXMOX_NODE"))

	snapUseCase := NewSnapListUseCase(ticket)
	snapshots := []SnapShotData{}

	for _, vm := range vmRresult.VmData {
		res := snapUseCase.Execute(os.Getenv("PROXMOX_NODE"), vm.Vmid)
		for _, snap := range res.VmData {
			snap.Vmid = vm.Vmid
			snapshots = append(snapshots, snap)
		}
	}
	log.Println(snapshots)
}

func TestDate(t *testing.T) {

	snapTime := time.Unix(1702157648, 0)
	now := time.Now()

	delta := now.Sub(snapTime)
	log.Println(int(delta.Hours() / 24))

}
