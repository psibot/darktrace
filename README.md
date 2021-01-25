# Darktrace API Library

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/rfizzle/darktrace)

A golang API library for Darktrace appliances.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Install

This project uses [go modules](https://golang.org/ref/mod)

You can install it locally by running:

```sh
$ go get github.com/rfizzle/darktrace
```

## Usage

Import the library into your project, setup the client, and call the API with valid parameters:

```go
package main

import (
    "github.com/rfizzle/darktrace"
    "log"
)

func main() {
  client, err := darktrace.NewClient(
    "https://darktrace.example.com",
    "publicToken",
    "privateToken",
  )
    
  if err != nil {
    log.Fatal(err)
  }
  
  results, err := client.EventList(
  	darktrace.Param("pbid", "6000000053951"),
  	darktrace.Param("includetotalbytes", "true"), 
  )

  if err != nil {
  	log.Fatal(err)
  }

  if len(results) > 0 {
  	return
  }

  return
}
```

## Maintainers

[@rfizzle](https://github.com/rfizzle)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/rfizzle/darktrace/issues/new) or submit PRs.

## License

[MIT](LICENSE) © Coleton Pierson