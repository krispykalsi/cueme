package model

type MsgRequest struct {
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Message string `json:"message"`
}

func (m MsgRequest) HasValidPhone() bool {
	return m.Phone != ""
}

func (m MsgRequest) HasValidEmail() bool {
	return m.Email != ""
}
