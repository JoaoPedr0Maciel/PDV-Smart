# Projeto PDV Smart

O **PDV Smart** é uma aplicação desenvolvida com o framework [Gin](https://github.com/gin-gonic/gin) para Go. Este README fornece instruções sobre como construir e executar a aplicação utilizando Docker.

## Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/) instalado em sua máquina.
- [Docker Compose](https://docs.docker.com/compose/install/) instalado em sua máquina.
- [Git](https://git-scm.com/downloads) instalado em sua máquina.

## Estrutura do Projeto

- `Dockerfile` e `docker-compose.yaml`: Arquivo de configuração para construir a imagem Docker da aplicação.
- `go.mod` e `go.sum`: Arquivos de dependência do Go.
- `main.go`: Código principal da aplicação.
- Diretórios `config`, `handlers`, `routes`, `schemas`, `database`: Contêm arquivos relacionados à configuração, manipulação de dados e definição de rotas da aplicação.

## Instalação
Para iniciar clone esse repositório e logo em seguida entre no mesmo:
```bash
git clone https://github.com/JoaoPedr0Maciel/PDV-Smart

cd pdv_smart
```

## Construindo a Imagem Docker e subindo o container

Para construir a imagem Docker da aplicação, utilize o seguinte comando:

```bash
sudo docker build -t pdv_smart .
```
Para subir o container, execute o seguinte comando:

```bash
sudo docker run -d --name pdv_smart_container --network pdv_smart_app-network -p 8080:8080 pdv_smart
```

Para subir o container do postgres e do pgadmin, execute:

```bash
sudo docker compose up -d
```

Para acessar o banco de dados, basta acessar:

```bash
http://localhost:8081
```

- email: admin@admin.com
- senha: admin

## Criar e conectar o banco de dados

1 - Crie uma nova "connection" (New connection)

- Nome da conexão: pdv_docker
- HOST/ADDRESS: [my_postgres_container]
- Maintance Database: pdv_smart
- Port: [5432]
- Username: postgres
- Password: postgres
