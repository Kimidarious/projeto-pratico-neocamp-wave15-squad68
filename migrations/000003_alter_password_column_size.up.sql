-- Alterar tamanho da coluna password de VARCHAR(10) para VARCHAR(255)
-- Necessário para armazenar hash bcrypt (60 caracteres)

ALTER TABLE users 
ALTER COLUMN password TYPE VARCHAR(255);

-- Adicionar comentário explicativo
COMMENT ON COLUMN users.password IS 'Senha do usuário armazenada como hash bcrypt (60 chars)';
