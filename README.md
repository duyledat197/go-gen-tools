[![Typing SVG](https://readme-typing-svg.demolab.com/?lines=Interview+Hao+Github+source&center=true)](https://git.io/typing-svg)

# Interview Hao

<p align="center">
<a href="https://github.com/harish-sethuraman/readme-components">
 <img  src="https://readme-components.vercel.app/api?component=logo&fill=black&logo=go&animation=spin&svgfill=15d8fe">  
 </a>
 <a href="https://github.com/harish-sethuraman/readme-components">
 <img  src="https://readme-components.vercel.app/api?component=logo&fill=black&logo=github&animation=spin">  
 </a>
</p>

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

<i>A curated list of Interview Hao READMEs</i>

<a href="https://github.com/duyledat197/interview-hao/stargazers"><img src="https://img.shields.io/github/stars/duyledat197/interview-hao" alt="Stars Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/network/members"><img src="https://img.shields.io/github/forks/duyledat197/interview-hao" alt="Forks Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/pulls"><img src="https://img.shields.io/github/issues-pr/duyledat197/interview-hao" alt="Pull Requests Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/issues"><img src="https://img.shields.io/github/issues/duyledat197/interview-hao" alt="Issues Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/duyledat197/interview-hao?color=2b9348"></a>
<a href="https://github.com/duyledat197/interview-hao/blob/master/LICENSE"><img src="https://img.shields.io/github/license/duyledat197/interview-hao?color=2b9348" alt="License Badge"/></a>

## Features

- Change method login to POST instead of GET.
- Add method register (POST /regisger).
- Split services layer to domain and transport layer.
- Write unit test for the domain layer.
- Write unit test for the repository/store layer.

## Project Structure

```sh
.
├── cmd
│   ├── gen-layer
│   │   ├── internal
│   │   │   ├── process.go
│   │   │   └── step.go
│   │   ├── main.go
│   │   ├── models
│   │   │   ├── cli_step.go
│   │   │   └── template.go
│   │   └── templates
│   │       ├── delivery
│   │       │   ├── create.tpl
│   │       │   ├── default.tpl
│   │       │   ├── delete.tpl
│   │       │   ├── list.tpl
│   │       │   ├── retrieve.tpl
│   │       │   └── update.tpl
│   │       ├── repository
│   │       │   └── default.tpl
│   │       └── service
│   │           ├── create.tpl
│   │           ├── default.tpl
│   │           ├── delete.tpl
│   │           ├── list.tpl
│   │           ├── retrieve.tpl
│   │           └── update.tpl
│   ├── protoc-gen-custom
│   │   ├── internal
│   │   │   └── generator.go
│   │   └── main.go
│   └── server
│       └── main.go
├── config
│   └── config.go
├── database
│   ├── migrations
│   │   ├── 0001_migrate.up.sql
│   │   ├── 0002_migrate.up.sql
│   │   └── 0003_migrate.up.sql
│   ├── queries
│   │   ├── hub.sql
│   │   └── user.sql
│   └── sqlc.yaml
├── developments
│   ├── docker-compose.yml
│   ├── gen-go.sh
│   ├── groonga-build.sh
│   ├── postgres.Dockerfile
│   ├── proto.Dockerfile
│   └── sqlc.yaml
├── docs
│   ├── html
│   │   └── index.html
│   └── swagger
│       ├── auth.swagger.json
│       ├── call.swagger.json
│       ├── customer.swagger.json
│       ├── data.swagger.json
│       ├── enum.swagger.json
│       ├── file.swagger.json
│       ├── project.swagger.json
│       ├── task.swagger.json
│       ├── team.swagger.json
│       └── user.swagger.json
├── go.mod
├── go.sum
├── internal
│   ├── deliveries
│   │   └── http
│   │       ├── auth.go
│   │       ├── http.go
│   │       └── user.go
│   ├── middleware
│   │   └── authentication.go
│   ├── models
│   │   ├── db.go
│   │   ├── hub.sql.go
│   │   ├── models.go
│   │   ├── querier.go
│   │   └── user.sql.go
│   ├── repositories
│   │   ├── team.go
│   │   └── user.go
│   └── services
│       └── user.go
├── Makefile
├── mocks
│   ├── DBTX.go
│   ├── Middleware.go
│   ├── Querier.go
│   ├── TeamRepository.go
│   └── UserRepository.go
├── pb
│   ├── auth_enums.pb.go
│   ├── auth_grpc.pb.go
│   ├── auth_methods.pb.go
│   ├── auth.pb.go
│   ├── auth.pb.gw.go
│   ├── auth.pb.validate.go
│   ├── enum_enums.pb.go
│   ├── enum_methods.pb.go
│   ├── enum.pb.go
│   ├── enum.pb.validate.go
│   ├── team_grpc.pb.go
│   ├── team.pb.go
│   ├── team.pb.gw.go
│   ├── team.pb.validate.go
│   ├── user_enums.pb.go
│   ├── user_grpc.pb.go
│   ├── user_methods.pb.go
│   ├── user.pb.go
│   ├── user.pb.gw.go
│   └── user.pb.validate.go
├── proto
│   ├── auth.proto
│   ├── enum.proto
│   ├── options
│   │   ├── annotations.pb.go
│   │   ├── annotations.proto
│   │   └── doc.go
│   └── user.proto
├── README.md
├── transform
│   ├── options.go
│   ├── team_transformer.go
│   ├── transform.go
│   └── user_transformer.go
└── utils
    ├── crypto
    │   ├── sha256.go
    │   └── sha256_test.go
    ├── helper
    │   ├── validation.go
    │   └── validation_test.go
    ├── metadata
    │   └── metadata.go
    ├── moduleutils
    │   └── module.go
    └── token.go
```

## Installation

Make sure you have Go installed ([download](https://golang.org/dl/)). Version `1.19` or higher is required.

Install make for start the server.

For Linux:

```sh
$sudo apt install make
```

For Macos:

```sh
$brew install make
```

## Start server

First of all, you must start postgres:

```sh
$make start-postgres
```

After that should migrate:

```sh
$make migrate
```

Start server with cmd/terminal:

```sh
$make run
```

Start server with docker:

```sh
$make docker-start
```

Your app should now be running on [localhost:5050](http://localhost:5050/).

## Tools:

Generate sql:

```sh
$make gen-sql
```

Generate proto:

```sh
$make gen-proto
```

Generate layer by DDD (delivery, service, repository):

```sh
$make gen-layer
```

## Unit test:

```sh
$make test
```

## License

MIT

**Free Software, Hell Yeah!**
