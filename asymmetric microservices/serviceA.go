package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// Генерация пары ключей
func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // Размер ключа 2048 бит
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// Сохранение ключа в файл
func saveKeyToFile(filename string, key interface{}) error {
	var keyPEM *pem.Block

	switch k := key.(type) {
	case *rsa.PrivateKey:
		keyBytes := x509.MarshalPKCS1PrivateKey(k)
		keyPEM = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: keyBytes,
		}
	case *rsa.PublicKey:
		keyBytes := x509.MarshalPKCS1PublicKey(k)
		keyPEM = &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: keyBytes,
		}
	default:
		return fmt.Errorf("неверный тип ключа")
	}

	err := ioutil.WriteFile(filename, pem.EncodeToMemory(keyPEM), 0644)
	return err
}


func createKeys() {
	// Генерация пары ключей для микросервиса A
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		fmt.Println("Ошибка при генерации ключей:", err)
		return
	}

	// Сохранение открытого ключа микросервиса A
	saveKeyToFile("serviceA_public.pem", publicKey)

	// Сохранение закрытого ключа микросервиса A
	saveKeyToFile("serviceA_private.pem", privateKey)
}

func sendMessage() {
	// Загрузка открытого ключа микросервиса B
	publicKeyB, err := loadPublicKey("serviceB_public.pem")
	if err != nil {
		fmt.Println("Ошибка при загрузке открытого ключа микросервиса B:", err)
		return
	}

	// Шифрование сообщения с открытым ключом микросервиса B
	message := "Привет, микросервис B!"
	encryptedMessage, err := encryptMessage(publicKeyB, []byte(message))
	if err != nil {
		fmt.Println("Ошибка при шифровании сообщения:", err)
		return
	}

	// Преобразование зашифрованного сообщения в Base64
	encryptedMessageBase64 := base64.StdEncoding.EncodeToString(encryptedMessage)

	// fmt.Println("Зашифрованное сообщение для микросервиса B:")
	// fmt.Println(encryptedMessageBase64)

	// Сохранение зашифрованного сообщения
	err = ioutil.WriteFile("messageA", []byte(encryptedMessageBase64), 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении зашифрованного сообщения:", err)
		return
	}
}

// Загрузка открытого ключа из файла
func loadPublicKey(filename string) (*rsa.PublicKey, error) {
	keyPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyPEM)
	if block == nil {
		return nil, fmt.Errorf("неверный формат PEM-блока")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

// Загрузка закрытого ключа из файла
func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	keyPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyPEM)
	if block == nil {
		return nil, fmt.Errorf("неверный формат PEM-блока")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// Дешифрование сообщения с использованием закрытого ключа
func decryptMessage(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	return decryptedMessage, err
}


// Шифрование сообщения с использованием открытого ключа
func encryptMessage(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	return encryptedMessage, err
}

func readeMessage() {
	// Загрузка зашифрованного сообщения
	encryptedMessage, err := ioutil.ReadFile("messageB")
	if err != nil {
		fmt.Println("Ошибка при загрузке зашифрованного сообщения:", err)
		return
	}

	// Дешифрование сообщения от микросервиса B
	privateKeyA, err := loadPrivateKey("serviceA_private.pem")
	if err != nil {
		fmt.Println("Ошибка при загрузке закрытого ключа микросервиса A:", err)
		return
	}

	decryptedMessage, err := decryptMessage(privateKeyA, encryptedMessage)
	if err != nil {
		fmt.Println("Ошибка при дешифровании сообщения:", err)
		return
	}

	fmt.Println("Дешифрованное сообщение от микросервиса B:")
	fmt.Println(string(decryptedMessage))
}

func main() {
	//createKeys()
	//sendMessage()
	readeMessage()
}