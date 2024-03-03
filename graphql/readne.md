

```sh
curl 'http://localhost:8080/product?query=\{product(id:1)\{name,info,price\}\}' | jq

curl 'http://localhost:8080/product?query=\{list\{id,name,info,price\}\}' | jq

curl 'http://localhost:8080/product?query=mutation+_\{update(id:1,price:3.97)\{id,name,info,price\}\}' | jq

curl 'http://localhost:8080/product?query=mutation+_\{delete(id:1)\{id,name,info,price\}\}'

curl 'http://localhost:8080/product?query=mutation+_\{create(name:"Inca+Kola",info:"Inca+Kola+is+a+soft+drink+that+was+created+in+Peru+in+1935+by+British+immigrant+Joseph+Robinson+Lindley+using+lemon+verbena+(wiki)",price:1.99)\{id,name,info,price\}\}' | jq



```