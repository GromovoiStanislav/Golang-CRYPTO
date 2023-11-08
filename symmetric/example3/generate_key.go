package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	// Генерация случайного симметричного ключа
	key := make([]byte, 32) // 32 байта для AES-256
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Ошибка при генерации ключа:", err)
		return
	}

	// Сохранение ключа в файл
	err := writeKeyToFile("symmetric_key.bin", key)
	if err != nil {
		fmt.Println("Ошибка при записи ключа в файл:", err)
		return
	}

	fmt.Println("Симметричный ключ сгенерирован и сохранен в symmetric_key.bin")
}

// Запись ключа в файл
func writeKeyToFile(filename string, key []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(key)
	return err
}
