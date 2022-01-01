# go-postgres
API server that supports the followoing to access Vehicles table in the Postgres DB
- CRUD operations
- Search operation

This server supports CRUD for the following formats:
- Json
- Xml
- Protobuf

## Expectation

The client is expected to set "Content-type" header in the request
- application/json := for JSON
- application/xml := for XML
- protobuf := for protobuf

Any other format is not supported.

Cache Headers
==============

The support for cache headers is added for GET requests. In case the request is cached, the browser should send 304.

## P.S.

The application server uses parameters in .env file to connect to POSTGRES DB

Dockerfile
==========

.env file : To enable users to configure postgres connection parameters

## Example command
sudo docker run -p 8080:8080 -d --name my_pgserver --network="host" pgserver

Testing
========
- used Postman to test APIs

## Unit tests 
Unit tests have been added for handlers with coverage > 70%

### Snippet
go test -coverprofile=coverage.o

coverage: 73.5% of statements
SUCCESS	go-postgres/handlers	0.554s

Unit Bench is added for Create Request

### Snippet

go test -run BenchmarkCreateVehicle10 -bench=.

PASS
ok  	go-postgres/handlers	1.857s

# Endpoints

## Create:  POST /api/newvehicle
Json Body :
  {
    "vin": "FDHFJ",
    "make": "Toyota",
    "model": "Camry",
    "color": "blue",
    "type": "car",
    "condition": "new"
}

## Read :  GET /api/vehicle/{id}
        GET /api/vehicle    //GET ALL VEHICLES
        

## Update: PUT /api/vehicle/{id}
Json Body :
{
    "vin": "FDHFJ",
    "make": "Toyota",
    "model": "Camry",
    "color": "blue",
    "type": "car",
    "condition": "new"
}

## Delete: DELETE /api/deletevehicle/{id}

Search : GET "/api/search/{query}
- Example: "/api/search/vehicle?color=black"


# Limitations

Search Operation only returns 1 query result


DATABASE script amd SQL file 
===============================
These files are contained under sql folder

client.py
==========

Script to exercise REST endpoints

Following installations needed to run the script:

apt-get install python-requests
sudo apt-get install -y python-simplejson






