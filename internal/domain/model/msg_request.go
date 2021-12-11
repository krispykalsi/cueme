package model

import "time"

type MsgRequest struct {
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Message     string `json:"message"`
	TimeRFC3339 string `json:"time,omitempty"`
}

func (m MsgRequest) HasValidPhone() bool {
	return m.Phone != ""
}

func (m MsgRequest) HasValidEmail() bool {
	return m.Email != ""
}

func (m MsgRequest) GetCueTime() (*time.Time, error) {
	if m.TimeRFC3339 == "" {
		t := time.Now()
		return &t, nil
	}
	t, err := time.Parse(time.RFC3339, m.TimeRFC3339)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
