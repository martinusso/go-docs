# go-docs

[CPF](https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas) / [CNPJ](https://en.wikipedia.org/wiki/CNPJ) validator and generator written in Go.

[![Circle CI](https://circleci.com/gh/martinusso/go-docs/tree/master.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/martinusso/go-docs/tree/master)
[![Build Status](https://travis-ci.org/martinusso/go-docs.svg?branch=master)](https://travis-ci.org/martinusso/go-docs)
[![Coverage Status](https://coveralls.io/repos/github/martinusso/go-docs/badge.svg?branch=master)](https://coveralls.io/github/martinusso/go-docs?branch=master)
[![GoDoc](https://godoc.org/github.com/martinusso/go-docs?status.svg)](https://godoc.org/github.com/martinusso/go-docs)

## Usage

### CPF

`import "github.com/martinusso/go-docs/cpf"`

Valid return a boolean

```
valid := cpf.Valid("08507460003")
```

AssertValid return a boolean and the error if any

```
valid, err := cpf.AssertValid("08507460003")
```

Generate return a random valid CPF

```
doc := cpf.Generate()
```

### CNPJ

`import "github.com/martinusso/go-docs/cnpj"`

Valid return a boolean

```
valid := cnpj.Valid("34700442000162")
```

AssertValid return a boolean and the error if any

```
valid, err := cnpj.AssertValid("34700442000162")
```

Generate return a random valid CNPJ

```
doc := cnpj.Generate()
```

## License

This software is open source, licensed under the The MIT License (MIT). See [LICENSE](https://github.com/martinusso/go-docs/blob/master/LICENSE) for details.
