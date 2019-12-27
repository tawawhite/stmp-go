// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 20:03:48
package stmp

import (
	"compress/flate"
	"compress/gzip"
	"github.com/google/brotli/go/cbrotli"
	"io"
)

var mapEncodingCodec = map[string]EncodingCodec{}

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

type flateCodec struct{}

func (f flateCodec) Name() string {
	return "deflate"
}

func (f flateCodec) Reader(r io.Reader) (io.ReadCloser, error) {
	return flate.NewReader(r), nil
}

func (f flateCodec) Writer(w io.Writer, level int) (EncodingWriter, error) {
	return flate.NewWriter(w, level)
}

func NewFlateCodec() EncodingCodec {
	return flateCodec{}
}

type brotliCodec struct{}

func (b brotliCodec) Name() string {
	return "br"
}

func (b brotliCodec) Reader(r io.Reader) (io.ReadCloser, error) {
	return cbrotli.NewReader(r), nil
}

func (b brotliCodec) Writer(w io.Writer, level int) (EncodingWriter, error) {
	return cbrotli.NewWriter(w, cbrotli.WriterOptions{Quality: level}), nil
}

func NewBrotliCodec() EncodingCodec {
	return brotliCodec{}
}
