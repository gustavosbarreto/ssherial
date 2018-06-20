package main

import (
	"io"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/pkg/term"
)

func sessionHandler(session ssh.Session) {
	_, _, isPty := session.Pty()
	if !isPty {
		return
	}

	terminal, _ := term.Open(os.Args[1], term.RawMode)
	term.Speed(115200)

	go func() {
		io.Copy(terminal, session)
	}()

	io.Copy(session, terminal)
}

func checkPassword(ctx ssh.Context, pass string) bool {
	return true
}

func main() {
	ssh.Handle(sessionHandler)
	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.PasswordAuth(checkPassword)))
}
