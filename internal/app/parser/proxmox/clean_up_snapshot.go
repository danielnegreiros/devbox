package proxmox

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/cli_handler/proxmox"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
)

func init() {
	model.AvailSubCmds = append(model.AvailSubCmds, model.SubCmd{
		Parent:      "proxmox",
		Name:        "--clean-up-snapshots\t",
		Description: "Manage Proxmox Snapshots",
		ParseFunc:   ParseSnapshotCommand,
		ExecFunc:    proxmox.CleanUpSnapshotHandler,
	})
}

func ParseSnapshotCommand(args []string) map[string]string {

	// Define flags
	var endpoint string
	var user string
	var password string
	var node string
	var days string
	var include string
	var exclude string

	snapCmd := flag.NewFlagSet("common", flag.ExitOnError)

	snapCmd.StringVar(&endpoint, "endpoint", "", "Proxmox HTTPS Endpoint\nAccepts export PROXMOX_ENDPOINT='<PROXMOX_ENDPOINT>'")
	snapCmd.StringVar(&user, "user", "", "SSH Proxmox host user\nAccepts export PROXMOX_USERNAME='<PROXMOX_USERNAME>'")
	snapCmd.StringVar(&password, "password", "", "SSH proxmox host user password\nAccepts export PROXMOX_PASSWORD='<PROXMOX_PASSWORD>'")
	snapCmd.StringVar(&node, "node", "", "Name of the Proxmox node\nAccepts export PROXMOX_NODE='<PROXMOX_NODE>'")
	snapCmd.StringVar(&days, "days", "", "Delete snapshots older than x days")
	snapCmd.StringVar(&include, "include", "dummy", "Regex of snapshots names to be included")
	snapCmd.StringVar(&exclude, "exclude", "dummy", "Regex of snapshots names to be skipped")

	snapCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: main.go [OPTIONS]\n")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Println()
		snapCmd.PrintDefaults()
		fmt.Println()

		builder := strings.Builder{}
		builder.WriteString("Setting Variables:\n\n")
		builder.WriteString("export PROXMOX_ENDPOINT='<https://PROXMOX-IP:PORT>'\nexport PROXMOX_NODE='<PROXMOX_NODE>'\nexport PROXMOX_USERNAME='<PROXMOX_USERNAME>'\nexport PROXMOX_PASSWORD='<PROXMOX_PASSWORD>'\n\n")

		builder.WriteString("Examples: \n\n")
		builder.WriteString("# Deleting ALL snapshots older than 5 days using acess data from Env varibles\n")
		builder.WriteString("$ go run cmd/proxmox/main.go proxmox --clean-up-snapshots -days 5 -include all\n\n")
		builder.WriteString("# Deleting snapshots older than 15 days that name contains CHANGE-SNAP but avoiding snap that name contains CHANGE-SNAP-PERM using all params from cli\n")
		builder.WriteString("$ go run cmd/proxmox/main.go proxmox --clean-up-snapshots -endpoint https://<IP>:8006 -user root@pam -password <PASS> -node <PROXMOX> -days 15 -include <CHANGE-SNAP> -exclude <CHANGE-SNAP-PERM>\n\n")
		builder.WriteString("# Deleting every single one of snapshots\n")
		builder.WriteString("$ go run cmd/proxmox/main.go proxmox --clean-up-snapshots -days -1 -include all\n")

		fmt.Println(builder.String())
	}

	snapCmd.Parse(args)

	fmt.Println()

	if endpoint == "" {
		endpoint = os.Getenv("PROXMOX_ENDPOINT")
	}
	if endpoint == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: endpoint cannot be empty\n")
		os.Exit(1)
	}

	if user == "" {
		user = os.Getenv("PROXMOX_USERNAME")
	}
	if user == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Uer cannot be empty\n")
		os.Exit(1)
	}

	if password == "" {
		password = os.Getenv("PROXMOX_PASSWORD")
	}
	if password == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Password cannot be empty\n")
		os.Exit(1)
	}

	if node == "" {
		node = os.Getenv("PROXMOX_NODE")
	}
	if node == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: node cannot be empty\n")
		os.Exit(1)
	}

	if days == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Days cannot be empty\n")
		os.Exit(1)
	}

	if include == "all" {
		include = ".*"
	}

	if include == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Include cannot be empty\n")
		os.Exit(1)
	}

	if exclude == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Exclude cannot be empty\n")
		os.Exit(1)
	}

	argsMap := make(map[string]string)

	argsMap["endpoint"] = endpoint
	argsMap["user"] = user
	argsMap["password"] = password
	argsMap["node"] = node
	argsMap["days"] = days
	argsMap["include"] = include
	argsMap["exclude"] = exclude

	return argsMap
}
