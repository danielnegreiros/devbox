package proxmox

import (
	"log"
	"os"
	"testing"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
)

func TestListVms(t *testing.T) {

	credentials := entity.NewProxmoxCredential(os.Getenv("PROXMOX_ENDPOINT"), os.Getenv("PROXMOX_USERNAME"), os.Getenv("PROXMOX_PASSWORD"))

	ticketUseCase := NewGetTicketUseCase(credentials)
	ticket := ticketUseCase.Execute()

	vmUseCase := NewListVMsUseCase(ticket)
	vmRresult := vmUseCase.Execute(os.Getenv("PROXMOX_NODE"))

	log.Println(vmRresult)

}
