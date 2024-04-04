package proxmox

import (
	"fmt"
	"net/http"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

type Member struct {
	Name      string `json:"name"`
	VMID      int    `json:"vmid"`
    // MaxDisk   int    `json:"maxdisk"`
    // NetIn     int    `json:"netin"`
    // Disk      int    `json:"disk"`
    // DiskWrite int    `json:"diskwrite"`
    // Uptime    int    `json:"uptime"`
    // Template  int    `json:"template"`
    // NetOut    int    `json:"netout"`
    // MaxMem    int    `json:"maxmem"`
    // ID        string `json:"id"`
    // Mem       int    `json:"mem"`
    // CPU       float64 `json:"cpu"`
    // Status    string `json:"status"`
    // MaxCPU    int    `json:"maxcpu"`
    // Node      string `json:"node"`
    // DiskRead  int    `json:"diskread"`
    // Type      string `json:"type"`
}

type Data struct {
    Members []Member `json:"members"`
    PoolID  string   `json:"poolid"`
}

type ResponsePool struct {
    Data Data `json:"data"`
}


type ListVMsFromPoolUseCase struct {
	ticket *entity.TicketData
}

func NewListVMsFromPoolUseCase(ticket *entity.TicketData) *ListVMsFromPoolUseCase {
	return &ListVMsFromPoolUseCase{
		ticket: ticket,
	}
}

func (u *ListVMsFromPoolUseCase) Execute(pool string) *ResponsePool {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/pools/%s", u.ticket.Host, pool),
		Method:        "GET",
		AcceptedCodes: []int{200},
		Data:          &ResponsePool{},
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
	}

	content, _ := httpReq.Execute()
	data := content.(*ResponsePool)

	return data

}
