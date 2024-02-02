package unix

import (
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/unix"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/unix/conf"
	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/adapters"
	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/ports"
)

func HardenSshHandler(argsMap map[string]string) {

	infra := getOsAdapter(argsMap)

	hardening := conf.GetDefaultHarden()
	unixHardenUC := unix.NewHardenSSHUseCase(infra, hardening)
	unixHardenUC.Execute()
}

func getOsAdapter(argsMap map[string]string) ports.InfraRepo {

	if argsMap["host"] == "local" || argsMap["host"] == "localhost" {
		return adapters.NewLocalAdapter()
	} else {
		return adapters.NewLinuxAdapter("", 0, "", "")
	}
}
