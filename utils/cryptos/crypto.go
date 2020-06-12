package cryptos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"github.com/appletouch/bookstore-users_api/logger"
	"golang.org/x/crypto/bcrypt"
	"io"
)

const (
	passphrase = "1234567890"
)

// priate: function will take a passphrase or any string, hash it, then return the hash as a hexadecimal value
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//private
func encrypt(data []byte) []byte {

	//First we create a new block cipher based on the hashed passphrase.
	//Once we have our block cipher, we want to wrap it in Galois Counter Mode (GCM) with a standard nonce length.
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error("APPLICATION CRASHED WITH ERROR", err)
		panic(err.Error())
	}

	//Before we can create the ciphertext, we need to create a nonce.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Error("APPLICATION CRASHED WITH ERROR", err)
		panic(err.Error())
	}
	//There are a few strategies that can be used to make sure our decryption nonce matches the encryption nonce.
	//One strategy would be to store the nonce alongside the encrypted data if it is going into a database.
	//Another option is to prepend or append the nonce to the encrypted data. We’ll prepending the nonce.
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext
}

//private
func decrypt(data []byte) []byte {
	//In the above code we create a new block cipher using a hashed passphrase.
	//We wrap the block cipher in Galois Counter Mode and get the nonce size.
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Error("APPLICATION CRASHED WITH ERROR", err)
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logger.Error("APPLICATION CRASHED WITH ERROR", err)
		panic(err.Error())
	}

	//Remember, we prefixed our encrypted data with the nonce.
	//This means that we need to separate the nonce and the encrypted data.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	//When we have our nonce and ciphertext separated, we can decrypt the data and return it as plaintext.
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logger.Error("APPLICATION CRASHED WITH ERROR", err)
		panic(err.Error())
	}
	return plaintext
}

//Public encryption
func Encryptpassword(password string) string {
	secret := encrypt([]byte(password))
	return string(secret)
}

//Public decryption
func DecryptPassword(secret string) string {
	password := decrypt([]byte(secret))
	return string(password)
}

func HashAndSaltPassword(password string) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logger.Error(err.Error(), err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
