-- Adicionar coluna password à tabela users (se não existir)
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name = 'users' 
        AND column_name = 'password'
    ) THEN
        ALTER TABLE users ADD COLUMN password VARCHAR(255) NOT NULL DEFAULT '';
    END IF;
END $$;

-- Comentário explicativo
COMMENT ON COLUMN users.password IS 'Senha do usuário (armazenada como hash bcrypt)';