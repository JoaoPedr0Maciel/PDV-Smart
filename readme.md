# Projeto PDV Smart

O PDV Smart é uma aplicação desenvolvida com o framework Gin para Go. Este README fornece instruções sobre como construir e executar a aplicação utilizando Docker.

## Pré-requisitos
- Docker instalado em sua máquina.
- Docker Compose instalado em sua máquina.
- Git instalado em sua máquina.

## Estrutura do Projeto
- `Dockerfile` e `docker-compose.yaml`: Arquivo de configuração para construir a imagem Docker da aplicação.
- `go.mod` e `go.sum`: Arquivos de dependência do Go.
- `main.go`: Código principal da aplicação.
- Diretórios `config`, `handlers`, `routes`, `schemas`, `database`: Contêm arquivos relacionados à configuração, manipulação de dados e definição de rotas da aplicação.

## Instalação
1. Clone esse repositório e entre no diretório do projeto:

    ```bash
    git clone https://github.com/JoaoPedr0Maciel/PDV-Smart
    cd PDV-Smart
    ```

2. Suba o container do Postgres e do PgAdmin utilizando Docker Compose:

    ```bash
    sudo docker compose up -d
    ```

    Isso cria a rede necessária e sobe os containers do Postgres e PgAdmin.

3. Para acessar o banco de dados, basta acessar:

    ```text
    http://localhost:8081
    ```

    - Email: `admin@admin.com`
    - Senha: `admin`

## Criar e conectar o banco de dados
    Crie uma nova "connection" (Nova conexão):

    - Nome da conexão: `pdv_docker`
    - HOST/ADDRESS: `[my_postgres_container]`
    - Maintenance Database: `pdv_smart`
    - Port: `5432`
    - Username: `postgres`
    - Password: `postgres`

4. Construa a imagem Docker da aplicação:

    ```bash
    sudo docker build -t pdv_smart .
    ```

5. Suba o container da aplicação:

    ```bash
    sudo docker run -d --name pdv_smart_container --network pdv-smart_app-network -p 8080:8080 pdv_smart
    ```
    
## Como Usar o Makefile

### 1. Executar Todos os Passos

Para parar e remover todos os containers, construir a imagem e iniciar o container, basta rodar:

```bash
make all
```

Este comando executa três ações em sequência:
- **clean**: Para e remove todos os containers.
- **build**: Executa `docker compose up -d` e cria a imagem Docker.
- **run**: Inicia o container com a rede e portas configuradas.

### 2. Parar e Remover Containers

Se você só precisa parar e remover os containers em execução, use:

```bash
make clean
```

Esse comando executa `docker stop` seguido de `docker rm` para remover os containers.

### 3. Subir o Docker Compose e Construir a Imagem

Para rodar o `docker compose` e construir a imagem Docker sem iniciar o container:

```bash
make build
```

Esse comando inicia os serviços com `docker compose` e depois constrói a imagem usando `docker build`.

### 4. Iniciar o Container

Para iniciar o container com as configurações de rede e portas:

```bash
make run
```

Este comando executa `docker run` com as opções configuradas (nome do container, rede e mapeamento de portas).

## Instalando o `make`

### Debian/Ubuntu:

1. Atualize os pacotes:
   ```bash
   sudo apt update
   ```
2. Instale o `make`:
   ```bash
   sudo apt install make
   ```

