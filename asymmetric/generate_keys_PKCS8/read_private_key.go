package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
    // Читаем закрытый ключ из файла
    privateKeyBytes, err := ioutil.ReadFile("private.pem")
    if err != nil {
		fmt.Println("Ошибка при чтении закрытого ключа:", err)
		return
    }

    // Декодируем закрытый ключ из формата PEM
    block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		fmt.Println("Не удалось декодировать закрытый ключ")
		return
	}

	// Преобразование закрытого ключа
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Ошибка при преобразовании закрытого ключа:", err)
		return
	}

	// Проверка, является ли закрытый ключ типом *rsa.PrivateKey
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Ошибка при приведении закрытого ключа к типу *rsa.PrivateKey")
		return
	}

    // Выводим закрытый ключ
    fmt.Println(rsaPrivateKey)
}