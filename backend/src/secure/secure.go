package secure

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func VerifyPass(hashPass, stringPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(stringPass))
}
