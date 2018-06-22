# Go-Api

Simple REST API written in Golang using MongoDB.


| Method     | URI        | Description           |
|------------|------------|-----------------------|
|**GET**     | /album     | List of music albums  |
|**GET**     | /album/:id | Get a single album    |
|**POST**    | /album     | Add a new album       |
|**PUT**     | /album/:id | Update an album by ID |
|**DELETE**  | /album/:id | Delete an album by ID |

## Run application

```
# Make sure "dep" is installed
go get -u github.com/golang/dep/cmd/dep

dep ensure

go run main.go
```

Note: This API defines the MongoDB connection details under `album/repository.go`
```
const (
	SERVER  = "127.0.0.1:27017"
	DBNAME  = "musicstore"
	DOCNAME = "albums"
)
```