package parser

import (
	"fmt"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/main_cmd"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/utils"
)

var availGroupCmd = main_cmd.GetAvailGroupCmds()

func init() {
	availGroupCmd = main_cmd.GetAvailGroupCmds()
}

// GetCommandFn parses the main command.
func GetMainCommandParserFn(args []string) model.MainCommandParserFn {

	usageText := `
Usage: crosscmd [OPTIONS] COMMAND [arg...]

Universal Command-line Interface to Manage Multiple Environments

Options:
  -h, --help    Show this help message and exit
  -v, --version Show version information

Commands:
`

	var textComplement strings.Builder
	for _, sub := range availGroupCmd {
		textComplement.WriteString(fmt.Sprintf("  %s\t%s\n", sub.Name, sub.Description))
	}

	usageText += textComplement.String()
	usageText += "\nRun 'crosscmd COMMAND --help' for more information on a command.\n"

	if len(args) < 2 {
		utils.Exit(usageText)
	}

	cmdInList, fn := isCmdinListCmds(args[1], model.AvailGroupCmd)
	if !cmdInList {
		utils.Exit(usageText)
	}
	return fn
}

func isCmdinListCmds(pattern string, optionsList []model.MainCmd) (bool, model.MainCommandParserFn) {
	for _, item := range optionsList {
		cleanName := strings.ReplaceAll(item.Name, "\t", "")
		if cleanName == pattern {
			return true, item.GroupParserFn
		}
	}
	return false, nil
}
