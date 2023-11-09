const crypto = require('crypto');
const fs = require('fs');

// Читаем ключ из файла
const keyHex = fs.readFileSync('symmetric_key.txt', 'utf-8');
const key = Buffer.from(keyHex, 'hex');

// Читаем зашифрованные данные из файла
const encryptedData = fs
  .readFileSync('encrypted_message.txt', 'utf-8')
  .split('\n');

// Выделяем IV и зашифрованное сообщение
const iv = Buffer.from(encryptedData[0], 'hex');
const encryptedMessage = encryptedData[1];

// Создаем объект для дешифрования
const decipher = crypto.createDecipheriv('aes-256-cbc', key, iv);

// Дешифруем сообщение
let decrypted = decipher.update(encryptedMessage, 'hex', 'utf-8');
decrypted += decipher.final('utf-8');

console.log('Дешифрованное сообщение:', decrypted);
