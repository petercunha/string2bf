# string2bf
A simple [Go](https://golang.org/) program that converts a string to [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) code

## Usage
After changing settings to
```go
const STR string = "Hello world!"
const COMPRESSION bool = true
```
And running

`$ go run string2bf.go`

We get
> ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++.+++++++..+++.-------------------------------------------------------------------------------.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.--------.+++.------.--------.-------------------------------------------------------------------.

This translates to `Hello world!`


With compression turned off, String2bf will generate code that produces the same result but is much longer.
