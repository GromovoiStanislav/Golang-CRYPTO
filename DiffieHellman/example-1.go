package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
)

// Генерация случайного простого числа
func generatePrime(bits int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bits)
}

// Генерация приватного ключа
func generatePrivateKey(bits int) (*big.Int, error) {
	return rand.Int(rand.Reader, new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits)), nil), big.NewInt(1)))
}

// Вычисление открытого ключа на основе приватного ключа и общего простого числа
func calculatePublicKey(privateKey, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), privateKey, p)
}

// Обмен открытыми ключами и вычисление общего секрета
func calculateSharedSecret(privateKey, publicKey, p *big.Int) *big.Int {
	return new(big.Int).Exp(publicKey, privateKey, p)
}

func truncateSharedSecret(sharedSecret *big.Int) []byte {
	hash := sha256.New()
	hash.Write(sharedSecret.Bytes())
	return hash.Sum(nil)[:32] // Усечение до 32 байт
}

// Шифрование данных с использованием AES-GCM
func encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, data, nil)
	return append(nonce, ciphertext...), nil
}

// Расшифровка данных с использованием AES-GCM
func decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	// Биты для генерации простого числа
	bits := 2048

	// Генерация простого числа p
	p, err := generatePrime(bits)
	if err != nil {
		fmt.Println("Ошибка при генерации простого числа:", err)
		return
	}

	// Генерация случайного приватного ключа a
	privateKeyA, err := generatePrivateKey(bits)
	if err != nil {
		fmt.Println("Ошибка при генерации приватного ключа A:", err)
		return
	}

	// Генерация открытого ключа A
	g := int64(2) // базовый генератор
	publicKeyA := calculatePublicKey(privateKeyA, p, g)

	// Генерация случайного приватного ключа b
	privateKeyB, err := generatePrivateKey(bits)
	if err != nil {
		fmt.Println("Ошибка при генерации приватного ключа B:", err)
		return
	}

	// Генерация открытого ключа B
	publicKeyB := calculatePublicKey(privateKeyB, p, g)

	// Обмен открытыми ключами

	// Вычисление общего секрета на стороне A
	sharedSecretA := calculateSharedSecret(privateKeyA, publicKeyB, p)

	// Вычисление общего секрета на стороне B
	sharedSecretB := calculateSharedSecret(privateKeyB, publicKeyA, p)

	// Проверка совпадения общих секретов
	if sharedSecretA.Cmp(sharedSecretB) == 0 {
		fmt.Println("Общий секрет совпадает:", sharedSecretA)
	} else {
		fmt.Println("Общий секрет не совпадает")
	}

	// Использование общего секрета, например, для дальнейшего шифрования
	hash := sha256.New()
	hash.Write(sharedSecretA.Bytes())
	encryptionKey := hash.Sum(nil)

	fmt.Printf("Используйте общий секрет для шифрования: %x\n", encryptionKey)

	// Шифрование и расшифровка с использованием общего секрета

	// Сообщение для шифрования
	message := []byte("Hello, Diffie-Hellman!")

	// Преобразование общего секрета в ключ фиксированной длины
	encryptionKeyA := truncateSharedSecret(sharedSecretA)
	encryptionKeyB := truncateSharedSecret(sharedSecretB)

	// Шифрование сообщения с использованием общего секрета
	ciphertext, err := encrypt(message, encryptionKeyA)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	fmt.Printf("Зашифрованное сообщение: %x\n", ciphertext)

	// Расшифровка сообщения с использованием общего секрета
	decryptedMessage, err := decrypt(ciphertext, encryptionKeyB)
	if err != nil {
		fmt.Println("Ошибка при расшифровке:", err)
		return
	}

	fmt.Printf("Расшифрованное сообщение: %s\n", decryptedMessage)

}
