# Othala - Product System
System to handle Goods acquisition

## Getting Started

Clone this project 

```sh
git clone git@github.com:DanielFCubides/Othala-invoice.git
```

### Prerequisites

You will need:
- Docker
- Go 


### Run the project Locally
Setup the environment variables that you will need, these variables are shown in the `.env` file.

Set up your database:

```sh
docker run --name pets -p 5432:5432 -e POSTGRES_PASSWORD=othala -e POSTGRES_USER=othala -e POSTGRES_DB=othala -d postgres
```
  
  
Build the binary:  

```sh
go build -o othala main.go
```

Run the binary:  
```sh
./othala
```

### Run the project with docker-compose

To run the project in a docker container, build the images with the docker-compose:

```sh
docker-compose build .
```

And run the containers:

```sh
docker-compose up 
```

## Running the tests

### Unit test

To run the unit test

```sh
go test ./...
```

