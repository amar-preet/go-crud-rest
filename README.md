# go-crud-rest

A simple CRUD API using golang and gin with storage on postgreSQL 

<details>
<summary>
GET Albums
</summary>

`curl localhost:8080/albums`
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

`
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "Fire","artist": "Tiesto","price": 29}'
`
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

`curl -X PUT -d '{"price" : 12}' localhost:8080/albums/1`
</details>
