package ises

//get Connection-Info, addr-info
//get User-info
//get ClientType




type Session interface {
	GetConnection() *Connection
	GetAccount() *Account
	GetCache() *Cache
}

func main() {
	//net.Addr
}