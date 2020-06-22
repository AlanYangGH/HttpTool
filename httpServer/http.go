package httpServer

import (
	"context"
	"email/config"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"sync"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func NewServer() *Server {
	s := &Server{
		wg: &sync.WaitGroup{},
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	go s.serveHTTP()
	return s
}

//Close Close
func (s *Server) Close() {
	s.cancel()
	s.wg.Wait()
}

func (s *Server) serveHTTP() {
	s.wg.Add(1)
	quit := make(chan struct{})
	defer func() {
		if p := recover(); p != nil {
			zap.L().Error("Start error")
		}
		if quit != nil {
			close(quit)
			quit = nil
		}
		s.wg.Done()
	}()
	go func() {
		http.HandleFunc("/", s.processHTTP)
		http.HandleFunc("/send_email", s.SendEmail)
		port := config.EmailConfig.HTTP.Port
		zap.L().Info("The HTTP server listen", zap.Int("port", port))
		err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		if err != nil {
			zap.L().Error("HTTP server listen error", zap.Error(err))
		}
		close(quit)
		quit = nil
	}()
	select {
	case <-s.ctx.Done():
		return
	case <-quit:
		return
	}
}

func (s *Server) processHTTP(w http.ResponseWriter, req *http.Request) {
	query := req.PostForm

	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(fmt.Sprintf("%v,%v", query, string(body)))
	fmt.Println(fmt.Sprintf("%v", query["code"]))

	w.Write([]byte(fmt.Sprintf("%v", query["code"])))
}
