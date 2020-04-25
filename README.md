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

The app is set up to automatically hot reload when running via docker-compose. It uses ```CompileDaemon``` to check for changes to .go files.

If you would like to have the same effect when running via the docker file directly (docker run), then uncomment the ```entrypoint``` and comment out the ```cmd``` in the ```Dockerfile```

#### Multistage build

The ```Dockerfilemulti``` file creates a minimal container from the app and the Alpine linux base image. The file has two stages:

1. Creates a docker image based on the golang official base image. This is basically the same image created by the original Dockerfile. The only difference is that this image will also contain a binary called ```goWebAppLinAlp``` which is created to be able to run 
our app in a linux Alpine environment.
2. Creates a linux Alpine image that can run the application

Size biggest difference between the golang image and the Alpine image. While the golang image is 858MB, the Alpine version is only 16MB (about **98%** size reduction).

To create an image using the ```Dockerfilemulti```  multistage build, run:
```bash
docker build -t gowebappalpine -f Dockerfilemulti .      

```
To run the above image:
```bash
docker run -p 8080:8080 --name gowebappalpinecont gowebappalpine  
```
**Note**: We could technically delete the original Docker file and just use the multistage version. To create an image similar to the golang image curently created by the original Docker file using the multistage build, we could just run:
```bash
docker build --target Dev  -t gowebappfs -f Dockerfilemulti .
```