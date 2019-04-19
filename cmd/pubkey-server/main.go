package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
)

type options struct {
	Addr    string `short:"a" long:"address" default:"" description:"Listen address (default is empty string)."`
	Port    int    `short:"p" long:"port" default:"80" description:"Listen port (default is 8080)."`
	KeyPath string `long:"key-path" default:"keys" description:"Directory path which holds public keys."`
}

// Key represents one key.
type Key struct {
	Ktype     string `json:"kty"`
	Use       string `json:"use"`
	Algorithm string `json:"alg"`
	Kid       string `json:"kid"`
	N         string `json:"n"`
	E         string `json:"e"`
}

// KeySet represents the whole key sets which will be returned by pubkey-server.
type KeySet struct {
	Keys []Key `json:"keys"`
}

var (
	keySet KeySet
	logger *zap.Logger
)

func main() {
	logger, _ = zap.NewProduction()
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

	files, err := ioutil.ReadDir(opts.KeyPath)
	if err != nil {
		logger.Fatal("failed to read directory", zap.Error(err))
	}
	keySet.Keys = make([]Key, 0, 10)
	hasher := sha256.New()
	b64encoder := base64.URLEncoding.WithPadding(base64.NoPadding)
	for _, f := range files {
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", opts.KeyPath, f.Name()))
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse key file '%s'", f.Name()), zap.Error(err))
			continue
		}
		pubkey, err := jwt.ParseRSAPublicKeyFromPEM(b)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse key file '%s'", f.Name()), zap.Error(err))
			continue
		}
		hasher.Reset()
		hasher.Write(b)
		keySet.Keys = append(keySet.Keys, Key{
			Ktype:     "RSA",
			Use:       "sig",
			Algorithm: "RS256",
			Kid:       fmt.Sprintf("%x", hasher.Sum(nil)),
			N:         b64encoder.EncodeToString(pubkey.N.Bytes()),
			E:         b64encoder.EncodeToString(intToBytes(pubkey.E)),
		})
	}
	if len(keySet.Keys) == 0 {
		logger.Fatal("no key is found")
	}

	r := mux.NewRouter()
	r.HandleFunc("/keys", keysHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", opts.Addr, opts.Port), r)
}

func intToBytes(n int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(n))
	for i := 0; i < 4; i++ {
		if b[i] != 0x0 {
			b = b[i:]
			break
		}
	}
	return b
}

func keysHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(keySet); err != nil {
		logger.Error("failed to write JSON", zap.Error(err))
	}
}
