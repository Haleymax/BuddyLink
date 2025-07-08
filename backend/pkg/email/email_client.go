package email

import (
	"buddylink/config"
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type EmailClient interface {
	SendMail(to string, subject, content string) error
	Close() error
}

type EmailClientImpl struct {
	client *smtp.Client
}

func NewEmailClient() (EmailClient, error) {
	cfg := config.GetConfig()
	smtpConfig := cfg.SMTP
	addr := fmt.Sprintf("%s:%s", smtpConfig.Host, smtpConfig.Port)
	var conn *smtp.Client
	var err error
	if smtpConfig.SSL {
		tlsConfig := &tls.Config{
			ServerName: smtpConfig.Host,
			MinVersion: tls.VersionTLS12,
		}
		tlsconn, terr := tls.Dial("tcp", addr, tlsConfig)
		if terr != nil {
			return nil, terr
		}
		conn, err = smtp.NewClient(tlsconn, smtpConfig.Host)
	} else {
		conn, err = smtp.Dial(addr)
		if err != nil {
			return nil, err
		}
	}

	auth := smtp.PlainAuth("", smtpConfig.Email, smtpConfig.Password, smtpConfig.Host)
	if err = conn.Auth(auth); err != nil {
		return nil, err
	}
	return &EmailClientImpl{
		client: conn,
	}, nil
}

func (e EmailClientImpl) SendMail(to string, subject, content string) error {

	cfg := config.GetConfig()

	// 设置发件人
	if err := e.client.Mail(cfg.SMTP.Email); err != nil {
		return err
	}

	// 设置收件人
	if err := e.client.Rcpt(to); err != nil {
		return err
	}

	// 发送邮件内容
	w, err := e.client.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	// 构造邮件头
	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n", cfg.SMTP.Email, to[0], subject)

	// 写入邮件内容
	if _, err = fmt.Fprintf(w, "%s%s", headers, content); err != nil {
		return err
	}

	return nil
}

func (e EmailClientImpl) Close() error {
	if e.client != nil {
		return e.client.Quit()
	}
	return nil
}
