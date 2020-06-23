package server

import (
	"tool/config"
	"tool/server"
	"flag"
	"fmt"
	"os"
	"testing"
)

func Setup() {
	fmt.Println("--------------------> setup")

	configPath := flag.String("cfg", "../../", "path for config file directory")
	if *configPath == "" {
		fmt.Println("You must specify the config file path with -path")
		return
	}
	if err := config.LoadConfig(*configPath); err != nil {
		fmt.Println("LoadConfig Failed.", err.Error())
		return
	}
}

func Teardown() {
	fmt.Println("--------------------> teardown")
}

func TestMain(m *testing.M) {
	// Setup code
	Setup()
	// Run test
	code := m.Run()
	// Teardown code
	Teardown()
	// Exit
	os.Exit(code)
}

// go test -v -run TestSmsServer_SendSMSUseQiXun -args 182xxxxx
func TestSmsServer_SendSMSUseQiXun(t *testing.T) {
	type args struct {
		to      string
		content string
	}
	tt := struct {
		name string
		args args
	}{
		"QiXun",
		args{
			to:      os.Args[3],
			content: fmt.Sprintf(config.C.Sms.ContentCn, server.RandInt(111111, 999999)),
		},
	}

	t.Run(tt.name, func(t *testing.T) {
		sms := &SmsServer{}
		got, err := sms.SendSMSUseQiXun(tt.args.to, tt.args.content)
		t.Logf("got: %v, err: %v", got, err)
	})
}

// go test -v -run TestSmsServer_SendSMSUseTwilio -args 63xxxxxx
func TestSmsServer_SendSMSUseTwilio(t *testing.T) {
	type args struct {
		to      string
		content string
	}
	tt := struct {
		name string
		args args
	}{
		"Twilio",
		args{
			to:      os.Args[3],
			content: fmt.Sprintf(config.C.Sms.ContentEn, server.RandInt(111111, 999999)),
		},
	}

	t.Run(tt.name, func(t *testing.T) {
		sms := &SmsServer{}
		got, err := sms.SendSMSUseTwilio(tt.args.to, tt.args.content)
		t.Logf("got: %v, err: %v", got, err)
	})
}
