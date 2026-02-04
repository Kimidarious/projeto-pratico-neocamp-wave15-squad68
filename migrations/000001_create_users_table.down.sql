-- Rollback: remover tabela users
DROP INDEX IF EXISTS idx_users_user_type;
DROP INDEX IF EXISTS idx_users_user_name;
DROP TABLE IF EXISTS users;