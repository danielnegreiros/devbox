package main_cmd

import (
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/proxmox"
)

func init() {
	model.AvailGroupCmd = append(model.AvailGroupCmd, model.MainCmd{
		Name:          "proxmox",
		Description:   "Manage Proxmox Environment",
		GroupParserFn: proxmox.ParseProxmoxCommand,
	})
}
