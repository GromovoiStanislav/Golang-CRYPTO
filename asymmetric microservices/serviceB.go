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
	saveKeyToFile("serviceB_public.pem", publicKey)

	// Сохранение закрытого ключа микросервиса A
	saveKeyToFile("serviceB_private.pem", privateKey)
}

func sendMessage() {
	// Загрузка открытого ключа микросервиса B
	publicKeyA, err := loadPublicKey("serviceA_public.pem")
	if err != nil {
		fmt.Println("Ошибка при загрузке открытого ключа микросервиса B:", err)
		return
	}

	// Шифрование сообщения с открытым ключом микросервиса B
	message := "Привет, микросервис A!"
	encryptedMessage, err := encryptMessage(publicKeyA, []byte(message))
	if err != nil {
		fmt.Println("Ошибка при шифровании сообщения:", err)
		return
	}

	// fmt.Println("Зашифрованное сообщение для микросервиса A:")
	// fmt.Println(string(encryptedMessage))

	// Сохранение зашифрованного сообщения
	err = ioutil.WriteFile("messageB", encryptedMessage, 0644)
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
	// Загрузка зашифрованного сообщения в формате Base64
	encryptedMessageBase64, err := ioutil.ReadFile("messageA")
	if err != nil {
		fmt.Println("Ошибка при загрузке зашифрованного сообщения:", err)
		return
	}

	// Декодирование сообщения из Base64 в байты
	encryptedMessage, err := base64.StdEncoding.DecodeString(string(encryptedMessageBase64))
	if err != nil {
		fmt.Println("Ошибка при декодировании сообщения из Base64:", err)
		return
	}

	// Дешифрование сообщения от микросервиса B
	privateKeyA, err := loadPrivateKey("serviceB_private.pem")
	if err != nil {
		fmt.Println("Ошибка при загрузке закрытого ключа микросервиса A:", err)
		return
	}

	decryptedMessage, err := decryptMessage(privateKeyA, encryptedMessage)
	if err != nil {
		fmt.Println("Ошибка при дешифровании сообщения:", err)
		return
	}

	fmt.Println("Дешифрованное сообщение от микросервиса A:")
	fmt.Println(string(decryptedMessage))
}


func main() {
	//createKeys()
	//sendMessage()
	readeMessage()
}