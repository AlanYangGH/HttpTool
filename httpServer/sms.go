package httpServer

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tool/config"
	smsServer "tool/server/sms"
)

func (s *Server) SendSms(w http.ResponseWriter, req *http.Request) {
	token := req.PostFormValue("token")
	area := req.PostFormValue("area")
	areaVal, _ := strconv.Atoi(area)
	to := req.PostFormValue("to")
	toVal, _ := strconv.Atoi(to)
	code := req.PostFormValue("code")
	codeVal, _ := strconv.Atoi(code)

	zap.L().Info("Param", zap.String("token", token), zap.String("area", area), zap.String("to", to))

	type SendReturn struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	if toVal == 0 || areaVal == 0 || codeVal == 0 {
		jsons, _ := json.Marshal(SendReturn{400, "Parameter error!"})
		w.Write(jsons)
		return
	}

	// 验证Token
	if token != "b4ac1d1acdc899960e8225b7167f57fb" {
		zap.L().Info("Token authentication failed.")
		jsons, _ := json.Marshal(SendReturn{400, fmt.Sprintf("Token authentication failed. Token: %v", token)})
		w.Write(jsons)
		return
	}

	var sms smsServer.SmsServer
	var funName string
	var apiErr error
	if areaVal == 86 {
		funName = "SendSMSUseQiXun"
		_, apiErr = sms.SendSMSUseQiXun(to, fmt.Sprintf(config.C.Sms.ContentCn, code))
	} else {
		funName = "SendSMSUseTwilio"
		_, apiErr = sms.SendSMSUseTwilio(fmt.Sprintf("+%v%v", area, to), fmt.Sprintf(config.C.Sms.ContentEn, code))
	}

	if apiErr != nil {
		zap.L().Info(funName, zap.String("error", apiErr.Error()))
		jsons, _ := json.Marshal(SendReturn{400, funName + " Return Error."})

		w.Write(jsons)
		return
	}

	jsons, _ := json.Marshal(SendReturn{1, "SUCCESS"})
	w.Write(jsons)
	return
}
