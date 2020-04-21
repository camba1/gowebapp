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