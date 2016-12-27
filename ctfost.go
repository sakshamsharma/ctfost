package main

import (
	"net"
	"os"
	"strconv"

	"github.com/sakshamsharma/ctfost/logger"
	"github.com/sakshamsharma/ctfost/server"
)

var progname string

func main() {
	logger.Init()

	progname = os.Getenv("PROG_NAME")
	if progname == "" {
		progname = "program/service"
	}

	port, err := strconv.Atoi(os.Getenv("PROG_PORT"))
	if err != nil {
		port = 4002
	}

	err = server.Server{"0.0.0.0", port, "tcp"}.Listen(handler)
	if err != nil {
		logger.Error.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
}

func handler(scon *net.TCPConn) {
    defer scon.Close()

	nfile, err := scon.File()
	defer nfile.Close()

	if err != nil {
		logger.Error.Println("Error getting file from network: ", err.Error())
		return
	}

	var procattr os.ProcAttr
	procattr.Files = []*os.File{nfile, nfile, nfile}

	process, err := os.StartProcess(progname, []string{}, &procattr)

	if err != nil {
		logger.Error.Println("Start process failed:" + err.Error())
		return
	}

	_, err = process.Wait()
}
