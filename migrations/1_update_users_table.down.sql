-- Отмена изменений структуры таблицы users
ALTER TABLE users
RENAME COLUMN password_hash TO password;