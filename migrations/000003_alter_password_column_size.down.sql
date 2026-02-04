-- Rollback: Reverter tamanho da coluna password
-- ATENÇÃO: Pode causar perda de dados se houver senhas > 10 caracteres

ALTER TABLE users 
ALTER COLUMN password TYPE VARCHAR(10);
