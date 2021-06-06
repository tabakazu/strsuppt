# strsuppt

`strsuppt` is a library to string conversions like a ActiveSupport::Inflector in Rails.

# Installation
```
$ go get github.com/tabakazu/strsuppt
```

# Example

```go
package main

import (
	"github.com/tabakazu/strsuppt"
)

func main() {
	s := strsuppt.Underscore("StringSupport")
	println(s) // string_support
}
```
