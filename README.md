<h1>DevOps ToolBox</h3>

___

## 1. Reason

This app aims to help developers and operators by serving as a catalog for different automated operations for different environments.

This code can be execute on local environment and on remote environment.

Read desired use case documentation for more information and proper usage

___


## 2. Proxmox Use Cases


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
2024/04/04 23:14:30 Starting command: mkdir -pv /tmp/cloudinit
2024/04/04 23:14:30 Done
2024/04/04 23:14:30 Starting command: wget -P /tmp/cloudinit/ https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img
2024/04/04 23:15:18 Done
2024/04/04 23:15:18 Starting command: qm stop 1111
2024/04/04 23:15:19 Done
2024/04/04 23:15:19 Starting command: qm destroy 1111
2024/04/04 23:15:19 Done
2024/04/04 23:15:19 Starting command: qm create 1111 --memory 2048 --net0 virtio,bridge=vmbr0 --scsihw virtio-scsi-pci
2024/04/04 23:15:20 Done
2024/04/04 23:15:20 Starting command: qm set 1111 --scsi0 local:0,import-from=/tmp/cloudinit/jammy-server-cloudimg-amd64.img
2024/04/04 23:15:23 Done
2024/04/04 23:15:23 Starting command: qm set 1111 --ide2 local:cloudinit
2024/04/04 23:15:24 Done
2024/04/04 23:15:24 Starting command: qm set 1111 --boot order=scsi0
2024/04/04 23:15:24 Done
2024/04/04 23:15:24 Starting command: qm set 1111 --name ubuntu-tmpl
2024/04/04 23:15:25 Done
2024/04/04 23:15:25 Starting command: qm template 1111
2024/04/04 23:15:26 Done
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
2024/04/04 23:21:23 POST /api2/json/access/ticket 200
2024/04/04 23:21:23 GET /api2/json/pools 200
2024/04/04 23:21:23 POST /api2/json/nodes/proxmox/qemu/1111/clone 200
2024/04/04 23:21:23 PUT /api2/json/nodes/proxmox/qemu/1112/config 200
2024/04/04 23:21:23 POST /api2/json/nodes/proxmox/qemu/1112/status/start 200
``` 
</details>
<br />



To be completed

| Environment | Use Case                    | Connection       | Documentation | Status                       |
|-------------|-----------------------------|------------------|---------------|------------------------------|
| Proxmox     | Clean Up Snapshots          | Remote API       | Link          | :white_check_mark:           |
| Proxmox     | Create Cloud Init Templates | Local/Remote SSH | Link          | :white_check_mark:           |

___

## 3. Unix

To be completed

| Environment | Use Case                    | Connection       | Documentation | Status                       |
|-------------|-----------------------------|------------------|---------------|------------------------------|
| Unix        | Harden SSH                  | Remote SSH       | Link          | :chart_with_upwards_trend:   |

