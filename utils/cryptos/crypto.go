package cryptos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
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
		panic(err.Error())
	}

	//Before we can create the ciphertext, we need to create a nonce.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
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
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Remember, we prefixed our encrypted data with the nonce.
	//This means that we need to separate the nonce and the encrypted data.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	//When we have our nonce and ciphertext separated, we can decrypt the data and return it as plaintext.
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
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
