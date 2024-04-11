# goexpert-multithreading

## goroutines-with-waitgroups
``` shell
go run goroutines-with-waitgroups/main.go
```

## race-condition-mutex
``` shell
go run race-condition-mutex/main.go
ab -n 10000 -c 100 http://localhost:8080/mutex
curl http://localhost:8080/mutex
```

## race-condition-atomic
``` shell
go run race-condition-atomic/main.go
ab -n 10000 -c 100 http://localhost:8080/atomic
curl http://localhost:8080/atomic
```

## range-with-channels
``` shell
go run range-with-channels/main.go
```

## range-with-waitgroups
``` shell
go run range-with-waitgroups/main.go
```

## channel-directions
``` shell
go run channel-directions/main.go
```

## loadbalancer
``` shell
go run loadbalancer/main.go
```

## select-with-channels
``` shell
go run select-with-channels/main.go
```

## channel-with-buffers
``` shell
go run channel-with-buffers/main.go
```
