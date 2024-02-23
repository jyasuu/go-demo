# Go Demo


https://ithelp.ithome.com.tw/users/20107343/ironman/1892

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


```

