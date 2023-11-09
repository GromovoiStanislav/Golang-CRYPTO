const fs = require('fs');
const { execSync } = require('child_process');

// Генерация ключей
const privateKey = execSync('openssl genpkey -algorithm RSA -out private_key.pem');
const publicKey = execSync('openssl rsa -pubout -in private_key.pem -out public_key.pem');

console.log('Закрытый ключ успешно сохранен в файле private_key.pem.');
console.log('Открытый ключ успешно сохранен в файле public_key.pem.');