package stmp

import (
	"compress/gzip"
	"io"
	"net"
)

var mapEncodingCodec = map[string]EncodingCodec{}

// register a new compression algorithm
func RegisterEncodingCodec(codecs ...EncodingCodec) {
	for _, codec := range codecs {
		mapEncodingCodec[codec.Name()] = codec
	}
}

func GetEncodingCodec(name string) EncodingCodec {
	return mapEncodingCodec[name]
}

type EncodingWriter interface {
	io.WriteCloser
	Flush() error
}

type EncodingCodec interface {
	Name() string
	Reader(r io.Reader) (io.ReadCloser, error)
	Writer(w io.Writer, level int) (EncodingWriter, error)
}

type gzipCodec struct{}

func (g gzipCodec) Name() string {
	return "gzip"
}

func (g gzipCodec) Reader(r io.Reader) (io.ReadCloser, error) {
	return gzip.NewReader(r)
}

func (g gzipCodec) Writer(w io.Writer, level int) (EncodingWriter, error) {
	return gzip.NewWriterLevel(w, level)
}

func NewGzipCodec() EncodingCodec {
	return gzipCodec{}
}

type plainEncoding struct {
	net.Conn
}

func (plainEncoding) Close() error {
	return nil
}

func (plainEncoding) Flush() error {
	return nil
}

func init() {
	RegisterEncodingCodec(NewGzipCodec())
}
