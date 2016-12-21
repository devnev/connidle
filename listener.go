// Copyright 2016 Mark Nevill. All Rights Reserved.
// See LICENSE for licensing terms.

package connidle

import (
	"net"
	"time"
)

// WithTimeout adds an idle timeout to all connections returned from the listener.
// Whenever an operation is started, a deadline is set after which read and write operations will
// fail. Later operations override the deadline of earlier operations.
func WithTimeout(l net.Listener, idleTimeout time.Duration) net.Listener {
	if idleTimeout == 0 {
		return l
	}
	wrap := func(c net.Conn) net.Conn {
		c = &idleConn{
			Conn:        c,
			idleTimeout: idleTimeout,
		}
		return c
	}
	return &connWrapper{
		Listener: l,
		wrap:     wrap,
	}
}

type connWrapper struct {
	net.Listener
	wrap func(net.Conn) net.Conn
}

func (w *connWrapper) Accept() (net.Conn, error) {
	c, err := w.Listener.Accept()
	if err != nil || c == nil {
		return nil, err
	}
	return w.wrap(c), nil
}
