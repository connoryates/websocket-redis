# websocket-redis

Microservice for handling Redis pub/sub events and websocket connections.

Front-end client should make a websocket connection like so:

```javascript
var socket = new WebSocket("ws://127.0.0.1:8080/ws");

socket.onmessage = function(event) {
    console.log(event);
};
```

Which creates a Redis channel that sends all publish events back through the websocket.
