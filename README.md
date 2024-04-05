<h1>DevOps ToolBox</h3>

___

## 1. Reason

This app aims to help developers and operators by serving as a catalog for different automated operations for different environments.

This code can be execute on local environment and on remote environment.

Read desired use case documentation for more information and proper usage

___


## 2. Proxmox Use Cases

- Installation

```bash
sudo wget -O /usr/local/bin/devbox https://github.com/danielnegreiros/devbox/releases/download/v0.1.0/devbox
sudo chmod +x /usr/local/bin/devbox
```

- Configure Credentials 

```bash
export PROXMOX_ENDPOINT='https://192.168.0.99:8006'
export PROXMOX_NODE='proxmox'
export PROXMOX_USERNAME='root@pam'
export PROXMOX_PASSWORD='<PASS>'
```


- Create first VM to be used as template

```bash
devbox proxmox --create-template-cloud-init \
-action create \
-image ubuntu-latest \
-id 1111  \
-name ubuntu-tmpl  \
-storage local \
-host 192.168.0.99 \
-user root \
-password <SSH-PASS>
```

<details> <summary>Expand Logs</summary>  

```bash
2024/04/04 23:26:37 Starting command: mkdir -pv /tmp/cloudinit
2024/04/04 23:26:37 Starting command: wget -P /tmp/cloudinit/ https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img
2024/04/04 23:27:23 Starting command: qm stop 1111
2024/04/04 23:27:23 Starting command: qm destroy 1111
2024/04/04 23:27:24 Starting command: qm create 1111 --memory 2048 --net0 virtio,bridge=vmbr0 --scsihw virtio-scsi-pci
2024/04/04 23:27:25 Starting command: qm set 1111 --scsi0 local:0,import-from=/tmp/cloudinit/jammy-server-cloudimg-amd64.img
2024/04/04 23:27:27 Starting command: qm set 1111 --ide2 local:cloudinit
2024/04/04 23:27:28 Starting command: qm set 1111 --boot order=scsi0
2024/04/04 23:27:29 Starting command: qm set 1111 --name ubuntu-tmpl
2024/04/04 23:27:29 Starting command: qm template 1111
``` 
</details>
<br />

- Create VM

```bash
devbox proxmox --create-vm \
-vm_template_id 1111 \
-vm_id 1112 \
-vm_user my_user \
-vm_pass mypass \
-vm_pub_keys ~/.ssh/id_rsa.pub \
-vm_ip 10.10.100.3 \
-vm_netmask 24 \
-gateway 10.10.100.1 \
-vm_name myvm \
-pool test_pool
```

<details> <summary>Expand Logs</summary>  

```bash
2024/04/04 23:27:54 POST /api2/json/access/ticket 200
2024/04/04 23:27:54 GET /api2/json/pools 200
2024/04/04 23:27:54 POST /api2/json/nodes/proxmox/qemu/1111/clone 200
2024/04/04 23:27:54 PUT /api2/json/nodes/proxmox/qemu/1112/config 200
2024/04/04 23:27:54 POST /api2/json/nodes/proxmox/qemu/1112/status/start 200
``` 
</details>
<br />

- Create snapshots for all vms in a pool

```bash
devbox proxmox --create-snapshots -pool test_pool
```

<details> <summary>Expand Logs</summary>  

```bash
2024/04/04 23:28:51 POST /api2/json/access/ticket 200
2024/04/04 23:28:51 GET /api2/json/pools/test_pool 200
2024/04/04 23:28:51 POST /api2/json/nodes/proxmox/qemu/1112/snapshot 200
``` 
</details>
<br />

- Delete snapshots starting with daily older than 7 days for all VMs.

```bash
devbox proxmox --clean-up-snapshots -days 7 -include daily
```

<details> <summary>Expand Logs</summary>  

```bash
2024/04/04 23:30:27 POST /api2/json/access/ticket 200
2024/04/04 23:30:27 GET /api2/json/nodes/proxmox/qemu 200
2024/04/04 23:30:27 GET /api2/json/nodes/proxmox/qemu/1111/snapshot 200
2024/04/04 23:30:27 DELETE /api2/json/nodes/proxmox/qemu/1112/snapshot/daily_2024-04-04__23_28_51 200
``` 
</details>
<br />


- Configure crontab for automatic protection and cleanup

```bash
0 17 * * * devbox proxmox --create-snapshots -pool kubernetes
0 17 * * * devbox proxmox --create-snapshots -pool prod
0 18 * * * devbox proxmox --clean-up-snapshots -days 7 -include daily
```


___

## 3. Unix Use Cases

To be completed

| Environment | Use Case                    | Connection       | Documentation | Status                       |
|-------------|-----------------------------|------------------|---------------|------------------------------|
| Unix        | Harden SSH                  | Remote SSH       | Link          | :chart_with_upwards_trend:   |

