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

func generatePrime(bits int) *big.Int {
	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return prime
}


// truncateSharedSecret усекает общий секрет до нужной длины
func truncateSharedSecret(sharedSecret *big.Int) []byte {
	hash := sha256.New()
	hash.Write(sharedSecret.Bytes())
	return hash.Sum(nil)[:32] // Усечение до 32 байт
}

// encrypt использует AES-GCM для шифрования данных
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

// decrypt использует AES-GCM для расшифровки данных
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
	// Генерация простых чисел и базы для Диффи-Хеллмана
	bits := 2048
	prime := generatePrime(bits)
	base := big.NewInt(2)

	// Генерация ключей Алисы
	alicePrivateKey, err := rand.Int(rand.Reader, prime)
	if err != nil {
		panic(err)
	}
	alicePublicKey := new(big.Int).Exp(base, alicePrivateKey, prime)

	// Получение публичного ключа Алисы в виде строки (hex)
	alicePublicKeyHex := hex.EncodeToString(alicePublicKey.Bytes())

	// Генерация ключей Боба
	bobPrivateKey, err := rand.Int(rand.Reader, prime)
	if err != nil {
		panic(err)
	}
	bobPublicKey := new(big.Int).Exp(base, bobPrivateKey, prime)

	// Получение публичного ключа Боба в виде строки (hex)
	bobPublicKeyHex := hex.EncodeToString(bobPublicKey.Bytes())

	// Вычисление общего секретного ключа Алисы
	sharedSecretAlice := new(big.Int).Exp(bobPublicKey, alicePrivateKey, prime)
	sharedSecretAliceHex := hex.EncodeToString(sharedSecretAlice.Bytes())

	// Вычисление общего секретного ключа Боба
	sharedSecretBob := new(big.Int).Exp(alicePublicKey, bobPrivateKey, prime)
	sharedSecretBobHex := hex.EncodeToString(sharedSecretBob.Bytes())

	// Вывод результата
	fmt.Println("Публичный ключ Алисы:", alicePublicKeyHex)
	fmt.Println("Публичный ключ Боба:", bobPublicKeyHex)
	fmt.Println("Общий секретный ключ у Алисы:", sharedSecretAliceHex)
	fmt.Println("Общий секретный ключ у Боба:", sharedSecretBobHex)
	fmt.Println(sharedSecretAliceHex == sharedSecretBobHex)


	// Сообщение для шифрования
	message := []byte("Hello, Diffie-Hellman!")

	// Преобразование общего секрета в ключ фиксированной длины
	encryptionKeyAlice := truncateSharedSecret(sharedSecretAlice)
	encryptionKeyBob := truncateSharedSecret(sharedSecretBob)

	// Шифрование сообщения с использованием общего секрета Алисы
	ciphertext, err := encrypt(message, encryptionKeyAlice)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	fmt.Printf("Зашифрованное сообщение: %x\n", ciphertext)

	// Расшифровка сообщения с использованием общего секрета Боба
	decryptedMessage, err := decrypt(ciphertext, encryptionKeyBob)
	if err != nil {
		fmt.Println("Ошибка при расшифровке:", err)
		return
	}

	fmt.Printf("Расшифрованное сообщение: %s\n", decryptedMessage)
}
