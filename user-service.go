package main

import (
    "fmt"
    "os"

    "github.com/dlarkinc/nats2/transport"
    "github.com/nats-io/nats"
    "github.com/golang/protobuf/proto"
)

var users map[string]string
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

    users = make(map[string]string)
    users["1"] = "Bob"
    users["2"] = "John"
    users["3"] = "Dan"
    users["4"] = "Kate"

    nc.QueueSubscribe("UserNameById", "userNameByIdProviders", replyWithUserId)
    select {}
}

func replyWithUserId(m *nats.Msg) {

    myUser := Transport.User{}
    err := proto.Unmarshal(m.Data, &myUser)
    if err != nil {
        fmt.Println(err)
        return
    }

    myUser.Name = users[myUser.Id]
    data, err := proto.Marshal(&myUser)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("Replying to ", m.Reply)
    nc.Publish(m.Reply, data)
}