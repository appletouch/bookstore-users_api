package cryptos

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

//import "fmt"
//
//func main() {
//	fmt.Println("Starting the application...")
//	ciphertext := encrypt([]byte("Hello World"), "password")
//	fmt.Printf("Encrypted: %x\n", ciphertext)
//	plaintext := decrypt(ciphertext, "password")
//	fmt.Printf("Decrypted: %s\n", plaintext)
//	encryptFile("sample.txt", []byte("Hello World"), "password1")
//	fmt.Println(string(decryptFile("sample.txt", "password1")))
//}

func TestEncyptionStrings(t *testing.T) {
	fmt.Println("Starting the testing encryption with strings...")
	startPassword := "Peter01"
	fmt.Println("Password before test:", startPassword)

	//encrypt
	ciphertext := Encryptpassword(startPassword)
	fmt.Printf("Encrypted: %x\n", ciphertext)

	//decrypt
	password := DecryptPassword(ciphertext)

	fmt.Printf("Password after test: %s\n", password)

}

func TestEncyptionBytes(t *testing.T) {
	fmt.Println("Starting the testing encryption bytes...")
	//encrypt raw
	password := "Hello World"
	fmt.Printf("password before: %s\n", password)
	ciphertext := encrypt([]byte(password))

	str := string(ciphertext)
	bt := []byte(str)

	//decrypt raw
	plaintext := decrypt(bt)
	fmt.Printf("password after: %s\n", plaintext)
}

func TestHashAndSaltPassword(t *testing.T) {
	var startPassword string = "welkom01"
	fmt.Printf("The start password: %s\n", startPassword)
	secret := HashAndSaltPassword(startPassword)
	fmt.Printf("The start secret: %s\n", secret)

	err := bcrypt.CompareHashAndPassword([]byte(secret), []byte(startPassword))
	if err != nil {
		log.Println(err)
		t.Error(err)
	} else {
		t.Log("Test was succesful for:", startPassword)
	}
}
