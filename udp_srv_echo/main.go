package main

import (
  "fmt"
  "log"
  "net"
  "os"
)

func check (e error) {
  if e != nil {
    log.Fatal(e)
    os.Exit(1)
  }
}

func main () {
  if len(os.Args) < 2 {
    fmt.Println("Please provide address and port for UDP Server in format \"ADDRESS:PORT\".")
    os.Exit(1)
  }
  
  udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])
  check(err)

  conn, err := net.ListenUDP("udp", udpAddr)
  check(err)

  for {
    var buf [512]byte
    
    _, addr, err := conn.ReadFromUDP(buf[0:])
    check(err)

    fmt.Println("> ", string(buf[0:]))

    n, _, err := conn.WriteMsgUDP(buf[0:], nil, addr)
    check(err)
    fmt.Printf("Send %v byte message.\n", n)
  }
}
