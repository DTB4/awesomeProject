# awesomeProject

Simple educational HTML/CSS site with Golang API back-end

## Requirements

- MySQL
- Go 1.16

## Installation

clone develop branch

## Deploy

-Config create `config.env` from `config.env.dist` file and fill it all carefully ;-)

-Migrate:

```
  $go run cmd/migration/migration.go -U
```

-Create directory for log files and add its path to config file like `/logs` 
