package main

import (
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
)

var r *rand.Rand

func main() {
	Init()

	port, err := strconv.Atoi(os.Getenv("PROG_PORT"))
	if err != nil {
		port = 4002
	}

	r = rand.New(rand.NewSource(99))

	// Cannot be done in the docker build process
	exec.Command("cgconfigparser", "-l", "/etc/cgconfig.conf").Run()

	err = Server{"0.0.0.0", port, "tcp"}.Listen(handler)
	if err != nil {
		Error.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
}

func handler(scon *net.TCPConn) {
	defer scon.Close()

	nfile, err := scon.File()
	defer nfile.Close()

	if err != nil {
		Error.Println("Error getting file from network: ", err.Error())
		return
	}

	userId := r.Int31()%40000 + 2000

	exec.Command("user-create.sh", string(userId)).Run()
	defer exec.Command("user-delete.sh", string(userId)).Run()

	var procattr os.ProcAttr
	procattr.Files = []*os.File{nfile, nfile, nfile}

	process, err := os.StartProcess("user-run.sh", []string{}, &procattr)

	if err != nil {
		Error.Println("Start process failed:" + err.Error())
		return
	}

	_, err = process.Wait()
}
