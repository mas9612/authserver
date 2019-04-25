package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jessevdk/go-flags"
	gw "github.com/mas9612/authserver/pkg/authserver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type options struct {
	Port           int    `short:"p" long:"port" default:"8080" description:"Proxy port (default is 8080)."`
	AuthserverAddr string `long:"authserver-addr" default:"" description:"Authserver address (default is empty string)"`
	AuthserverPort int    `long:"authserver-port" default:"10000" description:"Authserver port (default is 10000)."`
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	opts := options{}
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	if _, err := parser.Parse(); err != nil {
		flagsErr := err.(*flags.Error)
		if flagsErr.Type == flags.ErrHelp {
			fmt.Printf("%s\n", flagsErr.Message)
			return
		}
		logger.Fatal("failed to parse command line flags", zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}
	if err := gw.RegisterAuthserverHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", opts.AuthserverAddr, opts.AuthserverPort), grpcOpts); err != nil {
		logger.Fatal("failed to register authserver handler", zap.Error(err))
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), mux); err != nil {
		logger.Fatal("error occured while serving http proxy", zap.Error(err))
	}
}
