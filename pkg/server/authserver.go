package server

import (
	context "context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/mas9612/authserver/pkg/authserver"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gopkg.in/ldap.v3"
)

// Authserver is the implementation of pb.AuthserverServer.
type Authserver struct {
	logger *zap.Logger
	conn   *ldap.Conn
}

type config struct {
	ldapaddr string
	ldapport int
}

// Option is the option to create authserver instance.
type Option func(*config)

// SetAddr sets the LDAP server address.
func SetAddr(addr string) Option {
	return func(c *config) {
		c.ldapaddr = addr
	}
}

// SetPort sets the LDAP server port.
func SetPort(port int) Option {
	return func(c *config) {
		c.ldapport = port
	}
}

// NewAuthserver creates new server instance.
func NewAuthserver(logger *zap.Logger, opts ...Option) (pb.AuthserverServer, error) {
	c := config{
		ldapaddr: "localhost",
		ldapport: 389,
	}
	for _, o := range opts {
		o(&c)
	}

	logger.Info("connecting LDAP server")
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", c.ldapaddr, c.ldapport))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to LDAP server")
	}
	logger.Info("connected")

	return &Authserver{
		logger: logger,
		conn:   conn,
	}, nil
}

// CreateToken creates and returns the new token.
func (s *Authserver) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.Token, error) {
	if err := s.conn.Bind(fmt.Sprintf("uid=%s,ou=Users,dc=ldap,dc=firefly,dc=kutc,dc=kansai-u,dc=ac,dc=jp", req.User), req.Password); err != nil {
		return nil, errors.Wrap(err, "bind failed")
	}

	signKeyBytes, err := ioutil.ReadFile("./test.pem")
	if err != nil {
		return nil, errors.Wrap(err, "failed to load signkey")
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse signkey")
	}

	nowUnix := time.Now().Unix()
	v4 := uuid.NewV4()
	claims := jwt.StandardClaims{
		Audience:  "firefly.kutc.kansai-u.ac.jp",
		ExpiresAt: nowUnix + 3600, // valid 1h
		Id:        v4.String(),
		IssuedAt:  nowUnix,
		Issuer:    "authserver.firefly.kutc.kansai-u.ac.jp",
		NotBefore: nowUnix - 5,
		Subject:   "access_token",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate JWT token")
	}

	return &pb.Token{
		Token: ss,
	}, nil
}
