package main

import (
	"log"
	"net"
	"nvim-share/server"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	cmd := exec.Command("nvim")
	tty, err := pty.Start(cmd)
	if err != nil {
		log.Fatalf("Falha ao iniciar pty para nvim: %v", err)
	}
	defer func() {
		tty.Close()
		cmd.Wait()
	}()

	log.Println("Neovim iniciado em um pseudo-terminal.")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
	defer listener.Close()
	log.Println("Servidor TCP iniciado na porta :8080, aguardando conexão...")

	server.Run(listener, tty)
}
