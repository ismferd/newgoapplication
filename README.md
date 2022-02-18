# newgoapplication

- [how to run newgoapplication](#how-to-run-newgoapplication)
    - [release newgoapplication](#release-newgoapplication)
- [how to run the tests](#how-to-run-the-tests)
- [more information about it](#more-information-about-it)
    - [what I would do next, given more time](#what-I-would-do-next,-given-more-time)
    - [are there bugs that I am aware of?](#are-there-bugs-that-I-am-aware-of?)

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

Just to agilize the development cycle I made a makefile with the following phonies:
```
 make help
 help  Display this help screen
 darwin  Build the binary for mac
 linux  Build the binary for linux
 release  Tests, package the image and push it to docker registry
 test  Launch the unit test
 clean  Remove previous build
 ```

## release newgoapplication

I added an option to build the dockerfile and push it to the docker registry but you have to set both env vars `DOCKER_PASS` and `DOCKER_USER`, then run `make release` also it will be released in each PR.

# how to run the tests

You can run unit tests locally running `make test`, this command gives you the result of the unit also the coverage
```
make tests
go test -cover -timeout 10s ./...
?       github.com/ismferd/newGoApplication     [no test files]
ok      github.com/ismferd/newGoApplication/pkg/hasher  3.057s  coverage: 96.2% of statements
ok      github.com/ismferd/newGoApplication/pkg/sanitizerwords  3.254s  coverage: 100.0% of statements
ok      github.com/ismferd/newGoApplication/pkg/sorter  3.147s  coverage: 100.0% of statements
```
Moreover, I have built a GHA worflow in order to cover lint test, this action will be triggered in each PR.
Take a look [here](https://github.com/ismferd/newgoapplication/actions/runs/1854934656)

# more information about it
## what I would do next, given more time

- To use goroutines in order to run newgoapplication using concurrency insted of using a single thread.
- More effort controlling errors.
- Move code from functions to OOP, why?
    - Code more ordered than now
    - Objects easy to extend
    - Code easy to follow the [SOLID principle](https://en.wikipedia.org/wiki/SOLID)
    - Finally, easy to test and also to mock
- The application is using a map to populate the data, this map is created in-memory, which means that we big files we can have memory limits. To improve that, maybe we can use an external engine, I was thinking of Redis or ElasticSeach, this last one can give us the advantage of their search engine.
- On the Hasher function find a way to move positions more than a real FIFO, now it is working in an weird way
```
	hasher[0] = hasher[1]
	hasher[1] = hasher[2]
```
- Observability and monitoring, add metrics, different log levels such as[ DEBUG, ERROR, WARN, and so on] and tracing. Then send them to a third-party tool that allows to take advantage of the available data in an efficient way
- To improve the unit tests, add unit test and make a kind of load test in order to see the behavior with load (large files)
- Improve the output and create a helper to know how `newgoapplication` works.

## are there bugs that I am aware of?

- I don't control the size and type of the files, it could be a real problem in production such as: consuming memory and having possibles OOM or problems with file descriptors.
- It doesn't manage the Unicode characters
- Passing 2 or more files you would see the top 100 per each file (it doesn't merge files)
