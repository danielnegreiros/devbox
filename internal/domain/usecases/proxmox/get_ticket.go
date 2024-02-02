package proxmox

import (
	"fmt"

	httpclient "github.com/danielnegreiros/go-proxmox-cli/internal/app/http_client"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

type Response struct {
	TicketData *entity.TicketData `json:"data"`
}

type GetTicketUseCase struct {
	credentials *entity.ProxmoxCredential
}

func NewGetTicketUseCase(credentials *entity.ProxmoxCredential) *GetTicketUseCase {
	return &GetTicketUseCase{
		credentials: credentials,
	}
}

func (u *GetTicketUseCase) Execute() *entity.TicketData {

	httpReq := rest.HttpRequest{
		Timeout: 10,
		EndPoint: fmt.Sprintf("%s/api2/json/access/ticket", u.credentials.Host),
		Method: "POST",
		Body: httpclient.ComposeCredentials(u.credentials.User, u.credentials.Pass),
		AcceptedCodes: []int{200},
		Data: &Response{},
	}

	content, _ := httpReq.Execute()

	data := content.(*Response)
	data.TicketData.Host = u.credentials.Host

	return data.TicketData

}
