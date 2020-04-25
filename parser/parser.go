package parser

import (
	"bytes"
	"time"
)

// Parse and change information in given text
func Parse(txt []byte, info string) []byte {
	dt := time.Now()

	txt = bytes.ReplaceAll(txt, []byte("%DATE%"), []byte(dt.Format("01-02-2006")))
	txt = bytes.ReplaceAll(txt, []byte("%TITLE%"), []byte(info))

	return txt
}
