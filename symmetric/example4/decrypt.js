const fs = require('fs');
const crypto = require('crypto');

// Считывание симметричного ключа из файла
const key = fs.readFileSync('symmetric_key.bin');

// Считывание зашифрованного сообщения из файла
const encryptedMessageWithAuthTag = fs.readFileSync('encrypted_message.bin');

// Отделяем аутентификационный тег от зашифрованных данных
const authTagLength = 16; // Длина аутентификационного тега для AES-GCM
const authTag = encryptedMessageWithAuthTag.slice(-authTagLength);
const encryptedMessage = encryptedMessageWithAuthTag.slice(0, -authTagLength);

// Используем первые 12 байт из зашифрованного сообщения как IV
const iv = encryptedMessage.slice(0, 12);

// Оставшиеся байты после первых 12 - это сами зашифрованные данные
const encryptedData = encryptedMessage.slice(12);

// Создаем объект decipher для дешифрования симметричным ключом AES
const decipher = crypto.createDecipheriv('aes-256-gcm', key, iv);

// Устанавливаем аутентификационный тег
decipher.setAuthTag(authTag);

// Обновим объект decipher с зашифрованными данными
let decryptedMessage;

try {
  decryptedMessage = decipher.update(encryptedData, 'binary', 'utf8');
  decryptedMessage += decipher.final('utf8');
} catch (error) {
  console.error('Ошибка дешифрования:', error.message);
  return;
}

console.log('Дешифрованное сообщение:', decryptedMessage);
