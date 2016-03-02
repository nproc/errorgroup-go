# Error Group

`ErrorGroup` is an error type that can hold other errors together.

## Installation

```
go get -u -v github.com/nproc/errorgroup
```

## Why?

Sometimes I have a loop running functions in parallel and because of the
existing API of most the Go libs I need to return only one error.

I don't want to throw away such important information, so I group them together.

## License

MIT
