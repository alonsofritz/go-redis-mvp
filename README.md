# Go Redis MVP - Aplicação de Demonstração com Golang e Redis

![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white) ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)

Uma aplicação criada para estudo sobre a utilização do Redis com a linguagem Go.

## Visão Geral

Este projeto é um MVP (Minimum Viable Product) que implementa um contador distribuído usando Redis como armazenamento persistente. A aplicação expõe um endpoint HTTP que incrementa e retorna o valor atual do contador a cada requisição.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação eficiente e de alto desempenho
- **Redis**: Banco de dados em memória para armazenamento do contador
- **Docker**: Containerização da aplicação e do Redis
- **Docker Compose**: Orquestração dos serviços

## Como Executar

1. Certifique-se de ter o Docker e Docker Compose instalados
2. Clone este repositório
3. Execute o comando:

```bash
docker-compose up --build
```

4. Acesse o Endpoint
```bash
curl http://localhost:8080/increment
```

## Funcionalidades
- Incrementa um contador a cada requisição
- Armazena o valor persistentemente no Redis
- Configuração simplificada com Docker Compose
- Health check para o serviço Redis

## Estrutura do Projeto
```bash
go-redis-mvp/
├── Dockerfile          # Configuração do container da aplicação
├── docker-compose.yml  # Definição dos serviços
├── go.mod              # Dependências do Go
└── main.go             # Código fonte da aplicação
```

## Escopo Geral
- Integração entre Go e Redis
- Configuração de containers Docker para desenvolvimento
- Uso de redes Docker para comunicação entre serviços
- Implementação de health checks para serviços dependentes