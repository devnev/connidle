// Copyright 2016 Mark Nevill. All Rights Reserved.
// See LICENSE for licensing terms.

package connidle

import (
	"net"
	"time"
)

type idleConn struct {
	net.Conn
	idleTimeout time.Duration
}

func (c *idleConn) Read(b []byte) (n int, err error) {
	err = c.Conn.SetDeadline(time.Now().Add(c.idleTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Read(b)
}

func (c *idleConn) Write(b []byte) (n int, err error) {
	err = c.Conn.SetDeadline(time.Now().Add(c.idleTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Write(b)
}
