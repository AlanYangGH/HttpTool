package main

import (
	"email/common"
	"email/config"
	"email/httpServer"
	"email/internel"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	_ "net/http/pprof"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	configPath := flag.String("cfg", "./", "path for config file directory")

	internel.SetupLogger("log/log.log", zap.DebugLevel)
	defer internel.CloseLogger()

	if *configPath == "" {
		log.Fatalln("You must specify the config file path with -path")
		return
	}

	if err := config.LoadConfig(*configPath); err != nil {
		fmt.Println("LoadConfig Failed.", err.Error())
		return
	}

	httpSvr := httpServer.NewServer()
	common.WaitSignalSynchronized()
	httpSvr.Close()
}
