# websocket-redis

Requires a running Redis instance.

Microservice for handling Redis pub/sub events and websocket connections.

Front-end client should make a websocket connection like so:

```javascript
var socket = new WebSocket("ws://127.0.0.1:8080/ws");

socket.onmessage = function(event) {
    console.log(event);
};
```

Which creates and subscribes to a Redis channel and sends all events back to the client.

# Running

As a Docker container:

```bash
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
docker build -t yourname/server .
```
