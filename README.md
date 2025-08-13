![BEON LOGO](https://beontech.wpengine.com/wp-content/themes/Archive/app/images/navbar-logo.png)

# Golang Coding Exercise

## Requirements

- [Docker and Docker Compose](https://docs.docker.com/compose/install/)
- Your preferred API Client

***

## Setting up the environment

First, clone the repository locally, then run:

```
docker-compose build
docker-compose up
```

You should now be able to go to `localhost:8080` and get the home view

Use Postman or any other API client to test the endpoints

## Running Tests

The project includes some test to validate handlers and functionality:

### Run All Tests
```bash
go test ./internal/app/ -v
```

***

## Helpful Resources

- [Golang Docs](https://golang.org/doc/)
- [GORM Docs](https://gorm.io/docs/)
- [Resty Docs](https://github.com/go-resty/resty)
- [Echo Docs](https://echo.labstack.com/)
- [Data Source - ARMS Data API](https://www.ers.usda.gov/developer/data-apis/arms-data-api/)

***

Copyright (C) BEON Tech Studio - All Rights Reserved
