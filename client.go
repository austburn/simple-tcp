package main

import(
  "net"
  "errors"
  "fmt"
  "bufio"
  "os"
)

func main() {
  localAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
  conn, err := net.DialTCP("tcp", nil, localAddr)

  if err != nil {
    errors.New("problem connecting.")
  }

  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("> ")
    msg, _ := reader.ReadBytes(0xA)

    // Kill the newline char
    msg = msg[:len(msg) - 1]

    _, err := conn.Write(msg)

    response := make([]byte, 1024)

    _, err = conn.Read(response)
    if err != nil {
      fmt.Print("Connection to the server was closed.\n")
      break
    }

    fmt.Printf("%s\n", response)
  }
}
