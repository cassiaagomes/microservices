# Projeto – Microservices com Go e gRPC

Este projeto implementa uma arquitetura de **microserviços** utilizando **Go**, **gRPC** e **Docker**, com separação clara de responsabilidades entre os serviços.

## Estrutura do Projeto
microservices/
├── order/
├── payment/
├── shipping/
├── docker-compose.yml
├── banco.sql
└── README.md


### Serviços
- **order**: responsável pela criação de pedidos e orquestração do fluxo
- **payment**: responsável pelo processamento de pagamentos
- **shipping**: responsável pelo cálculo de prazo de entrega

Os serviços se comunicam entre si via **gRPC**, utilizando contratos definidos em Protobuf.

---

## Tecnologias Utilizadas

- Go
- gRPC
- Protocol Buffers
- Docker e Docker Compose
- PostgreSQL
- pgAdmin
- GORM
- OpenTelemetry (tracing)

---

## Subindo o Banco de Dados

Na raiz do projeto, execute:

```bash
docker compose up -d

sso irá subir:
PostgreSQL na porta 5432
pgAdmin na porta 5050
Acesso ao pgAdmin
URL: http://localhost:5050
Email: admin@admin.com
Senha: admin
Configuração do servidor no pgAdmin:
Host: postgres
Port: 5432
User: postgres
Password: postgres
Database: microservices



