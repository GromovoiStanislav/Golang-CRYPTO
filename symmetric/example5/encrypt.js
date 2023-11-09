const crypto = require('crypto');
const fs = require('fs');

// Генерируем симметричный ключ
const key = crypto.randomBytes(32); // 256 бит

// Записываем ключ в файл
fs.writeFileSync('symmetric_key.txt', key.toString('hex'), 'utf-8');

// Сообщение для шифрования
const plaintext = 'Привет, мир!';

// Генерируем случайный IV (Initialization Vector)
const iv = crypto.randomBytes(16);

// Создаем объект для шифрования
const cipher = crypto.createCipheriv('aes-256-cbc', key, iv);

// Шифруем сообщение
let encrypted = cipher.update(plaintext, 'utf-8', 'hex');
encrypted += cipher.final('hex');

// Записываем IV и зашифрованное сообщение в файл
const encryptedData = `${iv.toString('hex')}\n${encrypted}`;
fs.writeFileSync('encrypted_message.txt', encryptedData, 'utf-8');

console.log('Шифрование завершено.');
