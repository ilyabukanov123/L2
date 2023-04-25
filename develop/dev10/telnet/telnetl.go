package telnet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Args struct {
	Timeout uint
}

func request(conn net.Conn, signalChan chan<- os.Signal, connChan chan<- error) {
	for {
		// Читаем из stdin
		reader := bufio.NewReader(os.Stdin)
		payload, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				signalChan <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			connChan <- err
		}
		// Пишем в сокет
		fmt.Fprintln(conn, payload)
	}
}

func response(conn net.Conn, connChan chan<- error) {
	for {
		// Читаем ответ от сервера
		reader := bufio.NewReader(conn)
		response, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				connChan <- fmt.Errorf("connection closed by foreign host")
				return
			}
			connChan <- err
		}
		// Выводим что прочитали
		fmt.Print(response)
	}
}

func Run(host, port string, args Args) error {
	address := net.JoinHostPort(host, port)
	var conn net.Conn
	var err error
	if args.Timeout > 0 {
		conn, err = net.DialTimeout("tcp", address, time.Duration(args.Timeout)*time.Second)
	} else {
		conn, err = net.Dial("tcp", address)
	}
	if err != nil {
		return err
	}
	defer conn.Close()

	// syscall chan
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	// connection error chan
	connChan := make(chan error, 1)
	go request(conn, signalChan, connChan)
	go response(conn, connChan)

	select {
	case <-signalChan:
		conn.Close()
	case err = <-connChan:
		if err != nil {
			return err
		}
	}
	return nil
}
