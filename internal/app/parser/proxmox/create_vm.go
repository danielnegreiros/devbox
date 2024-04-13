package proxmox

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/cli_handler/proxmox"
	"github.com/danielnegreiros/go-proxmox-cli/internal/app/parser/model"
)

type Config struct {
	Endpoint     string
	User         string
	Password     string
	Node         string
	VMUser       string
	VMPass       string
	VMPubKeys    string
	VMIP         string
	VMNetmask    string
	Gateway      string
	VMTemplateID string
	VMID         string
	VMName       string
	Pool         string
	Cores        string
	Sockets      string
	DiskSize     string
}

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

	var cfg Config

	templateCmd := flag.NewFlagSet("common", flag.ExitOnError)

	templateCmd.StringVar(&cfg.Endpoint, "endpoint", "", "Proxmox HTTPS Endpoint\nAccepts export PROXMOX_ENDPOINT='<PROXMOX_ENDPOINT>'")
	templateCmd.StringVar(&cfg.Node, "node", "", "Proxmox node to create the Virtual Machine")
	templateCmd.StringVar(&cfg.VMUser, "vm_user", "", "User to be configured in the VM")
	templateCmd.StringVar(&cfg.VMPass, "vm_pass", "", "Password to be initialized for the user")
	templateCmd.StringVar(&cfg.VMPubKeys, "vm_pub_keys", "", "ssh public key location for your user")
	templateCmd.StringVar(&cfg.VMIP, "vm_ip", "", "IP address to be configured in the VM")
	templateCmd.StringVar(&cfg.Cores, "cores", "", "How many CPU cores")
	templateCmd.StringVar(&cfg.Sockets, "sockets", "", "How many CPU sockets")
	templateCmd.StringVar(&cfg.DiskSize, "disk_size", "", "How many CPU sockets")
	templateCmd.StringVar(&cfg.VMNetmask, "vm_netmask", "", "Netmask in decimal format, example: 24")
	templateCmd.StringVar(&cfg.Gateway, "gateway", "", "Gateway IP")
	templateCmd.StringVar(&cfg.VMTemplateID, "vm_template_id", "", "Storage vm_netmask to store the template")
	templateCmd.StringVar(&cfg.VMID, "vm_id", "", "ID of the new Virtual Machine")
	templateCmd.StringVar(&cfg.VMName, "vm_name", "", "Name of the new Virtual Machine")
	templateCmd.StringVar(&cfg.Pool, "pool", "", "Pool of the new Virtual Machine")

	templateCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: main.go [OPTIONS]\n")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Println()
		templateCmd.PrintDefaults()
		fmt.Println()

		builder := strings.Builder{}
		builder.WriteString("Setting Variables:\n\n")
		builder.WriteString("export PROXMOX_ENDPOINT='<https://PROXMOX-IP:PORT>'\nexport PROXMOX_NODE='<PROXMOX_NODE>'\nexport PROXMOX_USERNAME='<PROXMOX_USERNAME>'\nexport PROXMOX_PASSWORD='<PROXMOX_PASSWORD>'\n\n")

		builder.WriteString("Examples: \n\n")
		builder.WriteString("$ devbox proxmox --create-vm -vm_template_id 1111 -vm_id 1112 -vm_user my_user -vm_pass mypass -sockets 1 -cores 1 -disk_size 30G -vm_pub_keys ~/.ssh/id_rsa.pub -vm_ip 10.10.100.3 -vm_netmask 24 -gateway 10.10.100.1 -vm_name myvm -pool test_pool\n")
		fmt.Println(builder.String())
	}

	templateCmd.Parse(args)

	fmt.Println()

	parse_err := false
	parse_msg := []string{}

	if cfg.Endpoint == "" {
		cfg.Endpoint = os.Getenv("PROXMOX_ENDPOINT")
	}
	if cfg.Endpoint == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Set proxmox endpoint. $ export PROXMOX_ENDPOINT='<PROXMOX_ENDPOINT>' ")
		os.Exit(1)
	}

	if cfg.User == "" {
		cfg.User = os.Getenv("PROXMOX_USERNAME")
	}
	if cfg.User == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Set proxmox username. $ export PROXMOX_USERNAME='<PROXMOX_USERNAME>' ")
		os.Exit(1)
	}

	if cfg.Password == "" {
		cfg.Password = os.Getenv("PROXMOX_PASSWORD")
	}
	if cfg.Password == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: Password cannot be empty\n")
		os.Exit(1)
	}

	if cfg.Node == "" {
		cfg.Node = os.Getenv("PROXMOX_NODE")
	}
	if cfg.Node == "" {
		templateCmd.Usage()
		fmt.Fprint(os.Stderr, "Error: node cannot be empty\n")
		os.Exit(1)
	}

	if cfg.VMTemplateID == "" {
		parse_msg = append(parse_msg, "Error: vm_template_id cannot be empty")
		parse_err = true
	}

	if cfg.VMID == "" {
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

	return mountData(cfg)
}

func mountData(cfg Config) map[string]string {
	argsMap := make(map[string]string)

	argsMap["endpoint"] = cfg.Endpoint
	argsMap["node"] = cfg.Node
	argsMap["user"] = cfg.User
	argsMap["password"] = cfg.Password

	argsMap["vm_template_id"] = cfg.VMTemplateID
	argsMap["vm_id"] = cfg.VMID

	if cfg.VMPubKeys != "" {
		argsMap["vm_pub_keys"] = cfg.VMPubKeys
	}
	if cfg.VMIP != "" {
		argsMap["vm_ip"] = cfg.VMIP
	}
	if cfg.VMNetmask != "" {
		argsMap["vm_netmask"] = cfg.VMNetmask
	}
	if cfg.Gateway != "" {
		argsMap["gateway"] = cfg.Gateway
	}
	if cfg.VMName != "" {
		argsMap["vm_name"] = cfg.VMName
	}
	if cfg.VMUser != "" {
		argsMap["vm_user"] = cfg.VMUser
	}
	if cfg.VMPass != "" {
		argsMap["vm_pass"] = cfg.VMPass
	}
	if cfg.Pool != "" {
		argsMap["pool"] = cfg.Pool
	}

	if cfg.Cores != "" {
		argsMap["cores"] = cfg.Cores
	}
	if cfg.Sockets != "" {
		argsMap["sockets"] = cfg.Sockets
	}

	if cfg.DiskSize != "" {
		argsMap["diskSize"] = cfg.DiskSize
	}

	return argsMap
}
