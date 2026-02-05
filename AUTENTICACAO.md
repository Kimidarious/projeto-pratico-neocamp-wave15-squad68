# 🔐 Guia de Autenticação - SocialMeli API

Este documento explica como funciona a autenticação no sistema e como testar.

## 📚 O que foi implementado?

### 1. **JWT (JSON Web Token)** 🎫
É como um "cartão de identificação digital". Funciona assim:
- Você faz login com username e password
- O servidor valida e retorna um TOKEN (uma string grande)
- Você usa esse TOKEN nas próximas requisições para provar quem você é
- O TOKEN expira em 24 horas

### 2. **Estrutura Criada**

#### **DTOs (Data Transfer Objects)**
- `LoginRequest` - Envia username e password
- `LoginResponse` - Recebe token e informações do usuário

#### **Service (Camada de Negócio)**
- `AuthService` - Valida credenciais e gera tokens JWT
  - `Login()` - Autentica usuário
  - `ValidateToken()` - Verifica se token é válido
  - `generateToken()` - Cria o token JWT

#### **Handler (Camada de Apresentação)**
- `AuthHandler` - Recebe requisições HTTP
  - `Login()` - Endpoint POST /api/v1/auth/login

#### **Middleware**
- `AuthMiddleware` - Intercepta requisições e valida token
  - Verifica header "Authorization: Bearer <token>"
  - Adiciona informações do usuário no contexto da requisição

## 🚀 Como Testar

### Passo 1: Configurar variáveis de ambiente

Copie o arquivo `.env.example` para `.env`:

```bash
cp .env.example .env
```

Edite o `.env` e configure uma chave secreta forte para JWT:

```env
JWT_SECRET=minha-chave-super-secreta-123
```

> **IMPORTANTE**: Em produção, use uma chave forte! Gere com: `openssl rand -base64 32`

### Passo 2: Iniciar o servidor

```bash
go run cmd/api/main.go
```

### Passo 3: Criar um usuário de teste

Primeiro, crie um usuário:

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "joao123",
    "user_type": "buyer",
    "password": "senha123"
  }'
```

**Resposta esperada:**
```json
{
  "user_id": 1,
  "user_name": "joao123",
  "user_type": "buyer",
  "created_at": "2026-02-05T16:30:00Z",
  "updated_at": "2026-02-05T16:30:00Z"
}
```

### Passo 4: Fazer Login

Agora faça login com o usuário criado:

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "user_name": "joao123",
    "password": "senha123"
  }'
```

**Resposta esperada:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJqb2FvMTIzIiwidXNlcl90eXBlIjoiYnV5ZXIiLCJleHAiOjE3MDczMzg0MDB9.abc123...",
  "user_id": 1,
  "user_name": "joao123",
  "user_type": "buyer"
}
```

**Copie o valor do campo `token`!** Você vai precisar dele nas próximas requisições.

### Passo 5: Usar o token em requisições protegidas

Para usar o middleware de autenticação, adicione o token no header `Authorization`:

```bash
curl -X GET http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer SEU_TOKEN_AQUI"
```

Substitua `SEU_TOKEN_AQUI` pelo token que você recebeu no login.

## 🔒 Como Proteger Rotas

Para proteger uma rota, adicione o middleware no `main.go`:

```go
// Exemplo: proteger rotas de usuários
users := v1.Group("/users")
users.Use(middleware.AuthMiddleware(authService)) // Adiciona autenticação
{
    users.GET("/:id", userCRUDHandler.GetUser)
    // ... outras rotas protegidas
}
```

## 📋 Testes com diferentes cenários

### ✅ Login com sucesso
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"user_name": "joao123", "password": "senha123"}'
```

### ❌ Login com senha errada
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"user_name": "joao123", "password": "senhaerrada"}'
```

**Resposta esperada:**
```json
{
  "error": "invalid username or password"
}
```

### ❌ Login com usuário inexistente
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"user_name": "naoexiste", "password": "senha123"}'
```

**Resposta esperada:**
```json
{
  "error": "invalid username or password"
}
```

### ❌ Acessar rota protegida sem token
```bash
curl -X GET http://localhost:8080/api/v1/users/1
```

**Resposta esperada (se a rota estiver protegida):**
```json
{
  "error": "authorization header required"
}
```

### ❌ Acessar rota com token inválido
```bash
curl -X GET http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer token_invalido"
```

**Resposta esperada:**
```json
{
  "error": "invalid or expired token"
}
```

## 🧪 Testar com Postman

1. **Criar Request de Login**
   - Method: POST
   - URL: `http://localhost:8080/api/v1/auth/login`
   - Body (JSON):
     ```json
     {
       "user_name": "joao123",
       "password": "senha123"
     }
     ```
   - Clique em "Send"
   - Copie o `token` da resposta

2. **Usar Token em outras requisições**
   - Vá na aba "Authorization"
   - Selecione "Bearer Token"
   - Cole o token no campo "Token"
   - Faça suas requisições normalmente

## 📖 Documentação Swagger

Acesse a documentação interativa em:
```
http://localhost:8080/swagger/index.html
```

Lá você pode:
- Ver todos os endpoints
- Testar diretamente no navegador
- Ver exemplos de request/response

## 🔑 Informações Importantes

### O que está no Token JWT?
O token contém:
- `user_id` - ID do usuário
- `user_name` - Nome do usuário
- `user_type` - Tipo do usuário (buyer, seller, both)
- `exp` - Data de expiração (24 horas)
- `iat` - Data de emissão

### Como o Middleware funciona?
1. Intercepta a requisição HTTP
2. Pega o header `Authorization: Bearer <token>`
3. Valida o token usando a chave secreta (`JWT_SECRET`)
4. Se válido, adiciona informações no contexto:
   - `c.Get("user_id")` - ID do usuário autenticado
   - `c.Get("user_name")` - Nome do usuário
   - `c.Get("user_type")` - Tipo do usuário
5. Se inválido, retorna erro 401 (Unauthorized)

### Usando informações do usuário autenticado no Handler

Dentro de um handler protegido, você pode acessar os dados do usuário:

```go
func (h *MyHandler) MyProtectedEndpoint(c *gin.Context) {
    userID, _ := c.Get("user_id")
    userName, _ := c.Get("user_name")
    userType, _ := c.Get("user_type")
    
    // Use as informações...
    fmt.Printf("Usuário %s (ID: %d) acessou este endpoint\n", userName, userID)
}
```

## 🐛 Problemas Comuns

### Erro: "authorization header required"
- Você esqueceu de enviar o header `Authorization`
- Solução: Adicione o header `Authorization: Bearer <token>`

### Erro: "invalid authorization header format"
- O formato do header está errado
- Solução: Use exatamente `Bearer <token>` (com espaço)

### Erro: "invalid or expired token"
- O token expirou (24 horas)
- O token foi modificado
- A chave `JWT_SECRET` mudou
- Solução: Faça login novamente para obter um novo token

### Erro: "invalid username or password"
- Username ou senha incorretos
- Solução: Verifique as credenciais

## 📝 Próximos Passos

Agora você pode:
1. ✅ Criar usuários
2. ✅ Fazer login e receber token
3. ✅ Usar token em requisições
4. ⏳ Proteger rotas específicas com o middleware
5. ⏳ Implementar logout (adicionar blacklist de tokens se necessário)
6. ⏳ Implementar refresh token (renovar token sem fazer login novamente)

Boa sorte nos testes! 🚀
