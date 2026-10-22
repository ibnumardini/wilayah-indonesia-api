# Wilayah Indonesia API

![Visitors](https://api.visitorbadge.io/api/visitors?path=https%3A%2F%2Fgithub.com%2Fibnumardini%2Fbadges&label=repository%20visits&countColor=%230c7ebe&style=flat&labelStyle=none)
![License](https://img.shields.io/github/license/ibnumardini/wilayah-indonesia-api)
![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=flat&logo=mysql&logoColor=white)

REST API for Indonesian regional data (Provinces, Regencies/Cities, Districts, and Sub-districts/Villages).

## Data Source

This API uses data fetched from BPS (Statistics Indonesia) using [@ibnumardini/wilayah-indonesia](https://www.npmjs.com/package/@ibnumardini/wilayah-indonesia) CLI package.

## Features

- Province data
- Regency/City data
- District data
- Sub-district/Village data
- Swagger Documentation

## Tech Stack

- Go 1.27.3
- Chi Router
- MySQL Database
- Swagger/OpenAPI

## API Endpoints

### Base URL

```sh
https://wilayah-indonesia-api.mardini.dev
```

### Endpoints

- `GET /` - Welcome message
- `GET /provinces` - List of provinces
- `GET /regencies` - List of regencies/cities
- `GET /districts` - List of districts
- `GET /subdistricts` - List of sub-districts/villages
- `GET /swagger/index.html` - API Documentation

## Development

### Prerequisites

- Go 1.27.3 or higher
- MySQL Database

### Installation

1. Clone repository

    ```bash
    git clone https://github.com/ibnumardini/wilayah-indonesia-api.git
    cd wilayah-indonesia-api
    ```

2. Install dependencies

   ```bash
   go mod tidy
   ```

## Local development

The project includes [Task](https://taskfile.dev/) commands and [Air](https://github.com/air-verse/air) configuration for hot reload. The `dev` task runs Air through `go run`, so Air does not need to be installed globally:

```bash
task dev
```

The API listens on `http://localhost:8080` by default. Configure the existing database connection with environment variables such as `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASS`, and `DB_USE_SSL`.

## Docker

Docker Compose runs only the API; it does not create or initialize a database. By default, the container connects to a database on the host through `host.docker.internal:3306`. Compose loads all container environment variables from `.env` using `env_file`.

Copy the template to `.env` and adjust the database values before starting:

```bash
cp .env.example .env

task docker:up
```

The included `.env` uses the same defaults and is ignored by Git, so replace its database credentials for your environment.

Source changes are mounted into the development container and automatically rebuilt by Air. To build the production image without running it:

```bash
task docker:build
```

## Documentation

API documentation is available at `/swagger/index.html` endpoint.
