package models

type Example struct {
	Message string `json:"message"`
}

func NewExample(id int, message string) *Example {
	return &Example{
		Message: message,
	}
}

func (m *Example) GeExampletMessage() Example {
	return *m
}
