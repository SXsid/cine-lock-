package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	root "github.com/SXsid/cine-lock"
)

func main() {
	port := flag.Int("port", 8080, "HTPP server port")
	flag.Parse()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: http.FileServerFS(root.StaticAssests),
	}
	go func(server *http.Server) {
		fmt.Printf("server is up and running at %s \n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}(&server)
	signal, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()
	<-signal.Done()
	fmt.Println("shutting donw gracefully")
	timeContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(timeContext); err != nil {
		fmt.Println("froced shutdown occured")
		panic(err)
	}
}
