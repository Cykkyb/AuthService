-- Изменение структуры таблицы users
ALTER TABLE users
RENAME COLUMN password TO password_hash;