package helper

import "golang.org/x/crypto/bcrypt"

// Check password from hashed password with the plain
// password. Returning nil when equal, failure when not equal
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
