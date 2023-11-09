const fs = require('fs');
const crypto = require('crypto');

// Генерация случайного симметричного ключа длиной 32 байта (256 бит)
const symmetricKey = crypto.randomBytes(32);

// Запись ключа в файл
fs.writeFileSync('symmetric_key.bin', symmetricKey);

console.log('Симметричный ключ успешно сгенерирован и записан в файл.');
