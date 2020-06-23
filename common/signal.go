package common

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

//WaitSignalSynchronized 同步等待系统信号并处理
func WaitSignalSynchronized() {
	sigsCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{})

	signal.Notify(sigsCh, syscall.SIGABRT, syscall.SIGPIPE, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.Signal(10))

	go func() {
		for {
			select {
			case sig := <-sigsCh:
				zap.L().Warn("Receive", zap.String("Signal", fmt.Sprintf("%v", sig)))
				switch sig {
				case syscall.Signal(10):
					close(doneCh)
					return
				case syscall.SIGINT:
					close(doneCh)
					return
				default:
					continue
				}
			}
		}
	}()
	zap.L().Warn("Waiting For Signals")
	<-doneCh
	zap.L().Warn("Need Exit")
}
