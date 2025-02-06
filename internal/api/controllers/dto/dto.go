package dto

type SendMailReq struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	IsHtml  bool     `json:"isHtml"`
}
