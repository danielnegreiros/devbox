package punix

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/cli_handler/unix"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
)

func init() {
	model.AvailSubCmds = append(model.AvailSubCmds, model.SubCmd{
		Parent:      "unix",
		Name:        "--harden-ssh\t",
		Description: "Manage Proxmox Snapshots",
		ParseFunc:   ParseSnapshotCommand,
		ExecFunc:    unix.HardenSshHandler,
	})
}

func ParseSnapshotCommand(args []string) map[string]string {

	// Define flags
	var action string
	var host string
	var user string
	var password string

	snapCmd := flag.NewFlagSet("common", flag.ExitOnError)

	snapCmd.StringVar(&action, "action", "", "Action: create, delete")
	snapCmd.StringVar(&host, "host", "", "IP or DNS Name of Proxmox host\nAccepts export PROXMOX_HOST='<PROXMOX-ENDPOINT>'")
	snapCmd.StringVar(&user, "user", "", "SSH Proxmox host user\nAccepts export PROXMOX_USER='<PROXMOX-USER>'")
	snapCmd.StringVar(&password, "password", "", "SSH proxmox host user password\nAccepts export PROXMOX_PASS='<PROXMOX-PASS>'")

	snapCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: main.go [OPTIONS]\n")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Println()
		snapCmd.PrintDefaults()
		fmt.Println()

		builder := strings.Builder{}
		builder.WriteString("Setting Variables:\n\n")
		builder.WriteString("export PROXMOX_HOST='<PROXMOX-IP>'\nexport PROXMOX_SSH_USER='<PROXMOX-USER>'\nexport PROXMOX_SSH_PASS='<PROXMOX-PASS>'\n\n")

		builder.WriteString("Examples: \n")
		builder.WriteString("$ go run cmd/proxmox/main.go template -action create -image fedora-38 -id 1111  -name fedora-38-tmpl  -storage local\n\n")
		builder.WriteString("$ go run cmd/proxmox/main.go template -host <prox-host> -port 22 -user root -password <pass> -action create -image ubuntu-latest -id <template-id> -name ubuntu-template-name -storage <storage>\n\n")
		builder.WriteString("$ go run cmd/proxmox/main.go template -host <prox-host> -port 22 -user root -password <pass> -action create -image fedora-38 -id <template-id> -name ubuntu-template-name -storage <storage>\n")
		fmt.Println(builder.String())
	}

	snapCmd.Parse(args)

	fmt.Println()

	if host == "" {
		host = os.Getenv("PROXMOX_HOST")
	}
	if host == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: host cannot be empty\n")
		os.Exit(1)
	}

	if user == "" {
		user = os.Getenv("PROXMOX_USER")
	}
	if user == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Uer cannot be empty\n")
		os.Exit(1)
	}

	if password == "" {
		password = os.Getenv("PROXMOX_PASS")
	}
	if password == "" {
		snapCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Password cannot be empty\n")
		os.Exit(1)
	}

	argsMap := make(map[string]string)

	argsMap["action"] = action
	argsMap["host"] = host
	argsMap["user"] = user
	argsMap["password"] = password

	return argsMap
}
