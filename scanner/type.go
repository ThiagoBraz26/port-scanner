package scanner

import ("net")

type DialResult struct {
	Conn net.Conn
	Err error
	Port int
}