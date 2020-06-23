package server

import (
	"context"
	"tool/config"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"go.uber.org/zap"
	"strings"
	"time"
)

// 使用MailGun发送邮件
// https://www.mailgun.com
// to 收件人邮件地址，多个使用 ";" 隔开
func (mail *MailServer) SendMailUseMailGun(subject, text, to string) (string, error) {
	zap.L().Info("MailGun", zap.String("to", to), zap.String("title", subject), zap.String("content", text))

	mg := mailgun.NewMailgun(config.C.MailMailGun.Domain, config.C.MailMailGun.ApiKey)
	m := mg.NewMessage(
		config.C.MailMailGun.From,
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

	fmt.Println("-----------" + config.C.MailMailGun.From)
	_, id, err := mg.Send(ctx, m)
	return id, err
}
