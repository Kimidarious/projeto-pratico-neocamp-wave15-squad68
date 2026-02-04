-- Rollback: remover coluna password
ALTER TABLE users 
DROP COLUMN IF EXISTS password;