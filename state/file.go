package state

import (
	"errors"
	"fmt"
	"github.com/stated/stated/parser"
	"log"
	"net"
	"net/http"
	"net/rpc"
	// "os"
	"os/exec"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

type File int

func Listen() {
	arith := new(Arith)
	file := new(File)
	rpc.Register(arith)
	rpc.Register(file)
	rpc.HandleHTTP()
	if l, err := net.Listen("tcp", "127.0.0.1:12789"); err != nil {
		log.Fatal("listen error:", err)
	} else {
		http.Serve(l, nil)
	}
}

func Serve(stated parser.Stated) (cmd *exec.Cmd) {
	plugin := "/home/baiju/mygo/src/github.com/stated/stated/stated"
	cmd = exec.Command(plugin, "-plugin", "test")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return cmd
}

func Client(stated parser.Stated) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:12789")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
