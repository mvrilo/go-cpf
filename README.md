# go-cpf

CPF validator and generator written in Go.

## Usage

#### Command-line example:
    $ go-cpf 832.715.604-77

    To generate, just run it without any arguments.

#### API example:
```go
package main

import "github.com/mvrilo/go-cpf"

func main() {
          println(cpf.GeneratePretty())
}
```

## LICENSE

See [LICENSE](https://github.com/mvrilo/go-cpf/blob/master/LICENSE)

## Reference

* [https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas]
* [https://github.com/mvrilo/go-cpf]
