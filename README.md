# URL Shortener

Este é um projeto de aprendizagem de Go, com um backend em Go e um frontend em React.

O objetivo é praticar conceitos de backend com Go, MongoDB e integração com frontend moderno.

## Sobre o Frontend

O frontend React foi gerado automaticamente por IA (GitHub Copilot) a pedido do autor, exclusivamente para fins de estudo e integração com a API Go definida em `main.go`.

## Como rodar o frontend localmente

1. Certifique-se de que o backend (API Go) está rodando localmente, conforme definido no `main.go`.
2. No terminal, acesse a pasta `frontend`:

```bash
cd frontend
```

3. Instale as dependências do projeto React:

```bash
npm install
```

4. Inicie o servidor de desenvolvimento:

```bash
npm run dev
```

5. Acesse o frontend em [http://localhost:5173](http://localhost:5173) (ou a porta indicada no terminal).

> **Nota:** O frontend espera que a API Go esteja disponível em `http://localhost:8080` (ou a porta configurada no backend). Se necessário, ajuste o proxy no `vite.config.js`.

---

README e frontend React gerados por IA (GitHub Copilot) para fins de estudo.
