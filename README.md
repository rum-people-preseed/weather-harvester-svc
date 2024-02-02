## Weather Data Harvester

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is a Go microservice that harvests weather data from the [OpenWeatherMap API](https://openweathermap.org/api)
and stores it in a [PostgreSQL](https://www.postgresql.org/) database.

### Packages

We use:

1. [echo](https://github.com/labstack/echo) for routing
2. [zap](https://github.com/uber-go/zap) for logging
3. [dig](https://github.com/uber-go/dig) for dependency injection
4. [gorm](https://github.com/go-gorm/gorm) for ORM

### Linter

We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint our code. To run the linter, run the
following command:

```bash
golangci-lint run
```
