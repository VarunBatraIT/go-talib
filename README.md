# go-talib

[![GoDoc](http://godoc.org/github.com/thrasher-corp/go-talib?status.svg)](http://godoc.org/github.com/thrasher-corp/go-talib) 

A pure [Go](http://golang.org/) port of [TA-Lib](http://ta-lib.org)

## Original:

This project is forked from [go-talib](github.com/markcheno/go-talib)

## Install

Install the package with:

```bash
go get github.com/thrasher-corp/go-talib
```

Import it with:

```go
import "github.com/thrasher-corp/go-talib"
```

and use `talib` as the package name inside the code.

## Testing

### Prerequisite


1) Build and install [talib](https://downloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz) 

```bash
wget https://downloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz
tar -xvf ta-lib-0.4.0-src.tar.gz
cd ta-lib/
./configure --prefix=/usr
make
sudo make install
```

2) install ta-lib python package 

```bash
sudo pip install ta-lib
```

3) run tests

All package:
```bash
go test ./indicators/... 
```
Indivual Test:
```bash
go test ./indicators/. -run TestRsi -v
```

## Example

```go
package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/thrasher-corp/go-talib/inidcators"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := inidcators.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
```

## License

MIT License  - see LICENSE for more details

# Contributors

- [xtda](https://github.com/xtda)
- [Markcheno](https://github.com/markcheno) 
- [Alessandro Sanino AKA saniales](https://github.com/saniales)
