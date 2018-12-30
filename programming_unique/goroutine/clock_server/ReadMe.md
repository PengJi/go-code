并发时钟服务器  
顺序时钟服务器，它以每秒一次的频率向客户端发送当前时间  

```$xslt
go build clock.go
./clock &
nc localhost 8000
```

```$xslt
go build netcat.go
./netcat
```

```$xslt
go build clock_go.go
./clock_go &
```
