# 📚 Netscape Bookmarks

Golang library to work with bookmark export files in [Netscape bookmark
format](https://msdn.microsoft.com/en-us/library/aa753582.aspx).

## Install

`go get github.com/suhodolskiy/netscape-bookmarks`

## Usage

### Convert bookmark export file to golang structs

```go
package main

import (
  ...

  bookmarks "github.com/suhodolskiy/netscape-bookmarks"
)

func main() {
  input, _ := os.OpenFile("./example.html", os.O_RDONLY, 0644)
  data, _ := bookmarks.Parse(input)
  fmt.Println("Result", data)
}
```

### Command line tool

`go run cmd/conver/main.go --input ./example.html --output ./example.json`

1. [Input file: bookmarks/example.html](https://github.com/suhodolskiy/netscape-bookmarks/blob/main/example.html)
2. [Output file: bookmarks/result.json](https://github.com/suhodolskiy/netscape-bookmarks/blob/main/result.json)

## Todo

- [ ] Add the ability to generate a bookmark file
- [ ] Add the ability to convert HTML file to other formats (YAML, XML)
