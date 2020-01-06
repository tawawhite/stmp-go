package stmp

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack"
)

// A media codec should Marshal/Unmarshal its Name specified content-type
type MediaCodec interface {
	Name() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var mapMediaCodec = map[string]MediaCodec{}

// register custom media codec
func RegisterMediaCodec(codecs ...MediaCodec) {
	for _, codec := range codecs {
		mapMediaCodec[codec.Name()] = codec
	}
}

// get media codec by content type
func GetMediaCodec(name string) MediaCodec {
	return mapMediaCodec[name]
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

// create json codec
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

// create msgpack codec
func NewMsgpackCodec() MediaCodec {
	return msgpackCodec{}
}

type protobufCodec struct{}

func (protobufCodec) Name() string {
	return "application/protobuf"
}

func (protobufCodec) Marshal(v interface{}) ([]byte, error) {
	if pb, ok := v.(proto.Message); ok {
		return proto.Marshal(pb)
	}
	return nil, errors.New("invalid protobuf message to encode")
}

func (protobufCodec) Unmarshal(data []byte, v interface{}) error {
	if pb, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, pb)
	}
	return errors.New("invalid protobuf message to decode")
}

// create protobuf codec
func NewProtobufCodec() MediaCodec {
	return protobufCodec{}
}

// payload map will cache the marshalled data for the same content-type
type PayloadMap struct {
	in       interface{}
	payloads map[string][]byte
}

// marshal and cache the result according to the content type
func (p *PayloadMap) Marshal(conn *Conn) ([]byte, error) {
	if p.in == nil {
		return nil, nil
	}
	payload, ok := p.payloads[conn.Media.Name()]
	if !ok {
		var err error
		payload, err = conn.Media.Marshal(p.in)
		if err != nil {
			return nil, err
		}
		p.payloads[conn.Media.Name()] = payload
	}
	return payload, nil
}

// create a new PayloadMap
func NewPayloadMap(in interface{}) *PayloadMap {
	return &PayloadMap{in: in, payloads: map[string][]byte{}}
}

func init() {
	// maybe should be optimized
	RegisterMediaCodec(NewProtobufCodec(), NewMsgpackCodec(), NewJsonCodec())
}
