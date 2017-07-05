// Project site: https://github.com/gotmc/visa
// Copyright (c) 2017 The visa developers. All rights reserved.
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tcpip

import (
	"fmt"
	"net"

	"github.com/gotmc/visa"
)

// tcpipDriver implements the ivi.Driver interface.
type Driver struct{}

func (d Driver) Open(address string) (visa.Resource, error) {
	var c Connection
	tcpAddress, err := getTcpAddress(address)
	if err != nil {
		return nil, fmt.Errorf("%s is not a TCPIP VISA resource address: %s", address, err)
	}
	c.conn, err = net.Dial("tcp", tcpAddress)
	if err != nil {
		return nil, fmt.Errorf("Problem connecting to TCP instrument at %s: %s", tcpAddress, err)
	}
	return &c, nil
}

type Connection struct {
	conn net.Conn
}

func (c *Connection) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (c *Connection) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}

func (c *Connection) WriteString(s string) (int, error) {
	return 0, nil
}

func (c *Connection) Query(s string) (value string, err error) {
	return "foo", nil
}

func getTcpAddress(address string) (string, error) {
	return "127.0.0.1:5025", nil
}

// init registers the driver with the program.
func init() {
	var driver Driver
	visa.Register(visa.TCPIP, driver)
}