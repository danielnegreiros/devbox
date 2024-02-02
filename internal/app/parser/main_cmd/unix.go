package main_cmd

import (
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/punix"
)

func init() {
	model.AvailGroupCmd = append(model.AvailGroupCmd, model.MainCmd{
		Name:          "unix\t",
		Description:   "Manage Unix Servers",
		GroupParserFn: punix.ParseUnixCommand,
	})
}
