// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-26 18:33:52
package stmp

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack"
)

type MediaCodec interface {
	Name() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var mediaCodecs = map[string]MediaCodec{}

func RegisterMediaCodec(codec MediaCodec) {
	mediaCodecs[codec.Name()] = codec
}

func GetMediaCodec(name string) MediaCodec {
	return mediaCodecs[name]
}

var (
	ErrCodecInvalidValue = errors.New("invalid value for codec")
)

func init() {
	RegisterMediaCodec(NewJsonCodec())
	RegisterMediaCodec(NewMsgpackCodec())
	RegisterMediaCodec(NewProtobufCodec())
}

type jsonCodec struct{}

func (j jsonCodec) Name() string {
	return "application/json"
}

func (j jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (j jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}

func NewJsonCodec() MediaCodec {
	return jsonCodec{}
}

type msgpackCodec struct{}

func (m msgpackCodec) Name() string {
	return "application/msgpack"
}

func (m msgpackCodec) Marshal(v interface{}) ([]byte, error) {
	if mv, ok := v.(msgp.Marshaler); ok {
		return mv.MarshalMsg(nil)
	}
	return msgpack.Marshal(v)
}

func (m msgpackCodec) Unmarshal(data []byte, v interface{}) error {
	if mv, ok := v.(msgp.Unmarshaler); ok {
		_, err := mv.UnmarshalMsg(data)
		return err
	}
	return msgpack.Unmarshal(data, v)
}

func NewMsgpackCodec() MediaCodec {
	return msgpackCodec{}
}

type protobufCodec struct{}

func (p protobufCodec) Name() string {
	return "application/protobuf"
}

func (p protobufCodec) Marshal(v interface{}) ([]byte, error) {
	if pv, ok := v.(proto.Marshaler); ok {
		return pv.Marshal()
	}
	return nil, ErrCodecInvalidValue
}

func (p protobufCodec) Unmarshal(data []byte, v interface{}) error {
	if pv, ok := v.(proto.Unmarshaler); ok {
		return pv.Unmarshal(data)
	}
	return ErrCodecInvalidValue
}

func NewProtobufCodec() MediaCodec {
	return protobufCodec{}
}