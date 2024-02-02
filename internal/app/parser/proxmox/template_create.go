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
		Name:        "--create-template-cloud-init",
		Description: "Create Proxmox Cloud Init Templates",
		ParseFunc:   ParseTemplateCreateCommand,
		ExecFunc:    proxmox.TemplateHandler,
	})
}

func ParseTemplateCreateCommand(args []string) map[string]string {

	// Define flags
	var action string
	var host string
	var port string
	var user string
	var password string
	var image string
	var id string
	var name string
	var storage string
	var net string

	templateCmd := flag.NewFlagSet("common", flag.ExitOnError)

	templateCmd.StringVar(&action, "action", "", "Action: create, destroy")
	templateCmd.StringVar(&host, "host", "", "IP or DNS Name of Proxmox host\nAccepts export PROXMOX_HOST='<PROXMOX-IP>'")
	templateCmd.StringVar(&user, "user", "", "SSH Proxmox host user\nAccepts export PROXMOX_SSH_USER='<PROXMOX-UER>'")
	templateCmd.StringVar(&password, "password", "", "SSH proxmox host user password\nAccepts export PROXMOX_SSH_PASS='<PROXMOX-SSH-PASS>'")
	templateCmd.StringVar(&port, "port", "22", "Connection port of Proxmox host")
	templateCmd.StringVar(&image, "image", "", "Cloud init image to be templated")
	templateCmd.StringVar(&id, "id", "", "ID of the newly created template")
	templateCmd.StringVar(&name, "name", "", "Template name to be breated")
	templateCmd.StringVar(&storage, "storage", "", "Storage name to store the template")
	templateCmd.StringVar(&net, "net", "vmbr0", "Net adapter to VM")

	templateCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: main.go [OPTIONS]\n")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Println()
		templateCmd.PrintDefaults()
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

	templateCmd.Parse(args)

	fmt.Println()

	if host == "" {
		host = os.Getenv("PROXMOX_SSH_HOST")
	}
	if host == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: host cannot be empty\n")
		os.Exit(1)
	}

	if user == "" {
		user = os.Getenv("PROXMOX_SSH_USER")
	}
	if user == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Uer cannot be empty\n")
		os.Exit(1)
	}

	if password == "" {
		password = os.Getenv("PROXMOX_SSH_PASS")
	}
	if password == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Password cannot be empty\n")
		os.Exit(1)
	}

	if name == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: name cannot be empty\n")
		os.Exit(1)
	}

	if image == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Image cannot be empty\n")
		os.Exit(1)
	}

	if id == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Template ID cannot be empty\n")
		os.Exit(1)
	}

	if storage == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Template Storage cannot be empty\n")
		os.Exit(1)
	}

	if !(action == "create" || action == "destroy") {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: action should be create or destroy\n")
		os.Exit(1)
	}

	argsMap := make(map[string]string)

	argsMap["action"] = action
	argsMap["host"] = host
	argsMap["port"] = port
	argsMap["user"] = user
	argsMap["password"] = password
	argsMap["image"] = image
	argsMap["id"] = id
	argsMap["name"] = name
	argsMap["storage"] = storage
	argsMap["net"] = net

	return argsMap
}
