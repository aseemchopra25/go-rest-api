## Simple REST API using Golang and Gin 

### To start the server

```
https://github.com/aseemchopra25/go-rest-api.git

cd go-rest-api

go run main.go

```

### Send requests

```
------------------------------------
GET:
<All Albums>
curl http://localhost:8080/albums

<Album by ID>
curl http://localhost:8080/albums/1

------------------------------------
POST:

<Add New Album>
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

