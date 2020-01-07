package stmp

import (
	"strconv"
)

const (
	MessageKindPing     = 0x0
	MessageKindPong     = 0x1
	MessageKindRequest  = 0x2
	MessageKindNotify   = 0x3
	MessageKindResponse = 0x4
	MessageKindClose    = 0x5
)

func shouldHeadOnly(kind byte) bool {
	return kind == MessageKindPing || kind == MessageKindPong
}

func hasMid(kind byte) bool {
	return kind == MessageKindRequest || kind == MessageKindResponse
}

func hasAction(kind byte) bool {
	return kind == MessageKindRequest || kind == MessageKindNotify
}

func hasStatus(kind byte) bool {
	return kind == MessageKindResponse || kind == MessageKindClose
}

func isValidKind(kind byte) bool {
	switch kind {
	case MessageKindPing, MessageKindPong, MessageKindRequest, MessageKindNotify, MessageKindResponse, MessageKindClose:
		return true
	default:
		return false
	}
}

const maskFin = 0x80
const maskHead = 0b1000
const offsetKind = 4

var mapTextKind = map[byte]byte{
	'I': MessageKindPing,
	'O': MessageKindPong,
	'Q': MessageKindRequest,
	'N': MessageKindNotify,
	'S': MessageKindResponse,
	'C': MessageKindClose,
	//'F': MessageKindFollowing,
}

var mapKindText = map[byte]byte{
	MessageKindPing:     'I',
	MessageKindPong:     'O',
	MessageKindRequest:  'Q',
	MessageKindNotify:   'N',
	MessageKindResponse: 'S',
	MessageKindClose:    'C',
}

type Status byte

const (
	// OK
	// for handshake, if status is not OK, will close directly
	StatusOk Status = 0x00
	// unknown
	StatusUnknown Status = 0x01
	// network error
	// sender write, sender read
	// receiver read or write error will omit the request
	// maybe timeout to close
	StatusNetworkError Status = 0x02
	// protocol error
	// for parse packet error
	StatusProtocolError Status = 0x03
	// sender error
	// for handshake, maybe: unsupported packet format, protocol version, content type, encoding
	// for requests, maybe: sender cannot marshal input, receiver unmarshal input error, handler emit
	StatusBadRequest Status = 0x20
	// authenticate error, or handler emit
	StatusUnauthorized Status = 0x21
	// action is not registered, or interceptors do not accept, or no registered handlers
	StatusNotFound Status = 0x22
	// sender cancelled, which means ctx.Done() returns before receive response
	StatusRequestTimeout Status = 0x23
	// packet too large, if is request, will send close with this status
	StatusRequestEntityTooLarge Status = 0x24
	// rate limit, not implemented
	StatusTooManyRequests Status = 0x25
	// server internal error
	// marshal output error, handler emit error
	StatusInternalServerError Status = 0x40
	// close connection when server close
	StatusServerShutdown Status = 0x41
)

var MapStatus = map[Status]string{
	StatusOk:                    "Ok",
	StatusNetworkError:          "NetworkError",
	StatusProtocolError:         "ProtocolError",
	StatusUnknown:               "Unknown",
	StatusBadRequest:            "BadRequest",
	StatusUnauthorized:          "Unauthorized",
	StatusNotFound:              "NotFound",
	StatusRequestTimeout:        "RequestTimeout",
	StatusRequestEntityTooLarge: "RequestEntityTooLarge",
	StatusTooManyRequests:       "TooManyRequests",
	StatusInternalServerError:   "InternalServerError",
	StatusServerShutdown:        "ServerShutdown",
}

// a status error
type StatusError interface {
	Error() string
	Spread() (Status, []byte)
	Code() Status
	Message() string
}

func (s Status) Error() string {
	if m, ok := MapStatus[s]; ok {
		return m
	}
	return "Unknown (0x" + strconv.FormatUint(uint64(s), 16) + ")"
}

func (s Status) Spread() (Status, []byte) {
	return s, []byte(s.Error())
}

func (s Status) Code() Status {
	return s
}

func (s Status) Message() string {
	return s.Error()
}

type detailedError struct {
	code    Status
	message string
}

func (e *detailedError) Code() Status {
	return e.code
}

func (e *detailedError) Message() string {
	return e.message
}

func (e *detailedError) Error() string {
	return "STMP " + e.code.Error() + ": " + e.message
}

func (e *detailedError) Spread() (Status, []byte) {
	return e.code, []byte(e.message)
}

func NewStatusError(code Status, data interface{}) StatusError {
	se := new(detailedError)
	se.code = code
	if m, ok := data.(string); ok {
		se.message = m
	} else if e, ok := data.(error); ok {
		se.message = e.Error()
	}
	if se.message == "" {
		se.message = se.code.Error()
	}
	return se
}

// Detect handler emitted error
//
// If the err is StatusError, will use it directly
// else will use the fallback as the status
func DetectError(err error, fallback Status) StatusError {
	if se, ok := err.(StatusError); ok {
		return se
	}
	return NewStatusError(fallback, err.Error())
}

const (
	AcceptContentType  = "Accept"
	AcceptEncoding     = "Accept-Encoding"
	AcceptPacketFormat = "Accept-Packet-Format"

	DetermineContentType  = "Content-Type"
	DetermineEncoding     = "Content-Encoding"
	DeterminePacketFormat = "Packet-Format"
	DetermineStmpVersion  = "Stmp-Version"
)
