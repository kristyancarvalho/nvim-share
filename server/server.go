package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var (
	activeConnections int32
	connMu            sync.Mutex
)

func handleConnection(conn net.Conn, tty io.ReadWriteCloser) {
	defer func() {
		connMu.Lock()
		activeConnections--
		connMu.Unlock()
		log.Printf("Cliente desconectado. Conexões ativas: %d", activeConnections)
		conn.Close()
	}()

	connMu.Lock()
	activeConnections++
	log.Printf("Novo cliente conectado. Conexões ativas: %d", activeConnections)
	connMu.Unlock()

	go func() {
		if _, err := io.Copy(tty, conn); err != nil && err != io.EOF {
			log.Printf("Erro de escrita para pty: %v", err)
		}
	}()

	if _, err := io.Copy(conn, tty); err != nil && err != io.EOF {
		log.Printf("Erro de escrita para cliente: %v", err)
	}
}

func Run(listener net.Listener, tty io.ReadWriteCloser) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Erro ao aceitar conexão: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if activeConnections >= 1 {
			log.Println("Já existe um cliente conectado. Rejeitando nova conexão.")
			conn.Write([]byte("Desculpe, já há uma sessão Neovim ativa. Tente novamente mais tarde.\n"))
			conn.Close()
			continue
		}

		go handleConnection(conn, tty)
		fmt.Fprintf(os.Stdout, "Sessão Neovim compartilhada iniciada.\n")
	}
}
