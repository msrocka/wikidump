package wikidump

import (
	"bufio"
	"compress/bzip2"
	"encoding/xml"
	"io"
	"os"
)

// Page contains the information of a Wikipedia page
type Page struct {

	// The page title
	Title string `xml:"title"`

	// The ID of the page
	ID int64 `xml:"id"`

	// The raw text in Wikimedia syntax
	Text string `xml:"revision>text"`
}

// Reader reads wiki-pages from a bzip2 compressed XML dump
type Reader struct {
	dump *os.File
	dec  *xml.Decoder
}

// NewReader creates a new reader for the pages of the bzip2 compressed XML dump
// in the file with the given path.
func NewReader(file string) (*Reader, error) {
	dump, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	buffer := bufio.NewReader(dump)
	reader := bzip2.NewReader(buffer)
	return &Reader{dump: dump, dec: xml.NewDecoder(reader)}, nil
}

// NextPage reads the next wiki-page from the dump. We expect that the dump-file
// is not a history dump and thus there is just one revision per page included.
// If there are no more pages in the dump it returns nil, io.EOF.
func (r *Reader) NextPage() (*Page, error) {
	for {
		t, err := r.dec.Token()
		if err != nil {
			return nil, err
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "page" {
				var p Page
				err = r.dec.DecodeElement(&p, &se)
				return &p, err
			}
		}
	}
}

// Close closes the underlying dump file.
func (r *Reader) Close() error {
	return r.dump.Close()
}

func (r *Reader) Read(fn func(*Page) bool) error {
	for {
		page, err := r.NextPage()
		if err == nil {
			if !fn(page) {
				break
			}
			continue
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
	return nil
}
