package models

import(
	"net"
)

type Client struct{
	id int
	connection net.Conn
}
