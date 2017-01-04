package server

import (
	"net"
	"os"
	"strconv"

	"../logger"
)

type Server struct {
	HostName       string
	PortNumber     int
	ConnectionType string
}

func (s Server) Listen(handler func(*net.TCPConn)) error {
	logger.Init()

	addr := net.TCPAddr{
		IP:   net.ParseIP(s.HostName),
		Port: s.PortNumber,
	}

	socket, err :=
		net.ListenTCP(s.ConnectionType, &addr)

	if err != nil {
		logger.Error.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	defer socket.Close()

	logger.Info.Println("Listening on " + s.HostName + ":" + strconv.Itoa(s.PortNumber))
	for {
		// Listen for an incoming connection.
		conn, err := socket.AcceptTCP()
		if err != nil {
			return err
		}

		// Handle connections in a new goroutine.
		go handler(conn)
	}
}
