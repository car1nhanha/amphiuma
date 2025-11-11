# 📝 GitHub Blog - Amphiuma

Um blog fullstack que renderiza posts em Markdown diretamente do GitHub!

![Vue.js](https://img.shields.io/badge/Vue.js-42b883?style=for-the-badge&logo=vue.js&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-646CFF?style=for-the-badge&logo=vite&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)

---

## 🎯 Sobre o Projeto

**GitHub Blog** é uma aplicação que permite visualizar e renderizar arquivos Markdown diretamente do GitHub. A aplicação consiste em um **frontend moderno** desenvolvido com Vue 3 + Vite e um **backend robusto** em Go com Gin.

Acesse repositórios GitHub através de uma URL simples e visualize seus posts em Markdown de forma elegante e responsiva!

---

## 🚀 Stack Tecnológico

### Frontend (`/app`)

- ⚡ **Vite** — Build tool ultra-rápida
- 🖼️ **Vue 3** — Framework progressivo reativo
- 📘 **TypeScript** — Type safety no JavaScript
- 🎨 **CSS3** — Styling responsivo
- 📅 **date-fns** — Formatação de datas
- ✍️ **Marked** — Parser de Markdown
- 🧭 **Vue Router** — Roteamento de páginas
- 🎨 **Iconify** — Biblioteca de ícones

### Backend (`/api`)

- 🐹 **Go 1.24.1** — Linguagem eficiente
- 🍸 **Gin** — Framework web de alto desempenho
- 🔐 **CORS** — Controle de requisições cross-origin

---

## 📂 Estrutura do Projeto

```
github-blog/
├── api/                           # Backend em Go
│   ├── cmd/
│   │   └── api/
│   │       └── main.go           # Entrada da aplicação
│   ├── internal/
│   │   ├── files/
│   │   │   ├── handler.go        # Handlers HTTP
│   │   │   └── get_file.go       # Lógica de busca de arquivos
│   │   └── interfaces/
│   │       └── github_response.go # Structs de resposta GitHub
│   ├── pkg/
│   │   └── utils/
│   │       └── base64.go         # Utilitários de encoding
│   └── go.mod
│
├── app/                           # Frontend em Vue 3 + TypeScript
│   ├── src/
│   │   ├── main.ts               # Entry point
│   │   ├── style.css             # Estilos globais
│   │   ├── components/
│   │   │   ├── atoms/            # Componentes pequenos
│   │   │   ├── molecules/        # Componentes médios
│   │   │   ├── organisms/        # Componentes grandes
│   │   │   ├── pages/            # Páginas de rota
│   │   │   └── templates/        # Templates reutilizáveis
│   │   ├── service/
│   │   │   └── api.ts            # Chamadas HTTP para o backend
│   │   └── assets/               # Imagens e recursos
│   ├── public/                    # Arquivos estáticos
│   ├── index.html
│   ├── package.json
│   ├── vite.config.ts
│   ├── tsconfig.json
│   └── README.md
│
└── README.md                      # Este arquivo
```

---

## 🛠️ Configuração e Instalação

### Pré-requisitos

- **Node.js** >= 18.x
- **npm** ou **yarn**
- **Go** >= 1.24.1

### 1️⃣ Clone o Repositório

```bash
git clone https://github.com/car1nhanha/amphiuma.git
cd amphiuma
```

### 2️⃣ Instale Dependências do Frontend

```bash
cd app
npm install
```

### 3️⃣ Instale Dependências do Backend

```bash
cd ../api
go mod download
```

---

## 🏃 Como Executar

### Frontend (Desenvolvimento)

```bash
cd app
npm run dev
```

A aplicação estará disponível em `http://localhost:5173`

### Frontend (Produção)

```bash
cd app
npm run build
npm run preview
```

### Backend

```bash
cd api
go run ./cmd/api
```

O servidor estará rodando em `http://localhost:8080`

---

## 📡 API Endpoints

### GET `/v1/:user/:path`

Busca um arquivo Markdown do repositório GitHub de um usuário.

**Parâmetros:**

- `user` — Usuário do GitHub
- `path` — Caminho do arquivo dentro do repositório

**Exemplo:**

```bash
curl http://localhost:8080/v1/car1nhanha/posts/meu-post.md
```

**Resposta:**

```json
{
  "content": "# Meu Post\n\nConteúdo em Markdown...",
  "encoding": "base64"
}
```

---

## 🎨 Componentes Vue

### Atoms

- `input-text.vue` — Input de texto reutilizável

### Molecules

- `Card-header.vue` — Cabeçalho de cards
- `Card-posts.vue` — Card para exibir posts

### Organisms

- `Header.vue` — Cabeçalho principal
- `Stylize-post.vue` — Renderizador de posts em Markdown

### Pages

- `Home.vue` — Página inicial com lista de posts
- `Post.vue` — Página de visualização de um post

### Templates

- `Default.vue` — Template padrão de layout

---

## 🔄 Fluxo da Aplicação

1. 👤 Usuário acessa a URL `/[usuario-github]`
2. 🏠 Frontend exibe a página inicial com lista de posts
3. 📄 Usuário clica em um post
4. 🔄 Frontend faz requisição ao backend: `/v1/:user/:path`
5. 🐙 Backend busca o arquivo no GitHub API
6. 📝 Arquivo é decodificado e retornado
7. ✨ Frontend renderiza o Markdown com estilo

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

## 👨‍💻 Autor

**Lucas Carinhanha**

- GitHub: [@car1nhanha](https://github.com/car1nhanha)
- Email: lucascarinhanha4@gmail.com

---

## 🆘 Suporte

Se tiver dúvidas ou encontrar problemas:

1. 🔍 Verifique se o GitHub está acessível
2. 🔐 Confira se o usuário/repositório existe
3. 📌 Veja os logs do backend para mais detalhes
4. 💬 Abra uma [issue](https://github.com/car1nhanha/amphiuma/issues)

---

**Feito com ódio por [Lucas Carinhanha](https://github.com/car1nhanha)**
