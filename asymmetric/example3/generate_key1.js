const fs = require('fs');
const crypto = require('crypto');

// Генерация ключей
const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
  modulusLength: 2048,
});

// Экспорт закрытого ключа в формат PKCS#8
const privateKeyPem = privateKey.export({
  type: 'pkcs8',
  format: 'pem',
});

// Сохранение закрытого ключа в файл
fs.writeFileSync('private_key.pem', privateKeyPem);

console.log('Закрытый ключ успешно сохранен в файле private_key.pem.');

// Экспорт открытого ключа в формат PKCS#8
const publicKeyPem = publicKey.export({
  type: 'spki',
  format: 'pem',
});

// Сохранение открытого ключа в файл
fs.writeFileSync('public_key.pem', publicKeyPem);

console.log('Открытый ключ успешно сохранен в файле public_key.pem.');
