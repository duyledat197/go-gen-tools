<div align="center">

<img height="120" alt="Thanks for visiting my repository" width="100%" src="https://raw.githubusercontent.com/BrunnerLivio/brunnerlivio/master/images/marquee.svg" />

# Interview Hao

<p align="center">
<a href="https://github.com/harish-sethuraman/readme-components">
 <img  src="https://readme-components.vercel.app/api?component=logo&fill=black&logo=go&animation=spin&svgfill=15d8fe">  
 </a>
 <a href="https://github.com/harish-sethuraman/readme-components">
 <img  src="https://readme-components.vercel.app/api?component=logo&fill=black&logo=github&animation=spin">  
 </a>
</p>

[![CI](https://github.com/idodod/protoc-gen-fieldmask/actions/workflows/ci.yml/badge.svg)](https://github.com/idodod/protoc-gen-fieldmask/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/idodod/protoc-gen-fieldmask)](https://goreportcard.com/report/github.com/idodod/protoc-gen-fieldmask)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/idodod/protoc-gen-fieldmask)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/idodod/protoc-gen-fieldmask)
![GitHub](https://img.shields.io/github/license/idodod/protoc-gen-fieldmask)

<i>A curated list of Interview Hao READMEs</i>

<a href="https://github.com/duyledat197/interview-hao/stargazers"><img src="https://img.shields.io/github/stars/duyledat197/interview-hao" alt="Stars Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/network/members"><img src="https://img.shields.io/github/forks/duyledat197/interview-hao" alt="Forks Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/pulls"><img src="https://img.shields.io/github/issues-pr/duyledat197/interview-hao" alt="Pull Requests Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/issues"><img src="https://img.shields.io/github/issues/duyledat197/interview-hao" alt="Issues Badge"/></a>
<a href="https://github.com/duyledat197/interview-hao/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/duyledat197/interview-hao?color=2b9348"></a>
<a href="https://github.com/duyledat197/interview-hao/blob/master/LICENSE"><img src="https://img.shields.io/github/license/duyledat197/interview-hao?color=2b9348" alt="License Badge"/></a>

</div>

## Features:

- Change method login to POST instead of GET.
- Add method register (POST /regisger).
- Split services layer to domain and transport layer.
- Write unit test for the domain layer.
- Write unit test for the repository/store layer.

## Project Structure:

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

<div align="center">

## Installation:

Make sure you have Go installed ([download](https://golang.org/dl/)). Version `1.19` or higher is required.

Install make for start the server.

For Linux:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">sudo apt install make</a></i></pre>
</h2>

For Macos:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">brew install make</a></i></pre>
</h2>

## How to start server:

First of all, you must start postgres:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make start-postgres</a></i></pre>
</h2>

After that should migrate:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make migrate</a></i></pre>
</h2>

Start server with cmd/terminal:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make run</a></i></pre>
</h2>

Start server with docker:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make docker-start</a></i></pre>
</h2>

## Tools:

Generate sql:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-sql</a></i></pre>
</h2>

Generate proto:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-sql</a></i></pre>
</h2>
```sh
$make gen-proto
```

Generate layer by DDD (delivery, service, repository):

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-layer</a></i></pre>
</h2>

## Unit test:

Generate mock:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-mock</a></i></pre>
</h2>

Run all test:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make test</a></i></pre>
</h2>

</div>

## License:

MIT

**Free Software, Hell Yeah!**
