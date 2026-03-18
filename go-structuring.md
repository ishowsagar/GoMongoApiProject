# GO LANGUAGE FOLDER STRUCTURING

1. ENTRY POINT -> **go app is build from cmd/api~server/main.go ==> this is where your main func init the go app.**

2. DB layer -> **db.go folder comprises of db.go files which pools connection to the database and stores all the methods needed for db connection or test checks.**

3. Services layer -> **this layer contains go file which comprises of all the methods which uses that db connection to either make queries for db call or all the methods needed for querying the database.**

4. Handler/Controller layer -> **This is the most important layer's go file which comprises of all the methods that invoking those db querying methods by passing required payloads they expects like id,body etc/. to execute those methods to query database and lastly make an actual response to the client.**

5. Router layer --> this layer makes use of handler methods and bind to the route paths to let client invoke those methods
   e.g => **_ Client makes an request on route path "/api/todos/all", as this would method bounded to it in router file --> which invokes handler methods --> which pass payload from client to db layer methods for querying the database --> which is connected by db pooled connection from db layer <---> all this ultimately is served from cmd/api/main.go file where server exists from http package and router handler is passed to it to route the requests. _**

## PROJECT FOLDER STRUCTURE

GoMongoApiProject/
├── cmd/
│ └── api/
│ └── main.go
├── db/
│ └── db.go
├── handlers/
│ ├── handlers.go
│ └── router.go
├── services/
│ ├── services.go
│ └── todo.go
├── docker-compose.yml
├── go.mod
├── makefile
└── ytmongotodos
