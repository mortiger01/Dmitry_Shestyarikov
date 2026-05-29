package utils

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, salt string) (string, error) {
	salted := password + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, salt, hash string) bool {
	salted := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salted))
	return err == nil
}

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashToken(token string) (string, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return "", err
	}
	salted := token + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return salt + ":" + string(bytes), nil
}

func CheckTokenHash(token, tokenHash string) bool {
	parts := splitN(tokenHash, ":", 2)
	if len(parts) != 2 {
		return false
	}
	salt, hash := parts[0], parts[1]
	salted := token + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salted))
	return err == nil
}

func splitN(s, sep string, n int) []string {
	result := []string{}
	for i := 0; i < n-1; i++ {
		idx := indexOf(s, sep)
		if idx < 0 {
			break
		}
		result = append(result, s[:idx])
		s = s[idx+len(sep):]
	}
	result = append(result, s)
	return result
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}