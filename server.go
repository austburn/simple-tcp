package main

import(
  "net"
  "errors"
  "time"
  "fmt"
)

func handleConnection(c *net.TCPConn) {
  fmt.Printf("Receving connection from %s\n", c.RemoteAddr())
  const layout = "Jan 2, 2006 at 3:04pm (MST)"
  for {
    msg := make([]byte, 32)
    c.Read(msg)

    fmt.Fprintf(c, "Echoing back %s at %s", msg, time.Now().Format(layout))
  }
}

func main() {
  tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
  ln, err := net.ListenTCP("tcp", tcpAddr)
  if err != nil {
    errors.New("problem setting up server")
  }

  for {
    conn, err := ln.AcceptTCP()

    if err != nil {
      errors.New("error connectin")
    }

    defer conn.Close()
    go handleConnection(conn)
  }
}
