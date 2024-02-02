package proxmox

import (
	"fmt"
	"net/http"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

type VMData struct {
	Name string `json:"name"`
	Vmid int    `json:"vmid"`
}

type VMListResponse struct {
	VmData []VMData `json:"data"`
}

type ListVMsUseCase struct {
	ticket *entity.TicketData
}

func NewListVMsUseCase(ticket *entity.TicketData) *ListVMsUseCase {
	return &ListVMsUseCase{
		ticket: ticket,
	}
}

func (u *ListVMsUseCase) Execute(node string) *VMListResponse {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu?full=0", u.ticket.Host, node),
		Method:        "GET",
		AcceptedCodes: []int{200},
		Data:          &VMListResponse{},
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
	}

	content, _ := httpReq.Execute()
	data := content.(*VMListResponse)

	return data

}
