# network

Minimal helpers around Golang `net` and `net/http` packages.

## Installation

```bash
go get github.com/benitogf/network
```

## Usage

```go
import "github.com/benitogf/network"
```

### Validate IP addresses

```go
ok := network.IsValidIP("127.0.0.1")
```

### HTTP clients

```go
client := network.NewHttpClient()     // ~5s timeouts (sensible timeouts)
fastClient := network.NewFastHttpClient() // ~1s timeouts (fast timeouts)
```

### Check if a host is reachable

```go
reachable := network.IsHostReachable(client, "example.com:80")
```

### Local IP utilities

```go
ips := network.GetLocalIPs()
self := network.SelfIP()
```
