package proxmox

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

type SnapCleanUpUseCase struct {
	ticket  *entity.TicketData
	days    int
	include string
	exclude string
}

func NewSnapCleanUpUseCase(ticket *entity.TicketData, days int, include string, exclude string) *SnapCleanUpUseCase {
	return &SnapCleanUpUseCase{
		ticket:  ticket,
		days:    days,
		include: include,
		exclude: exclude,
	}
}

func (u *SnapCleanUpUseCase) Execute(node string, snaps []SnapShotData) {
	log.Println("-------- START --------")

	resSnaps := getSnapsOlderThan(snaps, u.days)
	resSnaps = getIncludedOnlySnaps(resSnaps, u.include)
	resSnaps = getWithExcludedSnaps(resSnaps, u.exclude)

	for _, snap := range resSnaps {
		log.Printf("Starting deleting snapshot: %s from VM: %s\n", snap.Name, snap.VmName)
		u.deleteSnap(node, snap.Vmid, snap.Name)
		log.Printf("Finished deleting snapshot: %s from VM: %s\n", snap.Name, snap.VmName)
	}

	if len(resSnaps) == 0 {
		log.Println("No snapshots were found under requested filters")
	}

	log.Println("-------- END --------")
}

func getIncludedOnlySnaps(snaps []SnapShotData, include string) []SnapShotData {

	resSnaps := []SnapShotData{}
	regex := regexp.MustCompilePOSIX(include)

	for _, snap := range snaps {
		if regex.MatchString(snap.Name) {
			resSnaps = append(resSnaps, snap)
		}
	}

	return resSnaps

}

func getWithExcludedSnaps(snaps []SnapShotData, exclude string) []SnapShotData {

	resSnaps := []SnapShotData{}
	regex := regexp.MustCompilePOSIX(exclude)

	for _, snap := range snaps {
		if !regex.MatchString(snap.Name) {
			resSnaps = append(resSnaps, snap)
		}
	}

	return resSnaps

}

func getSnapsOlderThan(snaps []SnapShotData, days int) []SnapShotData {

	resSnaps := []SnapShotData{}
	now := time.Now()

	for _, snap := range snaps {
		snapTime := time.Unix(int64(snap.SnapTime), 0)
		delta := now.Sub(snapTime)

		if int(delta.Hours()/24) > days && snap.SnapTime != 0 {
			resSnaps = append(resSnaps, snap)
		}
	}

	return resSnaps
}

func (u *SnapCleanUpUseCase) deleteSnap(node string, vmid int, snapid string) {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%d/snapshot/%s", u.ticket.Host, node, vmid, snapid),
		Method:        "DELETE",
		AcceptedCodes: []int{200},
		Data:          &struct{}{},
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
		Header: map[string]string{
			"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
		},
	}

	httpReq.Execute()

}
