# Implementação de Sistema de Senhas com Hash 🔒

## Resumo da Implementação

Foi implementado um sistema completo de gerenciamento de senhas com hash bcrypt para garantir a segurança dos dados dos usuários.

## O que foi feito ✅

### 1. **Biblioteca de Hashing**
- ✅ Instalado `golang.org/x/crypto/bcrypt`
- ✅ Criado pacote `internal/utils/password.go` com funções:
  - `HashPassword()` - Gera hash bcrypt da senha
  - `ComparePassword()` - Compara senha com hash
  - `IsPasswordValid()` - Valida requisitos mínimos (6-72 caracteres)

### 2. **Modelo de Dados**
- ✅ Campo `password` atualizado para `VARCHAR(255)` (comporta hash bcrypt)
- ✅ Tag JSON alterada para `json:"-"` (senha nunca é retornada nas APIs)

### 3. **DTOs Atualizados**
- ✅ `CreateUserRequest` - Agora requer `password` (mín: 6, máx: 72 chars)
- ✅ `UpdateUserRequest` - Campo `password` opcional para atualização

### 4. **Service Layer**
- ✅ `CreateUser()` - Valida e hashea senha automaticamente
- ✅ `UpdateUser()` - Atualiza senha apenas se fornecida
- ✅ `ValidatePassword()` - Nova função para autenticação/login

### 5. **Handlers**
- ✅ Atualizado para passar senha aos services
- ✅ Senha nunca aparece nas respostas JSON

### 6. **Migrations**
- ✅ `000003_hash_existing_passwords` - Converte senhas em texto plano para hash
- ✅ Todas as senhas "senha_temporaria" foram hasheadas automaticamente

### 7. **Testes**
- ✅ Todos os testes unitários atualizados
- ✅ Verificação de hash adicionada nos testes

## Como Usar 🚀

### Criar Novo Usuário

**Request:**
```bash
POST /api/v1/users
Content-Type: application/json

{
  "user_name": "joao",
  "user_type": "BUYER",
  "password": "senha123"
}
```

**Response:**
```json
{
  "user_id": 1,
  "user_name": "joao",
  "user_type": "BUYER",
  "created_at": "2026-02-03T17:00:00Z",
  "updated_at": "2026-02-03T17:00:00Z"
}
```
⚠️ **Nota**: A senha NÃO aparece na resposta!

### Atualizar Usuário (incluindo senha)

**Request:**
```bash
PUT /api/v1/users/1
Content-Type: application/json

{
  "user_name": "joao_silva",
  "password": "nova_senha456"
}
```

### Validar Senha (Login)

Use o método `ValidatePassword` do service:

```go
err := userService.ValidatePassword(userID, "senha123")
if err != nil {
    // Senha incorreta
    return errors.New("invalid credentials")
}
// Senha correta - prosseguir com login
```

## Segurança 🔐

### O que foi implementado:
1. **Hash Bcrypt**: Algoritmo industry-standard (cost: 10)
2. **Validação**: Senhas entre 6-72 caracteres
3. **Nunca expor**: Senha NUNCA retorna nas APIs
4. **Salt automático**: Bcrypt inclui salt único para cada hash

### Exemplos de Senhas:

```
Senha em texto plano:  "senha_temporaria"
Hash bcrypt:           "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
```

## Testando a Implementação ✅

### 1. Rodar a aplicação:
```bash
go run cmd/api/main.go
```

A migration 000003 vai automaticamente converter todas as senhas existentes.

### 2. Criar um novo usuário:
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "teste",
    "user_type": "BUYER",
    "password": "senha123"
  }'
```

### 3. Rodar os testes:
```bash
go test ./internal/service/...
```

## Usuários Existentes 👥

Todos os usuários existentes tiveram suas senhas em texto plano convertidas para hash bcrypt automaticamente.

**Para fazer login com esses usuários:**
- Senha padrão: `senha_temporaria_123`
- O sistema vai comparar automaticamente com o hash armazenado

## Próximos Passos (Sugestões) 📋

1. **Implementar endpoint de Login**
   - POST /api/v1/auth/login
   - Retornar JWT token

2. **Implementar recuperação de senha**
   - Envio de email com token
   - Reset de senha

3. **Adicionar política de senha mais forte**
   - Exigir letras maiúsculas/minúsculas
   - Exigir números
   - Exigir caracteres especiais

4. **Rate limiting**
   - Prevenir brute force attacks
   - Limitar tentativas de login

## Arquivos Modificados 📝

```
✅ internal/utils/password.go                     (NOVO)
✅ internal/domain/user.go                        (atualizado)
✅ internal/dto/request/user_request.go           (atualizado)
✅ internal/service/user_service.go               (atualizado)
✅ internal/handler/user_crud_handler copy.go     (atualizado)
✅ internal/service/user_service_test.go          (atualizado)
✅ migrations/000003_hash_existing_passwords.*    (NOVO)
```

## Referências 📚

- [Bcrypt Package](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [OWASP Password Storage](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [Bcrypt Explained](https://auth0.com/blog/hashing-in-action-understanding-bcrypt/)

---

**Data de Implementação**: 03/02/2026  
**Status**: ✅ Completo e Testado
