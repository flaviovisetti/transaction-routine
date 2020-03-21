# Transaction Routine API

API responsável por processar transações financeiras.

## Dependências

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/get-started)

## Setup

**Gerar imagem docker**

Comando criar uma imagem base para utilizar no desenvolvimento.

```bash
docker-compose build
```

**Criar bando de dados**

Comando cria o banco de dados utilizado no ambiente de desenvolvimento e teste

```bash
docker-compose run --rm api ./scripts/create_database.sh
```

## Ambiente de desenvolvimento

Para acessar o terminal e executar comandos na sua aplicação.

```bash
docker-compose run --rm api bash
```

## Execução

### Rodar Aplicação

**Executar**

Comando executa a aplicação.

```bash
docker-compose up -d
```

### EXTRA

**Seeds**

Comando executa o arquivo de **seeds**. Para isso acontecer a aplicação precisar estar sendo executada.

```bash
docker-compose run --rm api ./scripts/seeds.sh
```

## Logs

Acessa os logs da aplicação e banco de dados

```bash
docker-compose logs -f
```