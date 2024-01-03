# my-disco

### run the code

1. Begin tracking the Gin module as a dependency.

At the command line, use go get to add the github.com/gin-gonic/gin module as a dependency for your module. Use a dot argument to mean “get dependencies for code in the current directory.”

```
$ go get .
go get: added github.com/gin-gonic/gin v1.7.2
```

Go resolved and downloaded this dependency to satisfy the import declaration you added in the previous step.

2. From the command line in the directory containing main.go, run the code. Use a dot argument to mean “run code in the current directory.”

```
$ go run .
```

Once the code is running, you have a running HTTP server to which you can send requests.

3. From a new command line window, use curl to make a request to your running web service.

```
$ curl http://localhost:8080/albums
```