# Leaves Module

To run the code you have to follow the instructions below:


## Pre-Requisites:

### Docker:

Install Docker is not already installed on your machine using this link:
[Docker](https://docs.docker.com/install/)

### Docker-Compose:

Install Docker-Compose is not already installed on your machine using this link:
[Docker-Compose](https://docs.docker.com/compose/install/)

---

## Clone Repo
1. Make a Directory using following command:

```bash
$ mkdir $GOPATH/github.com/reemishirsath
```
2.Navigate that Directory using following command:
```bash
$ cd $GOPATH/github.com/reemishirsath/
```
3.Clone the Repo using following command:
```bash
$ git clone https://github.com/reemishirsath/leave-module.git
```

4.Navigate to the code:
```bash
$ cd $GOPATH/github.com/reemishirsath/leave-module
```
---


## Running The Code using Docker Containers:

```bash
$ docker-compose up
```
#### Note: The above command will take time the first time to run.

## Postman Collection:

Postman collection is present in the postman folder, which can be imported in Postman applicaion.
---

# Project Architecture:

```bash
.
├── database                                           <-------------Directory of Database files
│   ├── Dockerfile                                     <-------------Dockerfile of Database Container
│   ├── leaves_module.sql 
│   └── mysql.go
├── docker-compose.yaml                                <-------------Docker-Compose file to run all the Containers at once
├── Dockerfile_Leaves                                  <-------------Dockerfile of Leaves-Service Container
├── Dockerfile_User                                    <-------------Dockerfile of User-Service Container
├── .gitignore
├── .gitlab-ci.yml                                     <-------------.gitlab-ci.yml file to enable CI Pipeline
├── go.mod                                             <-------------go.mod Go inbuilt dependancy management file(only works with  go-v1.11 & above)
├── go.sum                                             <-------------go.sum Go inbuilt dependancy management file(only works with  go-v1.11 & above)
├── logs
│   └── log.go
├── models                                             <-------------Directory of Database Models
│   ├── http_err.go
│   ├── leaves.go
│   ├── roles.go
│   └── users.go
├── README.md                                          <-------------The file you are reading now.
└── services                                           <-------------Directory of Services.
    ├── leave                                          <-------------Leave-Service.
    │   ├── cmd
    │   │   ├── main.go
    │   │   └── service
    │   │       └── service.go
    │   ├── Makefile                                   <-------------Makefile can be used alternatively to test code with the Docker Container.
    │   └── pkg
    │       └── v1
    │           ├── endpoints
    │           │   └── router.go
    │           ├── handlers
    │           │   ├── helpers.go
    │           │   ├── leave_handler.go
    │           ├── middleware
    │           │   └── middleware.go
    │           ├── repositories
    │           │   └── leave_repositories.go
    │           └── services
    │               └── leave_service.go
    └── user                                           <-------------User-Service.
        ├── cmd
        │   ├── main.go
        │   └── service
        │       └── service.go
        ├── Makefile                                   <-------------Makefile can be used alternatively to test code with the Docker Container.
        ├── mocks
        │   ├── user_repository_mock.go
        │   └── user_service_mock.go
        └── pkg
            └── v1
                ├── endpoints
                │   └── router.go
                ├── handlers
                │   ├── helpers.go
                │   ├── user_handler.go
                │   └── user_handler_test.go
                ├── middleware
                │   └── middleware.go
                ├── repositories
                │   └── user_repositories.go
                └── services
                    └── user_service.go
```