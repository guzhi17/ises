package ises



type ConnectionTypes int

const(
	STNone ConnectionTypes = iota
	STTcp
	STTls
	STUdp
	STHttp
	STWebsocket
)

type Connection struct {
	ConnType ConnectionTypes

}