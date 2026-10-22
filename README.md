# Wilayah Indonesia API

REST API for Indonesian administrative regions: provinces, regencies/cities,
districts, and villages. Built with Go and MySQL, with interactive
documentation via Swagger.

![License](https://img.shields.io/github/license/ibnumardini/wilayah-indonesia-api)
![Go](https://img.shields.io/badge/Go-1.27.3-00ADD8?style=flat&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=flat&logo=mysql&logoColor=white)

## Features

- REST API for Indonesian administrative region data
- Interactive API documentation with Swagger UI
- Hot reload for local development with Air
- Docker-based development workflow

## Data Source

Regional data is sourced from BPS (Statistics Indonesia) through the
[@ibnumardini/wilayah-indonesia](https://www.npmjs.com/package/@ibnumardini/wilayah-indonesia)
CLI package.

## Prerequisites

| Tool | Notes |
| --- | --- |
| Go 1.27.3+ | Required |
| MySQL | Required |
| [Task](https://taskfile.dev/) | Required for project commands |
| Docker and Docker Compose | Optional |

> **Note:** The database must already exist before starting the API. This
> project does not create, migrate, or seed a database.

## Getting Started

1. Clone the repository

   ```bash
   git clone https://github.com/ibnumardini/wilayah-indonesia-api.git
   cd wilayah-indonesia-api
   ```

2. Download dependencies

   ```bash
   go mod download
   ```

3. Configure the environment

   ```bash
   cp .env.example .env
   ```

   Update the database values in `.env` to match your environment. The
   `SWAGGER_HOST` variable is optional. Leave it empty to use the current API
   origin, or set it to a public host without a scheme:

   ```dotenv
   SWAGGER_HOST=api.example.com
   ```

   See [.env.example](./.env.example) for the complete list of variables.

4. Run the API

   With Air hot reload:

   ```bash
   task dev
   ```

   Without hot reload:

   ```bash
   task run
   ```

   The API is available at `http://localhost:8080`.

## API Documentation

Swagger UI is available at
[`http://localhost:8080/swagger/index.html`](http://localhost:8080/swagger/index.html)
when the API is running.

Regenerate the Swagger artifacts after changing endpoint annotations:

```bash
task swagger
```

This updates [docs/docs.go](./docs/docs.go), [docs/swagger.json](./docs/swagger.json),
and [docs/swagger.yaml](./docs/swagger.yaml).

## Project Structure

The API uses a module-based architecture. Each region module owns its
repository contract, service logic, HTTP handler, response mapping, and
router:

```text
modules/
├── province/
├── regency/
├── district/
└── subdistrict/
```

The dependency direction inside a module is:

```text
HTTP handler -> service -> repository interface
                                  ^
                                  |
                         MySQL repository
```

Services and repositories receive `context.Context` from the request. HTTP
response DTOs and query parsing stay in the handler layer, while database
implementations remain details of the module. Shared infrastructure and
helpers are kept under [`pkg/`](./pkg/).

## License

This project is licensed under the terms of the [LICENSE](./LICENSE) file.
