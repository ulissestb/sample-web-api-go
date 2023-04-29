package models

type Example struct {
	Message string `json:"message"`
}

func NewExample(message string) *Example {
	return &Example{
		Message: message,
	}
}

func (m *Example) GeExampletMessage() Example {
	return *m
}
