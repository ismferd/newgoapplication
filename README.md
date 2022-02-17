# newgoapplication

newgoapplication is a program executable from the command line that when given text(s) will return a list of the 100 most common three word sequences.

# how to run newgoapplication

You can run `newgoapplication` basically on 3 ways:

* Calling the main function:
```
go run main.go sample/example.txt
cat sample/example.txt | go run cmd/main.go
```
* Building the binary and passing files as params
```
new-go-application-darwin-amd64 sample/example.txt
cat sample/example.txt | new-go-application-darwin-amd64
```
* And using containers
```
docker run -v "$(pwd)"/examples:/app/examples newapp /app/examples/pg.txt
```

Just to agilize the cycle of development I made a makefile with the next phonies:
```
 make help
 help  Display this help screen
 darwin  Build the binary for mac
 linux  Build the binary for linux
 release  Tests, package the image and push it to docker registry
 test  Launch the unit test
 clean  Remove previous build
 ```
