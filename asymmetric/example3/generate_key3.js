const fs = require('fs');
const NodeRSA = require('node-rsa');

// Генерация ключей
const key = new NodeRSA({ b: 2048 });

// Экспорт закрытого ключа в формат PKCS#8
const privateKeyPem = key.exportKey('pkcs8-private');

// Сохранение закрытого ключа в файл
fs.writeFileSync('private_key.pem', privateKeyPem);

console.log('Закрытый ключ успешно сохранен в файле private_key.pem.');

// Экспорт открытого ключа в формат PKCS#8
const publicKeyPem = key.exportKey('pkcs8-public');

// Сохранение открытого ключа в файл
fs.writeFileSync('public_key.pem', publicKeyPem);

console.log('Открытый ключ успешно сохранен в файле public_key.pem.');
