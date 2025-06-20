# Pair-Programming no Neovim via TCP

Este projeto oferece um servidor TCP em Golang que permite a **colaboração remota em tempo real** em uma única sessão do Neovim. Em vez de sincronizar arquivos, o servidor cria um pseudo-terminal onde o Neovim é executado. Clientes remotos se conectam a ele e têm sua entrada de teclado e saída de tela redirecionadas para a sessão Neovim, permitindo que múltiplos usuários controlem o mesmo editor.

## Funcionalidades

* **Sessão Neovim Centralizada**: Uma única instância do Neovim é executada no servidor.
* **Controle Compartilhado**: Clientes conectados ao servidor TCP podem controlar a mesma sessão do Neovim.
* **Redirecionamento de I/O**: A entrada do teclado do cliente é enviada para o Neovim, e a saída da tela do Neovim é enviada de volta ao cliente, simulando um terminal compartilhado.
* **Conexão Única**: O servidor aceita apenas **uma conexão de cliente ativa** por vez para simplificar o controle.

## Como Usar

### 1. Clone o Repositório
Abra seu terminal e rode o comando:

```bash
git clone https://github.com/kristyancarvalho/nvim-share.git
```

### 2. Compile o Servidor
Na raiz do projeto (`nvim-share`) execute o seguinte comando:

```bash
make build
```

Isso criará um executável chamado `nvim-share`.

### 3. Inicie o Servidor
No mesmo terminal, inicie o servidor:

```bash
make run
```

Você verá mensagens indicando que o Neovim foi iniciado em um pseudo-terminal e o servidor está aguardando conexões.

### 4. Conecte com o Cliente
Abra um **novo terminal** (ou use um terminal de um computador remoto) e conecte-se ao servidor usando `socat`. Substitua `localhost` pelo IP do seu servidor, se necessário:

```bash
socat TCP4:localhost:8080 STDIO,raw,echo=0
```

Após a conexão, a **interface completa do Neovim** aparecerá no seu terminal cliente.
