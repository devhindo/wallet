package main

import (
	"os"
	"log"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"io"
)


func main() {
	// encrypt "encrypt" dir
	encryptDir("encrypt")
}

func encryptDir(dir string) {
	files := readFiles(dir)
	for _, file := range files {
		content, err := os.ReadFile("./encrypt/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		createBlockCipherAlgorithm(content)
	}
}

func readFiles(dir string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func generateAESKey() {
	key := make([]byte, 32)
    _, err := rand.Read(key)
    if err != nil {
        log.Fatal(err)
    }

    err = os.WriteFile("key", key, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func createBlockCipherAlgorithm(file []byte) {
	generateAESKey()
	key, err := os.ReadFile("key")
	if err != nil {
		log.Fatalf("failed to read key: %s", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("failed to create cipher: %s", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("failed to create gcm: %s", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
 	}

	cipherText := gcm.Seal(nonce, nonce, file, nil)

	err = os.WriteFile("encrypt/ciphertext.bin", cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}

