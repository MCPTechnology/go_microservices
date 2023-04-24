# Microservices with Goland

- [Microservices with Goland](#microservices-with-goland)
  - [Description](#description)
    - [What it does?](#what-it-does)
    - [What tech it uses this far?](#what-tech-it-uses-this-far)
    - [Tech Debt this far?](#tech-debt-this-far)
    - [Architeture and Design Patterns](#architeture-and-design-patterns)
  - [Materials this code follows](#materials-this-code-follows)
  - [Materials to take a look at](#materials-to-take-a-look-at)
  - [How to use the Project](#how-to-use-the-project)

*[Table of contents generated with markdown-toc](http://ecotrust-canada.github.io/markdown-toc/)*

## Description

This project is intended for experimenting with Golang for the creation of microservices

### What it does?

This is intended to be used as a guide and template for creating microservices with Golang and other stacks such as Rest APIs, Kafka, Docker, MongoDB, PostgreSQL, Kubernetes

### What tech it uses this far?

- Golang
  - net/http
  - [ServeMux](https://pkg.go.dev/net/http#ServeMux)

### Tech Debt this far?

- Products PUT API needs to be able to take only the fields that should be updated on a record
- MongoDB
- [Kafka](https://www.youtube.com/watch?v=-yVxChp7HoQ&ab_channel=AnthonyGG)
- [Rate Limiting Server](https://github.com/bitfield/tpg-tests/blob/main/req/1/req_test.go#L22)

### Architeture and Design Patterns

- Domain Driven Design
- [Repository Pattern](https://threedots.tech/post/repository-pattern-in-go/)
- [Factory Pattern](https://refactoring.guru/design-patterns/factory-method)
- [Singleton](https://refactoring.guru/design-patterns/singleton)

## Materials this code follows

1. [Introduction to Microservices with Go](https://www.youtube.com/watch?v=VzBGi_n65iU&list=RDCMUC2V1SxXFUa5YxVJvTsrCgyg&start_radio=1&rv=VzBGi_n65iU&t=52&ab_channel=NicJackson)
2. [DDD with Go](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)

## Materials to take a look at

1. [Practical Go](https://dave.cheney.net/practical-go)
2. [Go Wiki](https://github.com/golang/go/wiki)

## How to use the Project
