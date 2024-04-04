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
		Name:        "--create-snapshots\t",
		Description: "Manage Proxmox Snapshots",
		ParseFunc:   ParseCreateSnapshotCommand,
		ExecFunc:    proxmox.CreateSnapshotHandler,
	})
}

func ParseCreateSnapshotCommand(args []string) map[string]string {

	// Define flags
	var endpoint string
	var user string
	var password string
	var node string
	var pool string

	snapCmd := flag.NewFlagSet("common", flag.ExitOnError)

	snapCmd.StringVar(&endpoint, "endpoint", "", "Proxmox HTTPS Endpoint\nAccepts export PROXMOX_ENDPOINT='<PROXMOX_ENDPOINT>'")
	snapCmd.StringVar(&user, "user", "", "SSH Proxmox host user\nAccepts export PROXMOX_USERNAME='<PROXMOX_USERNAME>'")
	snapCmd.StringVar(&password, "password", "", "SSH proxmox host user password\nAccepts export PROXMOX_PASSWORD='<PROXMOX_PASSWORD>'")
	snapCmd.StringVar(&node, "node", "", "Name of the Proxmox node\nAccepts export PROXMOX_NODE='<PROXMOX_NODE>'")
	snapCmd.StringVar(&pool, "pool", "", "Pools of VMs to be created snapshot")

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
		builder.WriteString("# Creating snapshots for all Virtual Machines\n")
		builder.WriteString("$ go run cmd/proxmox/main.go proxmox --create-snapshots -pool all\n\n")
		builder.WriteString("# Creating snapshots for all Virtual Machines in a single pool\n")
		builder.WriteString("$ go run cmd/proxmox/main.go proxmox --create-snapshots -pool prod\n\n")
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

	if pool == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: pool cannot be empty\n")
		os.Exit(1)
	}


	argsMap := make(map[string]string)

	argsMap["endpoint"] = endpoint
	argsMap["user"] = user
	argsMap["password"] = password
	argsMap["node"] = node
	argsMap["pool"] = pool

	return argsMap
}
