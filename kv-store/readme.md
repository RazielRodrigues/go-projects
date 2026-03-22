# KV STORE EM GO

com base na serie de videos do huncoding: https://github.com/HunCoding/redis-like-go/tree/main

# Em um terminal, inicie o servidor:
go build -o server ./cmd/server && ./server --port 6379

# Em outro terminal, use netcat:
echo "SET mykey myvalue" | nc localhost 6379
echo "GET mykey" | nc localhost 6379
echo "EXISTS mykey" | nc localhost 6379
echo "PING" | nc localhost 6379
echo "INFO" | nc localhost 6379