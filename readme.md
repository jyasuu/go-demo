# Go Demo


https://ithelp.ithome.com.tw/users/20107343/ironman/1892

https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
https://www.mongodb.com/docs/drivers/go/current/fundamentals/connections/connection-guide/

https://redis.io/docs/connect/clients/go/
https://github.com/redis/go-redis

```sh

go mod init example.com/m

# generate unit test
go get -u github.com/cweill/gotests/...
gotests -all -w main.go main_test.go
go test -v -cover=true main_test.go main.go

# benchmark 
go test -bench=.  .

# gin framework
go get -u github.com/gin-gonic/gin

# gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go doc gorm.DB

# test
curl localhost:3000/ping | jq
curl Daniel:123456@localhost:3000/hello/Daniel/play\?firstnam=Sam\&lastname=Lucas -s | jq


```

