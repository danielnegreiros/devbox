package proxmox

import (
	"log"
	"os"
	"testing"

	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
)

func TestGetTicket(t *testing.T) {
	credentials := entity.NewProxmoxCredential(os.Getenv("PROXMOX_ENDPOINT"), os.Getenv("PROXMOX_USERNAME"), os.Getenv("PROXMOX_PASSWORD"))

	ticketUseCase := NewGetTicketUseCase(credentials)
	result := ticketUseCase.Execute()

	if len(result.Ticket) < 30 {
		t.Error("Invalid Ticket")
	}

	log.Println(result)
}
