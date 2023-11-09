const crypto = require('crypto');
const fs = require('fs');

// Генерация пары асимметричных ключей
const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
  modulusLength: 2048,
});

// Сохранение закрытого ключа в файл
const privateKeyPem = privateKey.export({
  type: 'pkcs1',
  format: 'pem',
});
fs.writeFileSync('private_key.pem', privateKeyPem);

// Сохранение открытого ключа в файл
const publicKeyPem = publicKey.export({
  type: 'pkcs1',
  format: 'pem',
});
fs.writeFileSync('public_key.pem', publicKeyPem);

console.log('Ключи успешно созданы и сохранены в файлах.');
