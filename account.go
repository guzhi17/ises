package ises

type AuthLevel int
const(
	ALNone 			AuthLevel 	= 0x0
	ALLoginBasic 	AuthLevel 	= 0x11
	ALLoginNormal 	AuthLevel 	= 0x14
	ALLoginHigh   	AuthLevel 	= 0x18
)

type Account struct {
	Auth AuthLevel

	Uid int64
	Realm string
	Os int32

}