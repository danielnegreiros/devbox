package proxmox

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

type CreateVmPoolSnapUseCase struct {
	ticket *entity.TicketData
	vmids  []string
}

func NewCreateVmPoolSnapUseCase(ticket *entity.TicketData, vmids []string) *CreateVmPoolSnapUseCase {
	return &CreateVmPoolSnapUseCase{
		ticket: ticket,
		vmids:  vmids,
	}
}

func (u *CreateVmPoolSnapUseCase) Execute(node string) {

	currentTime := time.Now().Format("2006-01-02__15_04_05")
	snapName := "daily_" + currentTime

	for _, vmid := range u.vmids {

		httpReq := rest.HttpRequest{
			Timeout:       10,
			EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%s/snapshot", u.ticket.Host, node, vmid),
			Method:        "POST",
			AcceptedCodes: []int{200},
			Data:          &struct{}{},
			Body:          []byte(fmt.Sprintf(`{"snapname": "%s"}`, snapName)),
			Cookie: &http.Cookie{
				Name:  "PVEAuthCookie",
				Value: u.ticket.Ticket,
			},
			Header: map[string]string{
				"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
				"Content-Type":        "application/json",
			},
		}

		content, code := httpReq.Execute()
		if code != http.StatusOK {
			log.Panic(content)
		}
	}

}
