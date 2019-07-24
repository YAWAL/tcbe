<h2>Prerequisites: </h3>

Go - https://golang.org/

Mongo - https://www.mongodb.com/

Dep - https://golang.github.io/dep/

<h2>Run programm: </h3>

1. Unzip archive and run `dep ensure` from root folder
2. Start MongoDB on `mongodb://127.0.0.1:27017` or choose your own address and fill it in config.json file
3. Run `go run main.go`

Your application is running on port `http://localhost/:8888`. For checking functionality you can use postman collection `tcbe.postman_collection.json`