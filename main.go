package main

import (
	"tool/common"
	"tool/config"
	"tool/httpServer"
	"tool/internel"
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
	internel.SetupLogger("log/log.log", zap.DebugLevel)
	defer internel.CloseLogger()

	configPath := flag.String("cfg", "./", "path for config file directory")

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
