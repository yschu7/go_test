package main

import (
    "fmt"
    "net"
    "net/rpc"
    "math/rand"
    "time"
    "flag"
    "os"
)

type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
    *reply = -i
    return nil
}

var PORT = ":9999"

func server() {
    rpc.Register(new(Server))
    ln, err := net.Listen("tcp", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        c, err := ln.Accept()
        if err != nil {
            continue
        }
        go rpc.ServeConn(c)
    }
}
func client() {
    svrip := "127.0.0.1"
    if flag.NArg() > 0 {
      svrip = flag.Args()[0]
      fmt.Println("Server IP: ", svrip)
    }
    c, err := rpc.Dial("tcp", svrip + PORT)
    if err != nil {
        fmt.Println(err)
        return
    }
    var rn, result int64
    rand.Seed(time.Now().Unix())
    for {
      rn = rand.Int63n(100000)
      err = c.Call("Server.Negate", rn, &result)
      if err != nil {
          fmt.Println(err)
          break
      } else {
          fmt.Printf("Server.Negate(%d) = %d\n", rn, result)
          fmt.Println("Go to sleep... ")
          time.Sleep(3000 * time.Millisecond)
      }
    }
}

func main() {
    var cmode, smode bool
    cmode, smode = setFlag()
    if cmode {
      go client()
    } else if smode {
      go server()
    } else {
      fmt.Println("Please specify client/server mode")
      os.Exit(1)
    }
    var input string
    fmt.Println("Press Enter to exit...")
    fmt.Scanln(&input)
}

func usage() {
  fmt.Fprintf(os.Stderr, "usage: %s -c|-s [server]\n", os.Args[0])
  flag.PrintDefaults()
  os.Exit(2)
}

func setFlag()(c, s bool) {
  flag.Usage = usage
  flag.BoolVar(&c, "c", false, "Client mode")
  flag.BoolVar(&s, "s", false, "Server mode")
  flag.Parse()
  return c, s
}
