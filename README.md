# Simple static HTTP server

## Installation

```
go get github.com/markwilson/static-server/...
```

## Usage

```
# run HTTP server from the current directory on port 8080
static-server

# run HTTP server from "public" folder on port 80
static-server -d public -p 80

# run HTTP server and bind to 0.0.0.0
static-server -x

# run HTTP server and bind to custom IP
static-server -i <custom IP>
```
