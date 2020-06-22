package server

import (
	"context"
	"email/config"
	"github.com/mailgun/mailgun-go/v3"
	"strings"
	"time"
)

// 使用MailGun发送邮件
// https://www.mailgun.com
// to 收件人邮件地址，多个使用 ";" 隔开
func (mail *MailServer) SendMailUsingMailGun(subject, text, to string) (string, error) {
	mg := mailgun.NewMailgun(config.C.MailGun.Domain, config.C.MailGun.ApiKey)
	m := mg.NewMessage(
		config.C.MailGun.From,
		subject,
		text,
	)

	toEmailMap := strings.Split(to, ";")
	for k, mail := range toEmailMap {
		prefix := strings.Split(mail, "@")[0]

		m.AddRecipientAndVariables(mail, map[string]interface{}{
			"first": prefix,
			"id":    k,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
