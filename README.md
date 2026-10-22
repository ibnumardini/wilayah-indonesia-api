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

- Go 1.23.1
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

- Go 1.23.1 or higher
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

## Documentation

API documentation is available at `/swagger/index.html` endpoint.
