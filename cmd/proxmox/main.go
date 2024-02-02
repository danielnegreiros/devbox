package main

import (
	"os"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/utils"
)

func main() {
	filteredArgs, remainingArgs := utils.GetArgs(os.Args, 0, 2)
	subCmdParserFn, subCmdCliHandlerFn := parseMainCommand(filteredArgs, remainingArgs)
	executeSubCommand(subCmdParserFn, subCmdCliHandlerFn, remainingArgs)
}

// parseMainCommand parses the first argument, that is the main command.
// and if the main command is valid return subcommand parser and subcommand cli handler executer
func parseMainCommand(filteredArgs []string, remainingArgs []string) (model.SubCmdParserFn, model.UseCaseCliHandlerFn) {
	groupParserFn := parser.GetMainCommandParserFn(filteredArgs)
	return groupParserFn(remainingArgs)
}

// executeSubCommand parses the subcommand according to remaining args. Excluding its own subcommand.
// then executes subcommand passing a map of arguments
func executeSubCommand(subCmdParserFn model.SubCmdParserFn, subCmdCliHandlerFn model.UseCaseCliHandlerFn, remainingArgs []string) {
	subCmdArgs, _ := utils.GetArgs(remainingArgs, 1, -1)
	argsMap := subCmdParserFn(subCmdArgs)
	subCmdCliHandlerFn(argsMap)
}
