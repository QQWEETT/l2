package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var timeout string

func init() {
	flag.StringVar(&timeout, "timeout", "10s", "time limit to establish connection")
	flag.Parse()
}

func main() {
	re := regexp.MustCompile(`\ds`)
	if !re.MatchString(timeout) {
		log.Fatal("invalid timeout format: " + timeout)
	}

	if len(flag.Args()) < 2 {
		log.Fatal("usage: --timeout=1s host port")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	toInt, _ := strconv.Atoi(timeout[:len(timeout)-1])
	to := time.Duration(toInt) * time.Second

	var conn net.Conn
	var err error

	start := time.Now()
	for time.Since(start) < to {
		conn, err = net.Dial("tcp", host+":"+port)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("unable to establish connection after timeout: %v", to)
	}
	defer conn.Close()
	log.Printf("connected to %s:%s", host, port)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("Message from server: " + message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		_, err := fmt.Fprintf(conn, in+"\n")
		if err != nil {
			log.Fatal("connection closed")
		}
	}
}
