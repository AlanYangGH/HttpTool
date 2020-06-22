package httpServer

import (
	server "email/server/email"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (s *Server) SendEmail(w http.ResponseWriter, req *http.Request) {
	token := req.PostFormValue("token")
	emailAddress := req.PostFormValue("email_address")
	title := req.PostFormValue("title")
	content := req.PostFormValue("content")

	zap.L().Info("Param",
		zap.String("token", token),
		zap.String("email_address", emailAddress),
		zap.String("content", content))

	type SendReturn struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	type SendCloudApiResponse struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}

	// 验证Token
	if token != "b4ac1d1acdc899960e8225b7167f57fb" {
		zap.L().Info("Token authentication failed.")
		jsons, _ := json.Marshal(SendReturn{400, fmt.Sprintf("Token authentication failed. Token: %v", token)})
		w.Write(jsons)
		return
	}

	// 发送邮件
	if emailAddress == "" {
		jsons, _ := json.Marshal(SendReturn{400, "Email address is invalid!"})
		w.Write(jsons)
		return
	}

	emailAddress = strings.Replace(strings.Replace(strings.Replace(emailAddress, `["`, "", -1), `"]`, "", -1), `","`, ";", -1)
	resp, err := server.SendSimpleMessage(title, content, emailAddress)
	if err != nil {
		zap.L().Info("MailGunReturnError", zap.String("error", err.Error()))
		jsons, _ := json.Marshal(SendReturn{400, "MailGun Return Error."})

		w.Write(jsons)
		return
	}

	zap.L().Info("MailGunReturnSuccess", zap.String("MailGun resp", resp))

	jsons, _ := json.Marshal(SendReturn{1, "SUCCESS"})
	w.Write(jsons)
	return
}


