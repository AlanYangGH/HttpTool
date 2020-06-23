package server

import (
	"tool/config"
	"tool/server"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 使用绮讯发送短信
// http://sdk.mb345.com
func (sms *SmsServer) SendSMSUseQiXun(to, content string) (string, error) {
	content = content + config.C.Sms.Sign
	zap.L().Info("QiXun", zap.String("to", to), zap.String("content", content))

	// Create HTTP request client
	client := &http.Client{}
	// https://mb345.com/ws/BatchSend2.aspx?CorpID=%v&Pwd=%v&Mobile=%v&Content=%v&Cell=&SendTime=
	gbkContent, _ := server.Utf8ToGbk([]byte(content))
	resp, err := client.PostForm(fmt.Sprintf(config.C.SmsQiXun.ApiUrl,
		config.C.SmsQiXun.Sid,
		config.C.SmsQiXun.Pwd,
		to,
		url.QueryEscape(string(gbkContent))), url.Values{})
	if err != nil {
		zap.L().Error("new request error", zap.Error(err))
		return "", err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		zap.L().Warn("response status code not 200", zap.Int("code", resp.StatusCode))
		return "", errors.New("response status code not 200")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("read response body error", zap.Error(err))
		return "", err
	}

	zap.L().Info("response", zap.String("body", string(respBody)))
	return string(respBody), nil
}
