package services

import (
	"buddylink/config"
	"buddylink/pkg/cache_client"
	"buddylink/pkg/email"
	"log"
	"math/rand"
	"time"
)

type StmpService interface {
	SendVerification(email string, redis cache_client.RedisClient) error
}

type StmpServiceImpl struct {
	Client email.EmailClient
}

func NewStmpService(cfg *config.SMTPConfig) StmpService {

	stmp_client, err := email.NewEmailClient(cfg)
	if err != nil {
		panic(err)
	}

	return &StmpServiceImpl{
		Client: stmp_client,
	}
}

func (s *StmpServiceImpl) SendVerification(email string, redis cache_client.RedisClient) error {
	verificationCode := generateVerificationCode()
	subject := "verification code"
	content := "your verification code is " + verificationCode
	err := s.Client.SendMail(email, subject, content)
	if err != nil {
		log.Printf("Error sending verification email %v", err)
		return err
	}

	key := email
	err = redis.Set(key, []byte(verificationCode), 2*time.Minute)
	if err != nil {
		log.Printf("Error sending verification email %v", err)
		return err
	}
	return nil
}

// 生成验证码
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = byte(rand.Intn(26) + 65)
	}
	return string(code)
}
