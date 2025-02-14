wwpackage main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func main() {
	key := "<INSERT_KEY>"
	iv := "<INSERT_IV>"
	encryptedText := "<INSERT_TEXT_TO_DECRYPT>"

	ivBytes, err := hex.DecodeString(iv)
	if err != nil {
		log.Fatalf("Error por decoding del IV: %v", err)
	}

	cipherTextBytes, err := hex.DecodeString(encryptedText)
	if err != nil {
		log.Fatalf("Error por decoding del texto cifrado: %v", err)
	}

	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatalf("Error al crear el bloque cifrado: %v", err)
	}

	if len(cipherTextBytes)%aes.BlockSize != 0 {
		log.Fatal("El texto cifrado no es múltiplo del tamaño del bloque")
	}

	mode := cipher.NewCBCDecrypter(block, ivBytes)

	plainText := make([]byte, len(cipherTextBytes))
	mode.CryptBlocks(plainText, cipherTextBytes)

	plainText, err = pkcs7Unpad(plainText)
	if err != nil {
		log.Fatalf("Error por sin padding: %v", err)
	}

	fmt.Printf("Texto decifrado: %s\n", plainText)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("Error por datos vacíos")
	}

	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		return nil, errors.New("Tamaño de padding inválido")
	}

	for i := 0; i < padding; i++ {
		if data[len(data)-1-i] != byte(padding) {
			return nil, errors.New("Padding inválido")
		}
	}

	return data[:len(data)-padding], nil
}