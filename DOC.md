## How to Use

This APIs can be accessed directly by running the service and hitting the endpoints with the above paths

#### Ways to run the service:

##### *make run*:
This will compile and run the local version of the code

##### *make docker-run*:
This will start a local container in foreground by pulling the specific version of the image

###### *curl https://localhost:3333/petstore/healthz*
###### *curl https://localhost:3333/petstore/pets* GET
###### *curl https://localhost:3333/petstore/pets/{id}* GET
###### *curl https://localhost:3333/petstore/pets* POST
###### *curl https://localhost:3333/xxx*

## Few useful commands:

##### *make docker-run*: Run the docker container

##### *make run*: Run the service using local code

##### *make docker-build*: Builds the docker image

##### *make build*: Builds the binary for the service

##### *make test*: Run the unit tests for the service endpoints

## Project Description:

### main.go:
This files is the entry point for the service. It contains the code for creating a router containing the set of routes

### pkg:
This package contains 3 files:

##### router.go:
Constructor for an http *Router* which is associated with the set of routes defined

##### routes.go
A struct *Route* represents an http Route with its related fields such as *HTTP Method*, *HTTP HandlerFunc*, *Endpoint* and the *Path*

#### svc:
This folder contains the 2 files:

##### endpoint.go:
All the service endpoints are implemented in this file. Such as *Health*, *AddPet*, *ListPet*, *GetPet*, *Reflect*

##### endpoint_test.go:
Contains the unit test functions for all the endpoints

### internal:
This folder contains the utility files such as *logger.go*, which initializes the logger for http requests

##### helpers.go:
This file contains the actual logic for encryption and decryption of the source input using ceaser cipher algorithms

### scripts:
This folder has the helper scripts used across the projects. It has *gocoverage.sh* currently,
which is used for checking the go code coverage of the project

### code config files:
##### .yamllint.conf:
Config file to be used by the yaml linter

##### .golangci.yaml:
Config file for the golangci-lint linter binary


### Dockerfile:
This is the docker file for the service. It is a multi-stage docker file where the base image is used as *golang:1.12.6-alpine*,
the service binary is built in that image and another image *alpine:latest*, is used to run the binary 