# WebApp

Small web application that allows text file manipulation. User can:

- List the files stored in a fix directory
- Create new files
- Edit new files

The backend is built with Go and the front end with HTML go templates along with Bulma CSS.

### Running the app

To run the app:

- Clone repo
- cd to the repo directory
- Build the app:
```bash
    go build
```
- call the executable
```bash
    webApp
```
- In your browser, go to:
```bash
    http://localhost:8080/
```
Note that depending on your system, you may be asked to allow the application to accept connections.

### Testing the app

To run all the packages within the app:

- cd to the root of the repo
- Execute tests:
```bash
go test ./...
```

To run the tests for a particular package:
- cd to the directory containing the package
- Execute test:
```bash
go test
```

Finally, if you want more details, you can use the -json flag eg:
```bash
go test -json
``` 

### Documentation
In typical Go manner the documentation is extracted directly from the comments in the application.

#### CLI

To view the high level documentation in the cli, run:
```bash
go doc
```
To view the documentation for publicly exposed functions, run:
```bash
go doc <function name>
```
To view the documentation for private functions, run:
```bash
go doc -u <function name>
```
To view all the documentation except for private functions
```bash
go doc -all
```
To view all the documentation including private functions
```bash
go doc -u - all
```