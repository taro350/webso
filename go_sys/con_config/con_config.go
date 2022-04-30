package con_config

import "crypto/rsa"

var (
	PubKeyFile  *rsa.PublicKey
	PrivKeyFile *rsa.PrivateKey
)

const (
	PKPWD = "Some pass"

	KeyCertPath = "var/www/keycertz"
)
