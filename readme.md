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

3. Construa a imagem Docker da aplicação:

    ```bash
    sudo docker build -t pdv_smart .
    ```

4. Suba o container da aplicação:

    ```bash
    sudo docker run -d --name pdv-smart_container --network pdv_smart_app-network -p 8080:8080 pdv_smart
    ```

5. Para acessar o banco de dados, basta acessar:

    ```text
    http://localhost:8081
    ```

    - Email: `admin@admin.com`
    - Senha: `admin`

## Criar e conectar o banco de dados
1. Crie uma nova "connection" (Nova conexão):

    - Nome da conexão: `pdv_docker`
    - HOST/ADDRESS: `[my_postgres_container]`
    - Maintenance Database: `pdv_smart`
    - Port: `5432`
    - Username: `postgres`
    - Password: `postgres`
