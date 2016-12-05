# string2bf
A simple Golang program that converts a string to Brainfuck code

## Usage
After changing settings to
```go
const STR string = "Hello world!"
const COMPRESSION bool = true
```
And running

`$ go run string2bf.go`

We get
> ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++.+++++++...+++.-------------------------------------------------------------------------------.+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.--------.+++.------.--------.-------------------------------------------------------------------.

**This translates to "Hello world!"**
