# URL Shortener Frontend (React + Vite)

Este projeto é um frontend React criado com Vite para consumir a API de encurtador de URL definida no backend (main.go).

## Como rodar localmente

1. Certifique-se de que o backend (API Go) está rodando localmente, por padrão em `http://localhost:8080`.
2. Instale as dependências do frontend (caso ainda não tenha feito):

   ```bash
   npm install
   ```

3. Inicie o frontend em modo desenvolvimento:

   ```bash
   npm run dev
   ```

4. Acesse o frontend em `http://localhost:5173` (ou a porta indicada no terminal).

## Configuração de Proxy (opcional)

Para evitar problemas de CORS durante o desenvolvimento, adicione um proxy no arquivo `vite.config.js`:

```js
// vite.config.js
export default {
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
};
```

Assim, as requisições para `/api/shorten` serão redirecionadas para o backend Go.

## Funcionalidades
- Informe uma URL e clique em "Encurtar URL".
- O resultado será exibido com o link encurtado.
- Mensagens de erro são exibidas em caso de falha.

---

Para dúvidas sobre a API, consulte o arquivo `main.go` no backend.
