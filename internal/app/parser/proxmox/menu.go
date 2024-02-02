package proxmox

import (
	"fmt"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/utils"
)

// ParseCommand parses the main command.
func ParseProxmoxCommand(args []string) (model.SubCmdParserFn, model.UseCaseCliHandlerFn) {

	usageText := `
Usage: crosscmd [OPTIONS] COMMAND [arg...]

Proxmox Command Line Interface. Simplify your operations.

Options:
  -h, --help    Show this help message and exit
  -v, --version Show version information

Commands:
`

	parent := "proxmox"

	var textComplement strings.Builder
	childSubCmds := []model.SubCmd{}

	for _, sub := range model.AvailSubCmds {
		if sub.Parent == parent {
			textComplement.WriteString(fmt.Sprintf("  %s\t%s\n", sub.Name, sub.Description))
			childSubCmds = append(childSubCmds, sub)
		}
	}

	usageText += textComplement.String()
	usageText += "\nRun 'devtoolbox COMMAND --help' for more information on a command.\n"

	if len(args) < 1 {
		utils.Exit(usageText)
	}

	cmdInList, subCmcMocdl := utils.IsOptioninChildSubCmds(args[0], childSubCmds)
	if !cmdInList {
		utils.Exit(usageText)
	}

	return subCmcMocdl.ParseFunc, subCmcMocdl.ExecFunc

}
