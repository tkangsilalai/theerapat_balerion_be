# Theerapat.K — Balerion Backend Assignment

`bahttext` converts `decimal.Decimal` amounts into Thai Baht text.

## Install (use in another Go project)

```bash
go get github.com/tkangsilalai/theerapat_balerion_be@v0.1.0
```

## Usage in another go module

```go
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tkangsilalai/theerapat_balerion_be/bahttext"
)

func main() {
	d := decimal.RequireFromString("33333.75")
	out, err := bahttext.ToThaiBahtText(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์
}
```

## Usage in this repo

```bash
go run .

Thai Baht Text Converter
Enter a number (or type 'exit' to quit)
> 33333.75
สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์
> 1000000000000
หนึ่งล้านล้านบาทถ้วน
> exit
Goodbye.
```

## Run tests (in this repo)

```bash
go test ./...
```
