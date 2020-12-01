package ises

//get Connection-Info, addr-info
//get User-info
//get ClientType

import (
	"google.golang.org/protobuf/proto"
)



type Session interface {
	GetConnection() *Connection
	GetAccount() *Account
	GetCache() *Cache
}

type Response interface {
	ResponseOK(proto.Message, ...int32)
	ResponseError(code int32, msg string)
}

type Request interface {
	Session
	Response
	Request(proto.Message)error
}

func main() {
	//net.Addr
}