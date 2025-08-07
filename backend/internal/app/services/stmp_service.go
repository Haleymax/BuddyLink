package services

import (
	"buddylink/pkg/cache_client"
	"buddylink/pkg/email"
	"log"
	"math/rand"
	"time"
)

type StmpService interface {
	SendVerification(email string, redis cache_client.RedisClient) error
	VerifyCode(email string, code string, redis cache_client.RedisClient) (bool, error)
}

type StmpServiceImpl struct {
	Client email.EmailClient
}

func NewStmpService() StmpService {
	//stmp_client, err := email.NewEmailClient()
	//if err != nil {
	//	panic(err)
	//}
	//
	//return &StmpServiceImpl{
	//	Client: stmp_client,
	//}
	return &StmpServiceImpl{
		Client: nil,
	}
}

func (s *StmpServiceImpl) SendVerification(email string, redis cache_client.RedisClient) error {
	verificationCode := generateVerificationCode()
	//subject := "verification code"
	//content := "your verification code is " + verificationCode
	//err := s.Client.SendMail(email, subject, content)
	//if err != nil {
	//	log.Printf("Error sending verification email %v", err)
	//	return err
	//}
	//
	key := email

	err := redis.Set(key, []byte(verificationCode), 2*time.Minute)
	if err != nil {
		log.Printf("Error sending verification email %v", err)
		return err
	}
	return nil
}

func (s *StmpServiceImpl) VerifyCode(email string, code string, redis cache_client.RedisClient) (bool, error) {
	key := email
	storedCode, err := redis.Get(key)
	if err != nil {
		log.Printf("Error retrieving verification code from Redis: %v", err)
		return false, err
	}
	if string(storedCode) != code {
		log.Printf("Verification code does not match for email %s", email)
		return false, nil
	}
	// 删除验证码
	err = redis.Del(key)
	if err != nil {
		log.Printf("Error deleting verification code from Redis: %v", err)
		return false, err
	}
	return true, nil
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
