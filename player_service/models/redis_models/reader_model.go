package redis_models

import (
	"net"
	"os"
)

type ReaderModel struct {
	File   *os.File
	SongId string
	Conn   net.Conn
	UserId string
}
