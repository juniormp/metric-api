# metric-api

The metric-api is based on Domain Driven Design and Clean Architecture practices following TDD and Test Fist approach.


**Tech stack**

- Golang
- Framework [gin-gonic](https://github.com/gin-gonic/gin)
- Docker
- Redis



### Objective

- [X] Save metrics and values 
- [X] Create a report based on the last hour

### Endpoints


Create a metric
```
Request
POST /metric/{key}
{
  "value": decimal
}

Response
Status 201
{}
```
 
 Get the metric report
 ```
 Request
GET /metric/{key}/sum

Response
{
  "value" : 20
}
```

Clean old time metrics
 ```
 Request
GET /clean-metrics

Response
Status 200
{}
```

### Project Setup

To build project image

 `$ docker build -t metric-api -f Dockerfile.dev .`
 
To start de project

`$ docker-compose up -d`

To run the tests

`$ go test ./tests/...`
