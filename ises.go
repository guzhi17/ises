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
	OK(proto.Message)
	Response(code int32, msg proto.Message)
	Error(code int32, msg string)
}

type Request interface {
	Session
	Response
	Request(proto.Message)error
}

func main() {
	//net.Addr
}