package parser

import (
	"bytes"
	"time"
)

// Parse and change information in given text
func Parse(txt []byte) []byte {
	dt := time.Now()

	txt = bytes.ReplaceAll(txt, []byte("%DATE%"), []byte(dt.Format("01-02-2006")))

	return txt
}
