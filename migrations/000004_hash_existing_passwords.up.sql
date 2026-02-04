-- Migration para hashear senhas existentes em texto plano
-- Esta migration detecta senhas que NÃO são hash bcrypt e as converte
-- Hash bcrypt sempre começa com $2a$, $2b$, $2x$, $2y$

-- Criar função temporária para identificar se é hash bcrypt
CREATE OR REPLACE FUNCTION is_bcrypt_hash(password_text TEXT) 
RETURNS BOOLEAN AS $$
BEGIN
    -- Hash bcrypt tem pelo menos 60 caracteres e começa com $2
    RETURN password_text ~ '^\\$2[abxy]\\$[0-9]{2}\\$.{53,}$';
END;
$$ LANGUAGE plpgsql;

-- Atualizar senhas que NÃO são hash bcrypt
-- Usa hash de "senha_temporaria_123" como padrão
UPDATE users 
SET password = '$2a$10$WxNYd/gauUUXb3G7jd0OpOVKUSkMWhrf3Zsomt0KMiSTAiymVv0.y'
WHERE NOT is_bcrypt_hash(password);

-- Remover função temporária
DROP FUNCTION IF EXISTS is_bcrypt_hash(TEXT);

-- Log: Mensagem informativa
DO $$ 
DECLARE
    updated_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO updated_count 
    FROM users 
    WHERE password = '$2a$10$WxNYd/gauUUXb3G7jd0OpOVKUSkMWhrf3Zsomt0KMiSTAiymVv0.y';
    
    RAISE NOTICE '✅ % user(s) passwords have been hashed. Default password: senha_temporaria_123', updated_count;
END $$;
