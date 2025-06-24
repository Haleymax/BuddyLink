package services

import "buddylink/pkg/email"

type StmpService interface {
	SendVerification(email string) error
	CheckVerification(email string) error
}

type StmpServiceImpl struct {
	Client email.EmailClient
}
