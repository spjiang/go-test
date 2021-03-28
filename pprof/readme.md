#说明

###图形安装
`brew install graphviz`
### Heap profile
``go tool pprof --text http://localhost:8080/debug/pprof/heap``
###CPU profile
`go tool pprof --text http://localhost:8080/debug/pprof/profile
`
### 
```go tool pprof -http=:8081 http://localhost:6060/debug/pprof/goroutine\?debug\=1```
å`go tool pprof http://localhost:6060/debug/pprof/block`