package mumbleclient

import (
	"bufio"
	"net"
)

type MumbleClient struct {
	conn       net.Conn      // TLS conn
	connReader *bufio.Reader // TLS conn reader

	Username string
	Password string
}
