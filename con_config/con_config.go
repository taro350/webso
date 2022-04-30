package con_config

import "crypto/rsa"

var (
	PubKeyFile  *rsa.PublicKey
	PrivKeyFile *rsa.PrivateKey
)

const (
	PKPWD = "Some pass"

	KeyCertPath  = "var/www/keycertz"
	PubCertPath  = "var/www/keycertz"
	PrivCertPath = "var/www/keycertz"

	MongoHost     = "127.0.0.1"
	MongoUser     = "mongod"
	MongoPassword = "some pass"
	MongoDb       = "admin"
)

// func init() {
// 	f, ok, err :=
// }
