package adapters

import (
	"log"
	"os/exec"

	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/ports"
)

type LocalAdapter struct{}

func NewLocalAdapter() *LocalAdapter {
	return &LocalAdapter{}
}

func (a *LocalAdapter) ExecuteCommand(command string) ports.CommandOutputDTO {

	log.Printf("Starting command: %s", command)
	// re := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)
	// args := re.FindAllString(command, -1)

	output, err := exec.Command("bash", "-c", command).CombinedOutput()

	log.Printf("Finished command: %s", command)
	return ports.NewCommandOutputDTO(string(output), "", err == nil, err)
}

func (a *LocalAdapter) Close() {}
