# Bank Server Backend

This project is a backend application for a bank. It includes a database schema, implemented and deployed with PostgreSQL in AWS RDS. The application also includes a set of RESTful HTTP APIs using Gin. The APIs are mocked for robust unit tests, handling errors, authenticating users, and securing the APIs with JWT and PASETO access tokens.

The project also includes a CI/CD pipeline using Github actions and AWS services (ECR, EKS, EC2, Secrets Manager, etc.). Which is shown in the action workflow.

**[updated] The aws EKS service is expensive, so I turn off the EKS cluster and EC2 instance for saving money!**

## Database Schema

The database schema is designed to store information about accounts, transfers, entries and users.

More details in the figure:

![Bank Schema](./bank_schema.png)

## APIs

The application exposes a set of RESTful APIs. Details in [API source code](https://github.com/luyi404/simplebank/tree/main/api)

## Authentication

The APIs are secured with JWT and PASETO access tokens.

## Workflow

The following diagram shows the continuous integration and continuous delivery pipeline:

[github workflow](https://github.com/luyi404/simplebank/tree/main/.github/workflows)

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [DB Docs](https://dbdocs.io/docs)

    ```bash
    npm install -g dbdocs
    dbdocs login
    ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

    ```bash
    npm install -g @dbml/cli
    dbml2sql --version
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

### Setup infrastructure

- Create the bank-network

    ``` bash
    make network
    ```

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```

### Documentation

- Generate DB documentation:

    ```bash
    make db_docs
    ```

- Access the DB documentation at [this address](https://dbdocs.io/techschool.guru/simple_bank). Password: `secret`

### How to generate code

- Generate schema SQL file with DBML:

    ```bash
    make db_schema
    ```

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

- Create a new db migration:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```
