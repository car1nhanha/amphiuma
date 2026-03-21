# Amphiuma

Aplicação fullstack para listar e renderizar conteúdo Markdown vindo do GitHub.

Stack principal:

- Frontend: Vue 3, Vite, TypeScript
- Backend: Go, Gin

## Visão geral

O frontend permite:

- Buscar conteúdos de um usuário do GitHub
- Filtrar por extensão de arquivo
- Abrir e renderizar um arquivo Markdown

O backend expõe endpoints que:

- Buscam arquivos via GitHub Search API
- Buscam conteúdo de arquivo no repositório
- Extraem front matter quando disponível

## Estrutura do projeto

amphiuma/

- api/
  - cmd/api/main.go
  - internal/files/
  - internal/interfaces/
  - pkg/utils/
- app/
  - src/
  - public/

## Pré-requisitos

- Node.js 18+
- npm
- Go 1.24+

## Variáveis de ambiente

Backend, arquivo api/.env:

- GH*PERSONAL_TOKEN=Token ghp*...
- IS_RUNNING_LOCAL=true

Frontend, arquivo app/.env:

- VITE_API_BACKEND=http://localhost:8080
- VITE_PROJECT_NAME=Amphiuma

Importante:

- Não comitar tokens reais no repositório.

## Instalação

Na raiz do projeto:

1. Frontend
   - cd app
   - npm install

2. Backend
   - cd ../api
   - go mod tidy

## Como rodar em desenvolvimento

1. Subir backend local
   - cd api
   - IS_RUNNING_LOCAL=true go run ./cmd/api

2. Subir frontend
   - cd app
   - npm run dev

URLs locais:

- Frontend: http://localhost:5173
- Backend: http://localhost:8080

## Build do frontend

- cd app
- npm run build
- npm run preview

## Rotas do frontend

- /
  - Landing
- /:user
  - Lista de arquivos do usuário
- /:user/:repo/:path
  - Visualização do arquivo

## Endpoints da API

1. GET /:user

Lista arquivos do usuário no GitHub.

Query params:

- extension, opcional, padrão md

Exemplo:

- GET http://localhost:8080/car1nhanha?extension=md

2. GET /:user/:repo/\*path

Busca e retorna o conteúdo de um arquivo específico.

Exemplo:

- GET http://localhost:8080/car1nhanha/amphiuma/README.md

3. GET /

Health check simples.

## Fluxo resumido

1. Usuário acessa a landing
2. Informa o username do GitHub
3. Frontend consulta GET /:user
4. Usuário abre um arquivo
5. Frontend consulta GET /:user/:repo/\*path
6. Frontend renderiza o Markdown

## Observações úteis

- Se o backend subir sem IS_RUNNING_LOCAL=true, ele entra em modo Lambda.
- O frontend usa VITE_API_BACKEND para apontar para a API.

## Autor

Lucas Carinhanha

- GitHub: https://github.com/car1nhanha
