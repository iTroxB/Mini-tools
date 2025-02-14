package main

import (
	"bytes"
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
	encryptedText := "<INSERT_TEXT_TO_ENCRYPT>"

	ivBytes, err := hex.DecodeString(iv)
	if err != nil {
		log.Fatalf("Error por decoding del IV: %v", err)
	}

	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatalf("Error al crear el bloque cifrado: %v", err)
	}

	paddedPlainText, err := pkcs7Pad([]byte(plainText), aes.BlockSize)
	if err != nil {
		log.Fatalf("Error de padding en el texto: %v", err)
	}

	cipherText := make([]byte, len(paddedPlainText))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, paddedPlainText)

	fmt.Printf("Texto cifrado: %s\n", hex.EncodeToString(cipherText))
}

func pkcs7Pad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 || blockSize > 255 {
		return nil, errors.New("Tamaño de bloque inválido")
	}
	padding := blockSize - len(data)%blockSize
	padded := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
	return padded, nil
}
