-- Rollback: Reverter hash para senha em texto plano
-- ATENÇÃO: Este rollback é apenas para desenvolvimento/testes
-- NUNCA use em produção!

UPDATE users 
SET password = 'senha_temporaria_123'
WHERE password = '$2a$10$WxNYd/gauUUXb3G7jd0OpOVKUSkMWhrf3Zsomt0KMiSTAiymVv0.y';
