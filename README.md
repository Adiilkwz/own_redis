# Own Redis - UDP Key-Value Store

A lightweight, high-performance, in-memory key-value store written in Go. Unlike traditional Redis which uses TCP, this application communicates exclusively over the **UDP protocol**, making it incredibly fast and stateless. 

The project is designed to successfully handle competitive, parallel client requests. It spins up independent goroutines for incoming network packets and protects the underlying data map with a `sync.RWMutex`, ensuring thread safety and completely preventing data races. It also supports key expiration (TTL) out of the box.

## Project Structure

```text
own-redis/
├── main.go               # Server entry point
├── cmd/
│   └── client.go         # Custom interactive UDP client
├── internal/
│   ├── config/
│   │   └── config.go     # CLI Flag parsing (--port, --help)
│   ├── protocol/
│   │   └── handler.go    # Command parsing logic
│   ├── server/
│   │   └── udp.go        # UDP networking loop and Goroutine dispatcher
│   └── storage/
│       ├── engine.go     # Thread-safe map using sync.RWMutex
│       └── item.go       # TTL (Time-To-Live) expiration logic
└── go.mod                # Go module definition
```

## How to Build and Run

### 1. Build the Server
Open your terminal, navigate to the root directory of the project, and compile the server executable:
```bash
go build -o own-redis .
```

### 2. Start the Server
Start the server by providing a port number using the `--port` flag:
```bash
./own-redis --port [port_number]
```
*(You can view the exact usage instructions at any time by running `./own-redis --help`)*

### 3. Connect a Client
Leave the server terminal running. Open a **new terminal window**. 

Because standard UDP clients like `nc` (Netcat) are not always installed on every operating system by default, this project includes a custom interactive Go client to easily test the server.

Run the client and point it to the same port as your server:
```bash
go run cmd/client.go --port [port_number]
```

Once connected, you can type your commands directly into the client prompt to interact with the key-value store.