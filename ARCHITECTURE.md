# Architecture

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
implementations remain details of the module. Shared helpers and application
infrastructure are kept under [`internal/`](./internal/).

## Adding a Module

For a new API resource, create a directory under
[`modules/`](./modules/) with these responsibilities:

```text
modules/example/
├── entity.go              # Domain models
├── repository.go          # Repository interface
├── mysql_repository.go    # MySQL implementation
├── service.go             # Application logic
├── response.go            # HTTP response DTOs
├── handler.go             # HTTP handlers and mapping
├── router.go              # Module routes
└── module.go              # Dependency wiring
```

Keep the service independent from HTTP and SQL details. Add the module route
in [`api/index.go`](./api/index.go), add endpoint annotations, regenerate
Swagger with `task swagger`, and add service or handler tests before merging.
