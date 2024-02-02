<h1>Clean Up Snapshots</h3>


```
Usage: main.go [OPTIONS]
Options:

  -action string
        Action: create, destroy
  -host string
        IP or DNS Name of Proxmox host
        Accepts export PROXMOX_HOST='<PROXMOX-IP>'
  -id string
        ID of the newly created template
  -image string
        Cloud init image to be templated
  -name string
        Template name to be breated
  -net string
        Net adapter to VM (default "vmbr0")
  -password string
        SSH proxmox host user password
        Accepts export PROXMOX_SSH_PASS='<PROXMOX-SSH-PASS>'
  -port string
        Connection port of Proxmox host (default "22")
  -storage string
        Storage name to store the template
  -user string
        SSH Proxmox host user
        Accepts export PROXMOX_SSH_USER='<PROXMOX-UER>'

```

- Setting Variables

```bash
export PROXMOX_HOST='<PROXMOX-IP>'
export PROXMOX_SSH_USER='<PROXMOX-USER>'
export PROXMOX_SSH_PASS='<PROXMOX-PASS>'
```

- Examples

```bash
$ go run cmd/proxmox/main.go template -action create -image fedora-38 -id 1111  -name fedora-38-tmpl  -storage local

$ go run cmd/proxmox/main.go template -host <prox-host> -port 22 -user root -password <pass> -action create -image ubuntu-latest -id <template-id> -name ubuntu-template-name -storage <storage>

$ go run cmd/proxmox/main.go template -host <prox-host> -port 22 -user root -password <pass> -action create -image fedora-38 -id <template-id> -name ubuntu-template-name -storage <storage>
```
