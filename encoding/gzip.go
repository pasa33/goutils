package encoding_utils

import (
	"bytes"
	"compress/gzip"
	"io"
)

var Gzip thisGzip

type thisGzip struct{}

func (thisGzip) Encode(b []byte) []byte {
	var compressedBuffer bytes.Buffer
	writer := gzip.NewWriter(&compressedBuffer)

	_, err := writer.Write(b)
	if err != nil {
		return []byte{}
	}

	if err = writer.Flush(); err != nil {
		return []byte{}
	}
	if err = writer.Close(); err != nil {
		return []byte{}
	}

	return compressedBuffer.Bytes()
}

func (thisGzip) Decode(b []byte) []byte {

	r, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return []byte{}
	}
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return []byte{}
	}

	return decompressed
}
