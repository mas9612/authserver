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
	logger     *zap.Logger
	conn       *ldap.Conn
	userFormat string
	pemPath    string
	issuer     string
	audience   string
}

type config struct {
	ldapaddr   string
	ldapport   int
	userFormat string
	pem        string
	issuer     string
	audience   string
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

// SetUserFormat sets the user format used when bind to LDAP server.
func SetUserFormat(format string) Option {
	return func(c *config) {
		c.userFormat = format
	}
}

// SetPem sets the pem path used to sign JWT token.
func SetPem(pem string) Option {
	return func(c *config) {
		c.pem = pem
	}
}

// SetIssuer sets the issuer used in JWT claim.
func SetIssuer(issuer string) Option {
	return func(c *config) {
		c.issuer = issuer
	}
}

// SetAudience sets the audience used in JWT claim.
func SetAudience(audience string) Option {
	return func(c *config) {
		c.audience = audience
	}
}

// NewAuthserver creates new server instance.
func NewAuthserver(logger *zap.Logger, opts ...Option) (pb.AuthserverServer, error) {
	c := config{
		ldapaddr:   "localhost",
		ldapport:   389,
		userFormat: "%s",
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
		logger:     logger,
		conn:       conn,
		userFormat: c.userFormat,
		pemPath:    c.pem,
	}, nil
}

// CreateToken creates and returns the new token.
func (s *Authserver) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.Token, error) {
	if err := s.conn.Bind(fmt.Sprintf(s.userFormat, req.User), req.Password); err != nil {
		return nil, errors.Wrap(err, "bind failed")
	}

	signKeyBytes, err := ioutil.ReadFile(s.pemPath)
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
		Audience:  s.audience,
		ExpiresAt: nowUnix + 3600, // valid 1h
		Id:        v4.String(),
		IssuedAt:  nowUnix,
		Issuer:    s.issuer,
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
