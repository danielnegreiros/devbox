package unix

import (
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/usecases/unix/conf"
	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/ports"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/munix"
)

type hardenSshUseCase struct {
	infra   ports.InfraRepo
	harden conf.HardenConf
}

func NewHardenSSHUseCase(infra ports.InfraRepo, harden conf.HardenConf) *hardenSshUseCase {
	return &hardenSshUseCase{
		infra: infra,
		harden: harden,
	}
}

func (u hardenSshUseCase) Execute() {

	for _, file := range u.harden.Files {
		for _, change := range file.Changes {
			lineInFile := munix.NewLineInFile(
				file.Path,
				change.Regex,
				change.Line,
				change.State,
				change.Owner,
				change.Group,
				change.Mode,
				change.Validate,
				change.ShouldCreate,
			)

			lineInFile.Execute(u.infra)

		}
	}
}
