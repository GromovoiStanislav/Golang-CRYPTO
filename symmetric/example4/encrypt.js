const fs = require('fs');
const crypto = require('crypto');

// Считывание симметричного ключа из файла
const key = fs.readFileSync('symmetric_key.bin');

// Создание случайного вектора инициализации (IV) длиной 12 байт
const iv = crypto.randomBytes(12);

// Сообщение, которое нужно зашифровать
const message = 'Это секретное сообщение, которое мы хотим зашифровать.';

// Создание объекта cipher для шифрования симметричным ключом AES
const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);

// Шифрование сообщения
const encryptedMessage = Buffer.concat([
  cipher.update(message, 'utf8'),
  cipher.final(),
]);

// Получение аутентификационного тега
const authTag = cipher.getAuthTag();

// Комбинирование IV, зашифрованных данных и аутентификационного тега
const encryptedMessageWithAuthTag = Buffer.concat([
  iv,
  encryptedMessage,
  authTag,
]);

// Запись зашифрованного сообщения в файл
fs.writeFileSync('encrypted_message.bin', encryptedMessageWithAuthTag);

console.log('Сообщение успешно зашифровано и записано в файл.');
