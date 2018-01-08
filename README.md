# wikidump
`wikidump` is a small package for reading `bz2` compressed Wikipedia dump XML
files (e.g. `enwiki-latest-pages-articles.xml.bz2` from
https://dumps.wikimedia.org/enwiki/latest/):

```go
package main

import (
	"fmt"

	"github.com/msrocka/wikidump"
)

func main() {
	dump := "path/to/dump.bz2"
	reader, err := wikidump.NewReader(dump)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	// print all page titles in the wikidump
	err = reader.Read(func(p *wikidump.Page) bool {
		fmt.Println(p.Title)
		return true // return true to continue
	})
	if err != nil {
		panic(err)
	}
}

```
