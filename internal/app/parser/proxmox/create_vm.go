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
		Name:        "--create-vm\t\t",
		Description: "Create Proxmox Cloud Virtual Machine",
		ParseFunc:   ParseVmCreateTmpl,
		ExecFunc:    proxmox.CreateVmHandler,
	})
}

func ParseVmCreateTmpl(args []string) map[string]string {

	// Define flags
	var user string
	var password string
	var endpoint string
	var vm_user string
	var vm_pass string
	var vm_pub_keys string
	var vm_ip string
	var vm_netmask string
	var vm_template_id string
	var vm_id string
	var vm_name string
	var node string
	var pool string

	templateCmd := flag.NewFlagSet("common", flag.ExitOnError)

	templateCmd.StringVar(&endpoint, "endpoint", "", "Proxmox HTTPS Endpoint\nAccepts export PROXMOX_ENDPOINT='<PROXMOX_ENDPOINT>'")
	templateCmd.StringVar(&node, "node", "", "Proxmox node to create the Virtual Machine")
	templateCmd.StringVar(&vm_user, "vm_user", "", "User to be configured in the VM")
	templateCmd.StringVar(&vm_pass, "vm_pass", "", "Password to be initialized for the user")
	templateCmd.StringVar(&vm_pub_keys, "vm_pub_keys", "", "ssh public key location for your user")
	templateCmd.StringVar(&vm_ip, "vm_ip", "", "IP address to be configured in the VM")
	templateCmd.StringVar(&vm_netmask, "vm_netmask", "", "Netmask in decimal format, example: 24")
	templateCmd.StringVar(&vm_template_id, "vm_template_id", "", "Storage vm_netmask to store the template")
	templateCmd.StringVar(&vm_id, "vm_id", "", "ID of the new Virtual Machine")
	templateCmd.StringVar(&vm_name, "vm_name", "", "Name of the new Virtual Machine")
	templateCmd.StringVar(&pool, "pool", "", "Pool of the new Virtual Machine")

	templateCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: main.go [OPTIONS]\n")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Println()
		templateCmd.PrintDefaults()
		fmt.Println()

		builder := strings.Builder{}
		builder.WriteString("Setting Variables:\n\n")
		builder.WriteString("export PROXMOX_ENDPOINT='<https://PROXMOX-IP:PORT>'\nexport PROXMOX_NODE='<PROXMOX_NODE>'\nexport PROXMOX_USERNAME='<PROXMOX_USERNAME>'\nexport PROXMOX_PASSWORD='<PROXMOX_PASSWORD>'\n\n")

		builder.WriteString("Examples: \n")
		builder.WriteString("$ go run cmd/proxmox/main.go template -vm_user create -vm_pub_keys fedora-38 -vm_ip 1111  -vm_netmask fedora-38-tmpl  -vm_template_id local\n\n")
		builder.WriteString("$ go run cmd/proxmox/main.go template -host <prox-host> -vm_pass 22 -user root -password <pass> -vm_user create -vm_pub_keys ubuntu-latest -vm_ip <template-vm_ip> -vm_netmask ubuntu-template-vm_netmask -vm_template_id <vm_template_id>\n\n")
		builder.WriteString("$ go run cmd/proxmox/main.go template -host <prox-host> -vm_pass 22 -user root -password <pass> -vm_user create -vm_pub_keys fedora-38 -vm_ip <template-vm_ip> -vm_netmask ubuntu-template-vm_netmask -vm_template_id <vm_template_id>\n")
		fmt.Println(builder.String())
	}

	templateCmd.Parse(args)

	fmt.Println()

	parse_err := false
	parse_msg := []string{}

	if endpoint == "" {
		endpoint = os.Getenv("PROXMOX_ENDPOINT")
	}
	if endpoint == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Set proxmox endpoint. $ export PROXMOX_ENDPOINT='<PROXMOX_ENDPOINT>' ")
		os.Exit(1)
	}

	if user == "" {
		user = os.Getenv("PROXMOX_USERNAME")
	}
	if user == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Set proxmox username. $ export PROXMOX_USERNAME='<PROXMOX_USERNAME>' ")
		os.Exit(1)
	}

	if password == "" {
		password = os.Getenv("PROXMOX_PASSWORD")
	}
	if password == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Password cannot be empty\n")
		os.Exit(1)
	}

	if node == "" {
		node = os.Getenv("PROXMOX_NODE")
	}
	if node == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: node cannot be empty\n")
		os.Exit(1)
	}

	if vm_template_id == "" {
		parse_msg = append(parse_msg, "Error: vm_template_id cannot be empty")
		parse_err = true
	}

	if vm_id == "" {
		parse_msg = append(parse_msg, "Error: vm_id cannot be empty")
		parse_err = true
	}

	if parse_err {
		templateCmd.Usage()
		for _, msg := range parse_msg {
			fmt.Println(msg)
		}
		fmt.Println()
		os.Exit(1)
	}

	argsMap := make(map[string]string)

	argsMap["endpoint"] = endpoint
	argsMap["node"] = node
	argsMap["user"] = user
	argsMap["password"] = password

	argsMap["vm_template_id"] = vm_template_id
	argsMap["vm_id"] = vm_id

	if vm_pub_keys != "" { argsMap["vm_pub_keys"] = vm_pub_keys }
	if vm_ip != "" { argsMap["vm_ip"] = vm_ip }
	if vm_netmask != "" { argsMap["vm_netmask"] = vm_netmask }
	if vm_name != "" { argsMap["vm_name"] = vm_name }
	if vm_user != "" { argsMap["vm_user"] = vm_user }
	if vm_pass != "" { argsMap["vm_pass"] = vm_pass }
	if pool != "" { argsMap["pool"] = pool }

	return argsMap
}
