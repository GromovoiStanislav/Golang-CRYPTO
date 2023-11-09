const crypto = require('crypto');
const fs = require('fs');

// Считывание открытого ключа из файла
const publicKey = fs.readFileSync('public_key.pem', 'utf8');

// Сообщение, которое мы хотим зашифровать
const message = 'Это сообщение, которое мы хотим зашифровать.';

// Шифрование сообщения с использованием открытого ключа
const encryptedBuffer = crypto.publicEncrypt(
  {
    key: publicKey,
    padding: crypto.constants.RSA_PKCS1_PADDING,
  },
  Buffer.from(message, 'utf8')
);

// Сохранение зашифрованного сообщения в файл
fs.writeFileSync('encrypted_message.bin', encryptedBuffer, 'utf8');

console.log('Сообщение успешно зашифровано и сохранено в файле.');
