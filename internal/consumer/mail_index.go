package consumer

type MailIndex struct {
	Id       int64  `json:"id"`
	Uid      int64  `json:"uid"`
	Type     int    `json:"type"`
	MailFrom string `json:"from"`
	MailTo   string `json:"to"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
	SendTime int64  `json:"send_time"`
}

type MailMsg struct {
	Operation string `json:"operation"`
	MailIndex
}
