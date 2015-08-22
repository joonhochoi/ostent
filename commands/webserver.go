package commands

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ostrost/ostent/commands/extpoints"
	"github.com/ostrost/ostent/flags"
)

type webserver struct {
	Log          *extpoints.Log
	Bind         flags.Bind
	ServeFunc    func(net.Listener)
	FirstRunFunc func() bool
	ShutdownFunc func() bool
}

func (wr webserver) NetListen() net.Listener {
	listen, err := net.Listen("tcp", wr.Bind.String())
	if err != nil {
		wr.Log.Fatal(err)
	}
	return listen
}

// LogInit sets up global log
func InitStdLog() {
	log.SetPrefix(fmt.Sprintf("[%d][ostent] ", os.Getpid()))
	// goagain logging is useless without pid prefix
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func NewWebserver(defport int) *webserver {
	return &webserver{
		Log: NewLog(fmt.Sprintf("[%d][ostent webserver] ", os.Getpid()), func(l *extpoints.Log) {
			l.Flag |= log.Lmicroseconds
		}),
		Bind: flags.NewBind(defport),
	}
}

func (ws *webserver) AddCommandLine() *webserver {
	AddCommandLine(func(cli *flag.FlagSet) extpoints.CommandLineHandler {
		cli.Var(&ws.Bind, "b", "short for bind")
		cli.Var(&ws.Bind, "bind", "Bind `address`")
		return nil
	})
	return ws
}
