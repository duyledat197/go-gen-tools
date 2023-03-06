<a name="readme-top"></a>
<div align="center">

<img height="120" alt="Thanks for visiting my repository" width="100%" src="https://raw.githubusercontent.com/BrunnerLivio/brunnerlivio/master/images/marquee.svg" />

# Go Gen Tools

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

<i>A curated list of Go Gen Tools READMEs</i>

<a href="https://github.com/duyledat197/go-gen-tools/stargazers"><img src="https://img.shields.io/github/stars/duyledat197/go-gen-tools" alt="Stars Badge"/></a>
<a href="https://github.com/duyledat197/go-gen-tools/network/members"><img src="https://img.shields.io/github/forks/duyledat197/go-gen-tools" alt="Forks Badge"/></a>
<a href="https://github.com/duyledat197/go-gen-tools/pulls"><img src="https://img.shields.io/github/issues-pr/duyledat197/go-gen-tools" alt="Pull Requests Badge"/></a>
<a href="https://github.com/duyledat197/go-gen-tools/issues"><img src="https://img.shields.io/github/issues/duyledat197/go-gen-tools" alt="Issues Badge"/></a>
<a href="https://github.com/duyledat197/go-gen-tools/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/duyledat197/go-gen-tools?color=2b9348"></a>
<a href="https://github.com/duyledat197/go-gen-tools/blob/master/LICENSE"><img src="https://img.shields.io/github/license/duyledat197/go-gen-tools?color=2b9348" alt="License Badge"/></a>

</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

There are many great README templates available on GitHub; however, I didn't find one that really suited my needs so I created this enhanced one. I want to create a README template so amazing that it'll be the last one you ever need -- I think this is it.

Here's why:
* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like creating a README from scratch
* You should implement DRY principles to the rest of your life :smile:

Of course, no one template will serve all projects since your needs may be different. So I'll be adding more in the near future. You may also suggest changes by forking this repo and creating a pull request or opening an issue. Thanks to all the people have contributed to expanding this template!

Use the `BLANK_README.md` to get started.
<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.

* [![Next][Next.js]][Next-url]
* [![React][React.js]][React-url]
* [![Vue][Vue.js]][Vue-url]
* [![Angular][Angular.io]][Angular-url]
* [![Svelte][Svelte.dev]][Svelte-url]
* [![Laravel][Laravel.com]][Laravel-url]
* [![Bootstrap][Bootstrap.com]][Bootstrap-url]
* [![JQuery][JQuery.com]][JQuery-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



## Features:

- [x] Auto generate protobuf files.
- [x] Auto generate mock interface for DDD.
- [x] Auto generate all layer of DDD.
- [x] Auto generate sql query with struct mapping and entities.
- [x] Auto migrate for Postgres.
- [x] Support generate repository layer for postgres, mongo, inmem.
- [x] Support mono repo architecture.
- [x] Auto generate cli with [cobra-cli](https://github.com/spf13/cobra-cli).
- [x] Support graceful shutdown.
- [x] Start kubernetes with [Kind](https://kind.sigs.k8s.io/).
- [x] Manage kubernetes with [Helm](https://helm.sh/).
- [x] Support vscode settings.
- [x] Support github workflows.
- [ ] Support [Twilio](https://www.twilio.com/) client.
- [ ] Support [Sendgrid](https://sendgrid.com/) client.
- [ ] Support [AWS](https://aws.amazon.com/) client.
- [x] Support metrics with [Prometheus](https://prometheus.io/).
- [x] Support [Grafana](https://grafana.com/) for monitor.
- [x] Support [Elasticsearch](https://www.elastic.co/) client.
- [x] Support [Ethereum](https://ethereum.org/) client.
- [x] Support [Kafka](https://kafka.apache.org/), [Nats](https://nats.io/) for message queue.
- [x] Support [Redis](https://redis.io/) client.
- [x] Support rate limit.
- [x] Support configuration for grpc client, grpc server, http client, http server.
- [x] Support [Hystrix config](https://github.com/Netflix/Hystrix) for circuit breaker.
- [x] Support [Consul](https://www.consul.io/) client for load balancer.
- [x] Support open tracing with [Jeager tracing](https://www.jaegertracing.io/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Project Structure:

```sh
.
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.go
├── database
│   ├── migrations
│   │   ├── 0001_migrate.up.sql
│   │   ├── 0002_migrate.up.sql
│   │   └── 0003_migrate.up.sql
│   └── queries
│       ├── hub.sql
│       ├── team.sql
│       └── user.sql
├── developments
│   ├── bdd_test.Dockerfile
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
│       └── ...
├── features
│   └── bdd.go
├── go.mod
├── go.sum
├── intergration_test.go
├── internal
│   ├── deliveries
│   │   ├── grpc
│   │   │   └── ...
│   │   └── http
│   │       ├── http.go
│   │       └── ...
│   ├── middleware
│   │   └── authentication.go
│   ├── models
│   │   ├── db.go
│   │   ├── hub.sql.go
│   │   ├── models.go
│   │   ├── querier.go
│   │   └── ...
│   ├── mongo
│   ├── postgres
│   ├── repositories
│   │   └── ...
│   └── services
│       └── ...
├── Makefile
├── mocks
│   ├── dbtx.go
│   ├── querier.go
│   └── ...
├── pb
│   └── ...
├── pkg
│   ├── grpc_client
│   │   ├── grpc.go
│   │   └── option.go
│   ├── grpc_server
│   │   ├── grpc.go
│   │   ├── health_check.go
│   │   └── middleware.go
│   ├── http_server
│   │   ├── http.go
│   │   ├── middleware.go
│   │   └── middleware_test.go
│   ├── hystrix
│   │   └── config.go
│   ├── registry
│   │   ├── consul.go
│   │   └── consul_test.go
│   └── tracing
│       └── open_tracing.go
├── proto
│   ├── enum.proto
│   ├── hub.proto
│   ├── options
│   │   ├── annotations.pb.go
│   │   ├── annotations.proto
│   │   └── doc.go
│   ├── search.proto
│   ├── team.proto
│   └── user.proto
├── README.md
├── third_party
│   ├── sendgrid
│   │   └── email.go
│   └── twilio
│       ├── messaging.go
│       └── twilio.go
├── tools
│   ├── gen-layer
│   │   ├── internal
│   │   │   ├── generator.go
│   │   │   └── step.go
│   │   ├── main.go
│   │   ├── models
│   │   │   ├── cli_step.go
│   │   │   ├── feature.go
│   │   │   └── template.go
│   │   ├── protoc-gen-custom
│   │   │   └── internal
│   │   ├── templates
│   │   │   ├── cucumber
│   │   │   │   ├── create.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── delivery
│   │   │   │   ├── create.tpl
│   │   │   │   ├── default.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── godog
│   │   │   │   ├── create.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── postgres
│   │   │   │   ├── create.tpl
│   │   │   │   ├── default.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── repository
│   │   │   │   └── default.tpl
│   │   │   └── service
│   │   │       ├── create.tpl
│   │   │       ├── default.tpl
│   │   │       ├── delete.tpl
│   │   │       ├── list.tpl
│   │   │       ├── retrieve.tpl
│   │   │       └── update.tpl
│   │   └── utils
│   │       ├── parser
│   │       │   └── parser.go
│   │       └── steps.go
│   └── protoc-gen-custom
│       ├── internal
│       │   └── generator.go
│       └── main.go
├── transform
│   └── ...
└── utils
    ├── authenticate
    │   ├── authenticator.go
    │   ├── jwt.go
    │   ├── jwt_test.go
    │   ├── paseto.go
    │   ├── paseto_test.go
    │   ├── payload.go
    │   └── token.go
    ├── crypto
    │   ├── sha256.go
    │   └── sha256_test.go
    ├── helper
    │   ├── validation.go
    │   └── validation_test.go
    ├── ip.go
    ├── logger
    │   └── zap.go
    ├── metadata
    │   └── metadata.go
    ├── pathutils
    │   └── path.go
    ├── string.go
    ├── token.go
    └── transformhelpers
        └── helpers.go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


# Architechture:

![clean architecture](https://raw.githubusercontent.com/phungvandat/clean-architecture/dev/images/clean-arch.png)

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
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-proto</a></i></pre>
</h2>

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

## Pprof:
install graphviz:

```
$ go get -u github.com/google/pprof

$ apt-get install graphviz gv // for linux/debian 
$ brew install graphviz (mac)  // for mac
```

## License:

MIT

**Free Software, Hell Yeah!**
