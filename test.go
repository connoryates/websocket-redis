package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "github.com/gorilla/websocket"
)

var (
    gRedisConn  = func() (redis.Conn, error) {
        return redis.Dial("tcp", ":6379")
    }
)

const(
    testMessage = "Hello"
    testID = 123456
)

func connectWS() *websocket.Conn {
    // Returning 404
    u := "ws://127.0.0.1:8080?id=" + testID

    // Debug
    conn, r, err := websocket.DefaultDialer.Dial(u, nil)
    if err != nil {
        fmt.Println(r)
        panic(err)
    }

    return conn
}

func main() {
    go func() {
        if c, err := gRedisConn(); err != nil {
            fmt.Println("Error establishing Redis connection")
        } else {
            c.Do("PUBLISH", testID, testMessage)
        }
    }()

    ws := connectWS()
    for {
        _, resp, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("read:", err)
            return
        }

        fmt.Println("Got value", resp)
    }
}

