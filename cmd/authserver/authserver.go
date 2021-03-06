package main

import (
	"fmt"
	"net"

	"github.com/jessevdk/go-flags"
	pb "github.com/mas9612/authserver/pkg/authserver"
	"github.com/mas9612/authserver/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type options struct {
	Port       int    `short:"p" long:"port" default:"10000" description:"listen port."`
	LdapAddr   string `long:"ldap-addr" default:"localhost" description:"LDAP server address."`
	LdapPort   int    `long:"ldap-port" default:"389" description:"LDAP server port."`
	UserFormat string `short:"f" long:"user-format" default:"%s" description:"User format used when bind to LDAP server."`
	PemPath    string `long:"pem" default:"signkey.pem" description:"PEM encoded key used to sign JWT token.pem)."`
	PubkeyPath string `long:"pubkey" default:"authserver.pub" description:"Public key which is used to validate JWT token."`
	Issuer     string `long:"issuer" default:"authserver" description:"Issuer name used in JWT claim."`
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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		logger.Fatal("listen failed", zap.Error(err))
	}
	defer listener.Close()

	options := make([]server.Option, 0, 10)
	if opts.LdapAddr != "localhost" {
		options = append(options, server.SetAddr(opts.LdapAddr))
	}
	if opts.LdapPort != 389 {
		options = append(options, server.SetPort(opts.LdapPort))
	}
	if opts.UserFormat != "%s" {
		options = append(options, server.SetUserFormat(opts.UserFormat))
	}
	if opts.PemPath != "signkey.pem" {
		options = append(options, server.SetPem(opts.PemPath))
	}
	if opts.PubkeyPath != "authserver.pub" {
		options = append(options, server.SetPubkey(opts.PubkeyPath))
	}
	if opts.Issuer != "authserver" {
		options = append(options, server.SetIssuer(opts.Issuer))
	}

	s, err := server.NewAuthserver(logger, options...)
	if err != nil {
		logger.Fatal("failed to initialize server", zap.Error(err))
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthserverServer(grpcServer, s)
	logger.Info(fmt.Sprintf("listening on :%d", opts.Port))
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("service failed", zap.Error(err))
	}
}
