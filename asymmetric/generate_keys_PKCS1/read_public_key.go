package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	// Читаем публичный ключ из файла
	publicKeyBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		fmt.Println("Ошибка при чтении публичного ключа:", err)
		return
	}

	// Декодируем публичный ключ из формата PEM
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		fmt.Println("Не удалось декодировать публичный ключ")
		return
	}

	// Преобразование публичного ключа
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Ошибка при преобразовании публичного ключа:", err)
		return
	}

	// Выводим публичный ключ
	fmt.Println(publicKey)
}
