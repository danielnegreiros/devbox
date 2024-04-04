package entity

type ProxmoxCredential struct {
	Host string
	User string
	Pass string
}

func NewProxmoxCredential(host, user, pass string) *ProxmoxCredential {
	return &ProxmoxCredential{
		Host: host,
		User: user,
		Pass: pass,
	}
}
