# wikidump
`wikidump` is a small package for reading `bz2` compressed Wikipedia dump XML
files (e.g. `enwiki-latest-pages-articles.xml.bz2` from
https://dumps.wikimedia.org/enwiki/latest/).

```go
import (
    "github.com/msrocka/wikidump"
    "fmt"
)

func main() {
    reader, err := wikidump.NewReader("path/to/dump.xml.bz2")
    defer reader.Close()
    check(err)
    for {
        page, err := reader.NextPage()
        check(err)
        fmt.Println(page.Title)
    } 
}
```
