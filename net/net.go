//@auther foolbread
//@time 2015-08-18
//@file net.go
package net

import (
	"io"
	"net"
	"time"
)

func ReadByTimeout(con net.Conn, d []byte, timeout time.Duration) error {
	l := len(d)
	s := 0
	con.SetReadDeadline(time.Now().Add(timeout))
	for s < l {
		n, err := con.Read(d[s:])
		if err != nil {
			return err
		}
		s = s + n
	}

	return nil
}

func ReadByCnt(r io.Reader, d []byte) error {
	l := len(d)
	s := 0
	for s < l {
		n, err := r.Read(d[s:])
		if err != nil {
			return err
		}
		s = s + n
	}

	return nil
}

func WriteByCnt(r io.Writer, d []byte) error {
	l := len(d)
	s := 0
	for s < l {
		n, err := r.Write(d[s:])
		if err != nil {
			return err
		}

		s = s + n
	}

	return nil
}
