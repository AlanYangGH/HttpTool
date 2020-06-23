package server

import (
	"tool/config"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 使用Twilio发送短信
// https://www.twilio.com
func (sms *SmsServer) SendSMSUseTwilio(to, content string) (string, error) {
	content = config.C.Sms.Sign + content
	zap.L().Info("Twilio", zap.String("to", to), zap.String("content", content))

	// Set account keys & information
	urlStr := fmt.Sprintf(config.C.SmsTwilio.ApiUrl, config.C.SmsTwilio.Sid)

	// Create possible message bodies
	quotes := [1]string{content}

	rand.Seed(time.Now().Unix())

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", config.C.SmsTwilio.From)
	msgData.Set("Body", quotes[rand.Intn(len(quotes))])
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, &msgDataReader)
	if err != nil {
		zap.L().Error("new request error", zap.Error(err))
		return "", err
	}

	req.SetBasicAuth(config.C.SmsTwilio.Sid, config.C.SmsTwilio.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("http error", zap.Error(err))
		return "", err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		zap.L().Warn("response status code not 200 or 201", zap.Int("code", resp.StatusCode))
		return "", errors.New("response status code not 200 or 201")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("read response body error", zap.Error(err))
		return "", err
	}

	zap.L().Info("response", zap.String("body", string(respBody)))
	return string(respBody), nil
}
