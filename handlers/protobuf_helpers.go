package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    proto "github.com/golang/protobuf/proto"
    "log"
    "net/url"
    "io/ioutil"
    "go-postgres/handlers/vehicle"
    _ "github.com/lib/pq"
)


func CreateVehicleProto(r *http.Request, w http.ResponseWriter) ([]byte) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)
    request := &vehicle.VehicleRequest{}
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
	log.Printf("Unable to read message from request : %v", err)
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


func GetVehicleProto(r *http.Request, w http.ResponseWriter, id int64) []byte {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the select sql query
    sqlStatement := `SELECT * FROM vehicles WHERE id=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    vehInfo := &vehicle.VehicleRequest{}

    // unmarshal the row object to vehivle
    err := row.Scan(&vehInfo.ID, &vehInfo.VIN, &vehInfo.Make, &vehInfo.Model, &vehInfo.Color, &vehInfo.Type, &vehInfo.Condition)

    response := &vehicle.VehicleErrorResponse{}

    switch err {
    case sql.ErrNoRows:
        response.ID = -1
	response.Message = "No rows were returned"
    case nil:
        response.ID = id
	mess := fmt.Sprintf("%v",vehInfo)
	response.Message = mess 
    default:
        response.ID = id
	response.Message = "Unable to scan the row"
    }

    resp, err := proto.Marshal(response)

    return resp

}

func UpdateVehicleProto(r *http.Request, w http.ResponseWriter, id int64) ([]byte) {
    contentLength := r.ContentLength
    fmt.Printf("Content Length Received : %v\n", contentLength)
    request := &vehicle.VehicleRequest{}
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Printf("Unable to read message from request : %v", err)
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

    // create the update sql query
    sqlStatement := `UPDATE vehicles SET vin=$2, make=$3, model=$4, color=$5, type=$6, condition=$7 WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id, vin, mak, model, color, vType, condition)

    response := &vehicle.VehicleErrorResponse{}

    if err != nil {
        log.Printf("Unable to execute the query. %v", err)
	response.ID = -1
	response.Message = "Unable to execute the query"
    } else {

	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
	    response.ID = -1
	    response.Message = "No rows existing with this id"
	} else {
            log.Printf("Total rows/record affected %v", rowsAffected)
            log.Printf("Inserted a single record %v", id)
            response.ID = id
            response.Message = "Vehicle updated successfully"
	}
    }

    resp, err := proto.Marshal(response)

    return resp
}


func DeleteVehicleProto(r *http.Request, w http.ResponseWriter, id int64) []byte {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the select sql query
    sqlStatement := `DELETE FROM vehicles WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id)

    response := &vehicle.VehicleErrorResponse{}

    if err != nil {
        log.Printf("Unable to execute the query. %v", err)
        response.ID = -1
        response.Message = "Unable to execute the query"
    } else {

        // check how many rows affected
        rowsAffected, _ := res.RowsAffected()

        if rowsAffected == 0 {
            response.ID = -1
            response.Message = "No rows existing with this id"
        } else {
            log.Printf("Total rows/record affected %v", rowsAffected)
            log.Printf("Deleted a single record %v", id)
            response.ID = id
            response.Message = "Vehicle deleted successfully"
        }
    }

    resp, err := proto.Marshal(response)

    return resp

}

func SearchVehicleProto(r *http.Request, w http.ResponseWriter) []byte {
    fmt.Println(r.URL.RawQuery)
    m, _ := url.ParseQuery(r.URL.RawQuery)

    db := createConnection()

    // close the db connection
    defer db.Close()

    statement := ""
    for key, val := range m {
            statement = key + "= '" + val[0] + "'"
    }

    // create the select sql query
    sqlStatement :=  `SELECT * FROM vehicles WHERE ` + statement 

    // execute the sql statement
    rows, _ := db.Query(sqlStatement)

    // close the statement
    defer rows.Close()
    count := 0
    var vehicles []vehicle.VehicleRequest
    var id int64

    response := &vehicle.VehicleErrorResponse{}
    // iterate over the rows
    for rows.Next() {
	vehInfo := vehicle.VehicleRequest{}

        // unmarshal the row object to vehivle
        err := rows.Scan(&vehInfo.ID, &vehInfo.VIN, &vehInfo.Make, &vehInfo.Model, &vehInfo.Color, &vehInfo.Type, &vehInfo.Condition)

        if err != nil {
            log.Printf("Unable to scan the row. %v", err)
            response.ID = -1
	    response.Message = "Unable to scan the row"
        }
	id = vehInfo.ID
        vehicles = append(vehicles, vehInfo)
        count = count + 1
    }

    if count == 0 {
        log.Printf("Vehicle for the given search criteria not found!")
        response.ID = -1
	response.Message = "Vehicle for the given search criteria not found!"
    } else {
        response.ID = id 
        mess := fmt.Sprintf("%v",vehicles)
        response.Message = mess
    }

    resp, _ := proto.Marshal(response)

    return resp



}
