package codec

import "io"

type Header struct {
	ServiceMethod string //"服务名.方法名"
	Seq           uint64 //请求序号，用来区分不同的请求
	Error         string //客户端置为空，服务端发生错误，将错误信息置于Error中
}

// 对消息体进行编解码
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JSONType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
