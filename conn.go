package ises



type ConnectionTypes int

const(
	STNone ConnectionTypes = iota
	STTcp
	STTls
	STHttp
	STWebsocket
)

type Connection struct {
	ConnType ConnectionTypes

}