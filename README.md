# go-faker-cn
Random fake chinese data generator written in go

[![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit/v6?status.svg)](https://github.com/aszhc/go-faker-cn)[![license](https://camo.githubusercontent.com/a144ba63d68316751be370b83652223e0edcf4f6e401288d5c28183c5d868b6f/687474703a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d677265656e2e7376673f7374796c653d666c6174)](https://raw.githubusercontent.com/aszhc/go-faker-cn/main/LICENSE)


## Installation
```
go get github.com/aszhc/go-faker-cn
```

## Simple Usage
```go
package main

import (
	"fmt"
	"github.com/aszhc/go-faker-cn/faker"
)

func main() {
	faker.Seed(0)       // use crypto/rand
	fmt.Println(faker.Name())
	fmt.Println(faker.IdCard())
	fmt.Println(faker.Email())
	fmt.Println(faker.Car())
	fmt.Println(faker.Job())
	fmt.Println(faker.BirthDay())
}
```

## Functions

```go
Age()
BirthDay()
Car()
Color()
Email()
IPv4Address()
IPv6Address()
MacAddress()
HTTPStatusCode()
UserAgent()
Job()
Name()
Number(min int, max int) 
MobilePhone()
Date()
```

