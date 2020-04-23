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
### Docker

The application has been setup to be able to run using docker and docker-compose.

To build the docker image:
```bash
docker build -t gowebapp .
```

To run the app using the docker file:

```bash
docker run -p 8080:8080 --name gowebappcont gowebapp
```

For docker-compose, things got a bit more interesting since the host machine is running MacOS while the image is based on the golang linux image.
As such, wehn we mount the volume from the Mac to th elinux image, the application binary is not compatible with one of the two environments.
To get around this, there are two options:

- If you do not care to share the code between the host and the guest OS, then simply comment the lines in the docker-compose fiie and you can run normally using ```docker-compose up``` and ```docker-compose down```
```bash
   volumes:
      - .:/go/src/goWebApp
```

- If you want to be able to share the code, then use the commands below instead of docker-compose up and docker-compose down. The commands below will create/delete linux binaries of the application (goWebAppLin) so that they can be run in the container normally.
```bash
./mycomposeup
```
```bash
./mycomposedown
```

#### Hot reload

The app is setup to automatically hot reload when running via docker-compose. It uses ```CompileDaemon``` to check for changes to .go files.

If you would like to have the same effect when running via the docker file directly (docker run), then uncomment the ```entrypoint``` and comment out the ```cmd``` in the ```Dockerfile```

