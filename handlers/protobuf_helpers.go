package handlers

import (
    //"database/sql"
    "fmt"
    "net/http"
    proto "github.com/golang/protobuf/proto"
    //prot "google.golang.org/protobuf/proto"
    "log"
    "io/ioutil"
    "go-postgres/handlers/vehicle"
    _ "github.com/lib/pq"
)


func CreateVehicleProto(r *http.Request, w http.ResponseWriter) ([]byte) {
    fmt.Println("The content is Protobuf")
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)
    request := &vehicle.VehicleRequest{}
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
	log.Fatalf("Unable to read message from request : %v", err)
    }
    proto.Unmarshal(data, request)
    vin := request.GetVIN()
    mak := request.GetMake()
    model := request.GetModel()
    color := request.GetColor()
    vType := request.GetType()
    condition := request.GetCondition() 

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning id will return the id of the inserted vehicle
    sqlStatement := `INSERT INTO vehicles (vin, make, model, color, type, condition) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

    // the inserted id will store in this id
    var id int64

    // execute the sql statement
    err = db.QueryRow(sqlStatement, vin, mak, model, color, vType, condition).Scan(&id)

    response := &vehicle.VehicleErrorResponse{}
    if err != nil {
        fmt.Println("Unable to execute the query")
	response.ID = -1
	response.Message = ""
    } else {
	fmt.Printf("Inserted a single record %v", id)
	response.ID = id
	response.Message = "Vehicle created successfully"
    }

    resp, err := proto.Marshal(response)

    return resp

}
