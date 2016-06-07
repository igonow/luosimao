package luosimao

type SMSRequest struct {
	Mobile   string `json:"mobile"`
	Message  string `json:"message,string"`
	CallBack string `json:"cb,omitempty"`
}

type BatchSMSRequest struct {
	MobileList string `json:"mobile_list"`
	Message    string `json:"message,string"`
	Time       string `json:"time, string"`
	CallBack   string `json:"cb,omitempty"`
}
