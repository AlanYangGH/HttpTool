package server

import (
	"tool/config"
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

// go test -v -run TestMailServer_SendMailUseMailGun -args xxx@wehi.cc
func TestMailServer_SendMailUseMailGun(t *testing.T) {
	type args struct {
		subject string
		text    string
		to      string
	}
	tt := struct {
		name string
		args args
	}{
		"MailGun",
		args{
			subject: "Test",
			text:    "Hello",
			to:      os.Args[3],
		},
	}
	t.Run(tt.name, func(t *testing.T) {
		mail := &MailServer{}
		got, err := mail.SendMailUseMailGun(tt.args.subject, tt.args.text, tt.args.to)
		if err != nil {
			t.Error(err)
		}
		t.Logf("got: %v, err: %v", got, err)
	})
}
