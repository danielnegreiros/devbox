package proxmox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/errs"
	"github.com/danielnegreiros/go-proxmox-cli/internal/domain/entity"
	"github.com/danielnegreiros/go-proxmox-cli/pkg/modules/rest"
)

const contentTypeValue = "application/json"
const contentTypeKey = "Content-Type"


type PoolsData struct {
	Poolid string `json:"poolid"`
}

type PoolsListResponse struct {
	PoolsData []PoolsData `json:"data"`
}

type VmClone struct {
	TmplId  string `json:"vmid"`
	NewVmId string `json:"newid"`
	Name    string `json:"name,omitempty"`
	Node    string `json:"node"`
	Pool    string `json:"pool,omitempty"`
}

type CiConfig struct {
	CiUser      string `json:"ciuser,omitempty"`
	CiPassword  string `json:"cipassword,omitempty"`
	SshKeysFile string `json:"-"`
	Ipconfig    string `json:"ipconfig0,omitempty"`
	SshKeys     string `json:"sshkeys,omitempty"`
	Cores       string `json:"cores,omitempty"`
	Sockets     string `json:"sockets,omitempty"`
}

type DiskSize struct {
	Size string `json:"size"`
	Disk string `json:"disk"`
}

type CreateVmUseCase struct {
	ticket     *entity.TicketData
	ciConfig   CiConfig
	vmClone    VmClone
	vmDiskSize DiskSize
}

func NewCreateVmUseCase(ticket *entity.TicketData, ciConfig CiConfig, vmClone VmClone, vmDiskSize DiskSize) *CreateVmUseCase {
	return &CreateVmUseCase{
		ticket:     ticket,
		ciConfig:   ciConfig,
		vmClone:    vmClone,
		vmDiskSize: vmDiskSize,
	}
}

func (u *CreateVmUseCase) Execute() {

	if u.vmClone.Pool != "" && !u.doesPoolExists() {
		u.createPool()
	}

	u.createVm()

	if u.ciConfig.CiPassword != "" || u.ciConfig.CiUser != "" || u.ciConfig.Ipconfig != "" || u.ciConfig.SshKeysFile != "" {
		u.configCloudInit()
	}


	if u.vmDiskSize.Size != ""{
		u.resizeDisk()
	}
	

	u.start()

}

func (u *CreateVmUseCase) doesPoolExists() bool {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/pools", u.ticket.Host),
		Method:        "GET",
		AcceptedCodes: []int{200},
		Data:          &PoolsListResponse{},
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
	}

	content, _ := httpReq.Execute()

	pools := content.(*PoolsListResponse)
	for _, pool := range pools.PoolsData {
		if pool.Poolid == u.vmClone.Pool {
			return true
		}
	}

	return false

}

func (u *CreateVmUseCase) createPool() int {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/pools", u.ticket.Host),
		Method:        "POST",
		AcceptedCodes: []int{200},
		Data:          &struct{}{},
		Body:          []byte(fmt.Sprintf(`{"poolid": "%s"}`, u.vmClone.Pool)),
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
		Header: map[string]string{
			"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
			contentTypeKey:        contentTypeValue,
		},
	}

	_, code := httpReq.Execute()

	return code
}

func (u CreateVmUseCase) createVm() {

	body, _ := json.Marshal(u.vmClone)
	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%s/clone", u.ticket.Host, u.vmClone.Node, u.vmClone.TmplId),
		Method:        "POST",
		AcceptedCodes: []int{200},
		Data:          &struct{}{},
		Body:          body,
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
		Header: map[string]string{
			"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
			contentTypeKey:        contentTypeValue,
		},
	}

	httpReq.Execute()

}

func (u CreateVmUseCase) configCloudInit() {

	sshContent, err := os.ReadFile(u.ciConfig.SshKeysFile)
	errs.PanicIfErr(err)

	escapedKey := url.QueryEscape(string(sshContent))
	escapedKey = strings.ReplaceAll(escapedKey, "+", "%20")

	u.ciConfig.SshKeys = escapedKey
	body, _ := json.Marshal(u.ciConfig)

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%s/config", u.ticket.Host, u.vmClone.Node, u.vmClone.NewVmId),
		Method:        "PUT",
		AcceptedCodes: []int{200},
		Data:          &struct{}{},
		Body:          body,
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
		Header: map[string]string{
			"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
			contentTypeKey:        contentTypeValue,
		},
	}

	httpReq.Execute()

}

func (u CreateVmUseCase) resizeDisk() {

	body, _ := json.Marshal(u.vmDiskSize)
	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%s/resize", u.ticket.Host, u.vmClone.Node, u.vmClone.NewVmId),
		Method:        "PUT",
		AcceptedCodes: []int{200},
		Data:          &struct{}{},
		Body:          body,
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
		Header: map[string]string{
			"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
			contentTypeKey:        contentTypeValue,
		},
	}

	httpReq.Execute()

}

func (u CreateVmUseCase) start() {

	httpReq := rest.HttpRequest{
		Timeout:       10,
		EndPoint:      fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%s/status/start", u.ticket.Host, u.vmClone.Node, u.vmClone.NewVmId),
		Method:        "POST",
		AcceptedCodes: []int{200},
		Data:          &struct{}{},
		Cookie: &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: u.ticket.Ticket,
		},
		Header: map[string]string{
			"CSRFPreventionToken": u.ticket.CSRFPreventionToken,
		},
	}

	httpReq.Execute()
}
