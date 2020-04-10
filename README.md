# truticket_metric_ticker

Offers a quick way to track counter metrics reported with a key

current support:
only able to track a metric count reported hourly basis based on a metric name/identifier 

## file structure

```$xslt
cmd/webapp/main.go          main file to initiate service 
cmd/webapp/routes/routes.go handles all the route configurations
internal/metrics/counter.go handles the metric reporting
```


## Getting Started

```$xslt
# get the code base using
git clone https://github.com/ckalagara/truticket_metric_ticker.git

# go the main file
cd cmd/webapp

# build the app
go build

# run it
./webapp

eg:

T-MacBook-Pro:webapp t$ go build
T-MacBook-Pro:webapp t$ ./webapp 
INFO[0000] Initializing the service...                  
INFO[0000] Initializing the routes...                   
INFO[0000] Initializing http server over port :3333...  

```

## Running the tests

```$xslt
go test

eg:

T-MacBook-Pro:truticket_metric_ticker t$ cd internal/metrics/
T-MacBook-Pro:metrics t$ go test
PASS
ok      github.com/ckalagara/truticket_metric_ticker/internal/metrics   0.280s


```

# Dockerfile
### Docker build
```
docker build -t=github.com/ckalagara/truticket_metric_ticker:latest .
```

### Running docker image
```
docker run -e APP_PORT=":3333" github.com/ckalagara/truticket_metric_ticker:latest
```

## Usage

### add or update a metric

to add a metric to the system, call service using your metric name and body with value you want to increment.

```$xslt
POST http://localhost:3333/metric/{metric Name, eg: users_logged_in}

HEADERS
Content-Type: application/json

BODY:
{
"value": {metric value, eg: 20}
}

RESPONSE:
{
}
```

### get a metric

to get current sum of your reported metric in current hour, call using below sample

```$xslt
GET http://localhost:3333/metric/{metric Name, eg: users_logged_in}/sum

RESPOSNE:
{
"value": 20
}
```