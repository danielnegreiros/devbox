package entity

type TicketData struct {
	CSRFPreventionToken string `json:"CSRFPreventionToken"`
	Ticket              string `json:"ticket"`
	Username            string `json:"username"`
	Host                string
}

func NewTicketData(token string, ticket string, username string, host string) *TicketData {
	return &TicketData{
		CSRFPreventionToken: token,
		Ticket:              ticket,
		Username:            username,
		Host:                host,
	}
}
