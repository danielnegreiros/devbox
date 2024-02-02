package main_cmd

import "github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"

func GetAvailGroupCmds() []model.MainCmd {
	return model.AvailGroupCmd
}
