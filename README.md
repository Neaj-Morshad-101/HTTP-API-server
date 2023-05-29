## RESTful HTTP API server using [Go](https://github.com/golang), [Cobra CLI](https://github.com/spf13/cobra), [Go-chi](https://github.com/go-chi/chi)

### Description
This is a basic RESTful API server, build with Golang. In this API server I have implemented Cobra CLI for running the API from the CLI and also used go-chi instead of Go net/http.

---------------

### Installation
- `git clone git@github.com:Neaj-Morshad-101/HTTP-API-server.git`

## To Run the HTTP-API-server, Please follow the instructions bellow:

### Step 01: To start the api server, go to the project directory the run the following command:

`cd go/src/github.com/Neaj-Morshad-101/HTTP-API-server/`

`./HTTP-API-server start -p 8080`

Notes: You can give any port number just remember this to use it in request you send.

### Step 02: Start Postman by using the command: 
`postman`

### Step 03: Now hit the api server using the login info:
`curl -X POST -d '{"username":"Neaj Morshad","password":"1234"}' http://localhost:8080/login`

Notes:
You can do it from the postman also, select verb POST and hit http://localhost:8080/login 
in the body section: select raw and the these login data in json format

{
	"username" : "Neaj Morshad",
	"password" : "1234"
}

### Step 04: Set the token from the command line using the following command: 
`TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiTmVhaiBNb3JzaGFkIl0sImV4cCI6MTY3OTY0MjIxOX0.RMoAGj7KGFUtvq3ebpg_3ksQJqy4Q-gA5jAfOF4SCPQ"`

Notes: 
You will find this token in response section (Cookies) in Posman. 
Name: jwt
Value: TOKEN
For the posman you don’t have to set any thing. 
Go to bash, bacause ‘=’ is not supported in fish 

### Step 05: Now hit the API sever using any commands like this:
`curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8080/albums`

Notes: 
You can hit the API using postman also, Select correct verb and the API endpoints
(GET, PUT, POST, DELETE)

---------------


## Other Information:

### Some CLI Commands

- Start the API in default port : 8080 by `HTTP-API-server start`
- Start the API in your given port by `HTTP-API-server start -p=PORT_NUMBER`, give your port number in the place of PORT_NUMBER

--------------


### Credentials 
{ 
  "Neaj Morshad",
  "1234", 
}



### Run the API server in docker container using dockerfile

#### Create docker image from the dockerfile

- `docker build -t <image_name> .`
- or `docker build -t <docker_hub_username>/<image_name>:<tag> .` (if your do this then don't need to give tag before dockerhub push)

#### Build the docker image using these commands:
- `docker build -t neajmorshad/http-api-server:0.0.1 .`
- `docker push neajmorshad/http-api-server:0.0.1`

#### Run the API server from the docker image in docker container

- `docker run -p 8080:8080 <image_name>` (valid when used `CMD ["start", "-p", "8080"]` in Dockerfile)
- `docker run -p 8080:8080 <image_name> start -p "8080"` (when did not used CMD in Dockerfile)

--------------

#### Upload the image to [Docker Hub](https://hub.docker.com/)

- `docker login --username=<docker_hub_username>`
- `docker tag <id_of_the_created_image> <docker_hub_username>/<name_of_the_image>:<tag>`
- `docker push <docker_hub_username>/<name_of_the_image>:<tag>`

--------------

#### Run using volume (where did not gave .env file in docker image)


- `docker run -v <absolute_host_path/.env>:<container_path/.env> -p 8080:8080 <image_name> start -p 8080`


--------------

### The Endpoints of this REST API

WILL BE UPDATED SOON 

----------------

### Data Model

WILL BE UPDATED SOON

----------------

### JWT Authentication

- implemented JWT authentication
- first of all user need to hit `/login` endpoint with basic authentication then a token will be given and with that token for specific amount of time user can do other request
----------------

#### Run the API server

- `curl -X POST -H "Content-Type:application/json" -d '{"username":"neajmorshad","password":"1234"}' http://localhost:8080/login`

#### Do CRUD Requests: GET POST PUT DELETE (Hit any endpoint) 
----------------

### API Endpoints Testing

- Primarily tested the API endpoints by [Postman](https://github.com/postmanlabs)
- E2E Testing.
    - Checked response status code with our expected status code
