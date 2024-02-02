package proxmox

import (
	"fmt"
	"net/http"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

type SnapShotData struct {
	Name     string `json:"name"`
	SnapTime int    `json:"snaptime"`
	Vmid     int
	VmName   string
}

type SnapListResponse struct {
	VmData []SnapShotData `json:"data"`
}

type SnapListUseCase struct {
	ticket *entity.TicketData
}

func NewSnapListUseCase(ticket *entity.TicketData) *SnapListUseCase {
	return &SnapListUseCase{
		ticket: ticket,
	}
}

func (u *SnapListUseCase) Execute(node string, vmid int) *SnapListResponse {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%d/snapshot", u.ticket.Host, node, vmid),
		Method:        "GET",
		AcceptedCodes: []int{200},
		Data:          &SnapListResponse{},
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
	}

	content, _ := httpReq.Execute()
	data := content.(*SnapListResponse)

	return data
}
