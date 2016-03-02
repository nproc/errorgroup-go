[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/nproc/errorgroup-go)
![Codeship](https://img.shields.io/codeship/a2f991d0-c2b4-0133-6f72-4e6bd7c806c7.svg?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/nproc/errorgroup-go.svg?style=flat-square)](https://codecov.io/github/nproc/errorgroup-go)
[![Go Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat-square)](https://goreportcard.com/report/github.com/nproc/errorgroup-go)

# Error Group

`ErrorGroup` is an error type that can hold other errors together.

## Installation

```
go get -u github.com/nproc/errorgroup
```

## Why?

Sometimes I have a loop running functions in parallel and because of the
existing API of most the Go libs I need to return only one error.

I don't want to throw away such important information, so I group them together.

## License

MIT
