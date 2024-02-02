package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
)

func Exit(usageText string) {
	fmt.Fprint(os.Stderr, usageText)
	fmt.Println()
	os.Exit(1)
}

func IsOptioninChildSubCmds(pattern string, childSubCmds []model.SubCmd) (bool, model.SubCmd) {
	for _, item := range childSubCmds {
		cleanName := strings.ReplaceAll(item.Name, "\t", "")
		if cleanName == pattern {
			return true, item
		}
	}
	return false, model.SubCmd{}
}

// GetArgs returns two sllice of strings
// The first one is the filtered paramenters in the argument
// The second one is remaining args out of the filter.
// When end = -1, it will return until the last element.
func GetArgs(args []string, begin int, end int) ([]string, []string) {
	if len(args) < end || end == -1 {
		return args[begin:], []string{}
	}

	return args[begin:end], args[end:]
}
