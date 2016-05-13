package encryption



import (
	"api-uptoo-jwt/settings"
	"crypto/cipher"
	"log"
	"crypto/aes"
	"io"
	"crypto/rand"
	"encoding/base64"
)

func Encrypter(PlainText string) string {
	key := []byte(settings.Get().AESKEY)
	plaintext := []byte(PlainText)


	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err.Error())
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err.Error())
	}

	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypter(cryptoText string) string {
	key := []byte(settings.Get().AESKEY)
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext trop petit")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}
