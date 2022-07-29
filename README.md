# go-crud-rest

A simple CRUD API using `go` and `gin` with storage on postgreSQL

## Running the Service locally

Make sure local PostgreSQL database is running. 
Run
```
go mod tidy
make run
```


## Testing the endpoints

<details>
<summary>
GET Albums
</summary>

`curl http://localhost:8080/albums`
</details>

<details>
<summary>
GET Album By ID
</summary>

`curl http://localhost:8080/albums/2`
</details>

<details>
<summary>
POST an Album
</summary>

```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "Fire","artist": "Tiesto","price": 29}'
```
</details>

<details>
<summary>
DELETE Album By ID
</summary>

`curl -X DELETE http://localhost:8080/albums/2`
</details>

<details>
<summary>
UPDATE Album By ID
</summary>

`curl -X PUT -d '{"price" : 12}' http://localhost:8080/albums/1`
</details>
