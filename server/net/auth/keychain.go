package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/brieefly/server/config"
)

// PrivateKey - private rsa key for jwt
func PrivateKey(c *config.Config) (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s%s", config.ConfigFilePath, c.Auth.Private))
	if err != nil {
		return nil, err
	}

	data, _ := pem.Decode(bytes)
	if err != nil {
		return nil, err
	}

	rsaPriv, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	return rsaPriv, nil
}

// PublicKey - public rsa key for jwt
func PublicKey(c *config.Config) (*rsa.PublicKey, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s%s", config.ConfigFilePath, c.Auth.Public))
	if err != nil {
		return nil, err
	}

	data, _ := pem.Decode(bytes)
	if err != nil {
		return nil, err
	}

	pkey, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pkey.(*rsa.PublicKey)
	if !ok {
		return nil, err
	}

	return rsaPub, nil
}
