package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lestrrat-go/server-starter/listener"
)

func main() {
	// initialize signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// check socket from Server::Starter
	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	// start web-server
	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid: %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	// shutdown when received SIGTERM
	<-signals
	server.Shutdown(context.Background())
}
