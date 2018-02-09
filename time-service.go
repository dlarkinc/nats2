package main

import (
    "github.com/nats-io/nats"
    "fmt"
    "github.com/dlarkinc/nats2/transport"
    "github.com/golang/protobuf/proto"
    "os"
    "time"
)

// We use globals because it's a small application demonstrating NATS.

var nc *nats.Conn

func main() {

    if len(os.Args) != 2 {
        fmt.Println("Wrong number of arguments. Need NATS server address.")
        return
    }
    var err error

    nc, err = nats.Connect(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Connected to NATS server " + os.Args[1])

    nc.QueueSubscribe("TimeTeller", "TimeTellers", replyWithTime)
    select {} // Block forever
}

func replyWithTime(m *nats.Msg) {
    curTime := Transport.Time{time.Now().Format(time.RFC3339)}

    data, err := proto.Marshal(&curTime)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Replying to ", m.Reply)
    nc.Publish(m.Reply, data)

}