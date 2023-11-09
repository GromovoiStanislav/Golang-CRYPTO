const crypto = require('crypto');
const fs = require('fs');

// Считывание закрытого ключа из файла
const privateKey = fs.readFileSync('private_key.pem', 'utf8');

// Считывание зашифрованного сообщения из файла
const encryptedMessage = fs.readFileSync('encrypted_message.bin');

// Расшифровка сообщения с использованием приватного ключа
const decryptedBuffer = crypto.privateDecrypt(
  {
    key: privateKey,
    padding: crypto.constants.RSA_PKCS1_PADDING,
  },
  encryptedMessage
);

const decryptedMessage = decryptedBuffer.toString('utf8');
console.log('Расшифрованное сообщение:', decryptedMessage);
