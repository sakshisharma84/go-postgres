package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "net/url"
    "encoding/json"
    "go-postgres/models"
    "log"
    _ "github.com/lib/pq"
)

type JsonResponse struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

type JsonResponseVehicle struct {
    Error   string `json:"error,omitempty"`
    VehicleDetails models.JSONVehicle `json:"details,omitempty"`
}

type JsonResponseList struct {
    Error   string `json:"error,omitempty"`
    VehicleDetails []models.JSONVehicle `json:"details,omitempty"`
}


func CreateVehicleJson(r *http.Request, w http.ResponseWriter) JsonResponse{
    fmt.Println("The content is JSON")
    var vehicle models.JSONVehicle
    var res JsonResponse

    // decode the json request to vehicle
    err := json.NewDecoder(r.Body).Decode(&vehicle)

    if err != nil {
        log.Printf("Unable to decode the request body.  %v", err)
    }
    insertID, err := InsertVehicle(vehicle)

    if err != nil {
        w.WriteHeader(400)
        res = JsonResponse{
            ID:      -1,
            Message: err.Error(),
        }
    } else {
         // format a response object
        res = JsonResponse{
            ID:      insertID,
            Message: "Vehicle created successfully",
        }
    }
    return res
}

func GetVehicleJson(r *http.Request, w http.ResponseWriter, id int64) JsonResponseVehicle {

    var res JsonResponseVehicle
    // call the getVehicle function using id to retrieve a vehicle
    vehicle, err := FetchVehicle(int64(id))

    if err != nil {
        w.WriteHeader(500)
        res = JsonResponseVehicle{
                Error: err.Error(),
                VehicleDetails: models.JSONVehicle{},
        }

    } else {
         res = JsonResponseVehicle{
                Error: "",
                VehicleDetails: vehicle,
        }
    }
    return res
}

func UpdateVehicleJson(r *http.Request, w http.ResponseWriter, id int64) JsonResponse {
    var vehicle models.JSONVehicle
    var res JsonResponse

    // decode the json request to vehivle
    err := json.NewDecoder(r.Body).Decode(&vehicle)

    if err != nil {
        log.Printf("Unable to decode the request body.  %v", err)
    }

    // call updateVehicle to update the vehicle
    updatedRows, err := ModifyVehicle(id, vehicle)

    if err != nil {
        w.WriteHeader(500)
        res = JsonResponse{
            ID:      -1,
            Message: err.Error(),
        }
    } else {
        // format the message string
        msg := fmt.Sprintf("Vehicle updated successfully. Total rows/record affected %v", updatedRows)

        // format the response message
        res = JsonResponse{
            ID:      id,
            Message: msg,
        }
    }

    return res
}

func DeleteVehicleJson(r *http.Request, w http.ResponseWriter, id int64) JsonResponse {

    var res JsonResponse
    deletedRows, err := RemoveVehicle(id)

    if err != nil {
        w.WriteHeader(500)
        res = JsonResponse{
            ID:      -1,
            Message: err.Error(),
        }
    } else {
        // format the message string
        msg := fmt.Sprintf("Vehicle deleted successfully. Total rows/record affected %v", deletedRows,)

        // format the response message
        res = JsonResponse{
            ID:      int64(id),
            Message: msg,
        }
    }

    return res
}

func GetAllVehicleJson() ([]models.JSONVehicle, error) {
    // get all the vehicles in the db
    vehicles, err := FetchAllVehicle()
    return vehicles, err

}
// insert vehicle in the DB
func InsertVehicle(vehicle models.JSONVehicle) (int64, error) {

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
    err := db.QueryRow(sqlStatement, vehicle.VIN, vehicle.Make, vehicle.Model, vehicle.Color, vehicle.Type, vehicle.Condition).Scan(&id)


    if err != nil {
        fmt.Println("Unable to execute the query")
        return id, err
    }

    fmt.Printf("Inserted a single record %v", id)

    // return the inserted id
    return id, nil
}

func SearchVehicleJson(r *http.Request, w http.ResponseWriter) JsonResponseList {
    var res JsonResponseList
    fmt.Println(r.URL.RawQuery)

    m, _ := url.ParseQuery(r.URL.RawQuery)

    vehicles, err := GetSearchedVehicle(m)

    if err != nil {
        w.WriteHeader(500)
        res = JsonResponseList {
                Error: err.Error(),
                VehicleDetails: vehicles,
        }

    } else {
         res = JsonResponseList {
                Error: "",
                VehicleDetails: vehicles,
        }
    }

    return res

}


func FetchVehicle(id int64) (models.JSONVehicle, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var vehicle models.JSONVehicle

    // create the select sql query
    sqlStatement := `SELECT * FROM vehicles WHERE id=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    // unmarshal the row object to vehivle
    err := row.Scan(&vehicle.ID, &vehicle.VIN, &vehicle.Make, &vehicle.Model, &vehicle.Color, &vehicle.Type, &vehicle.Condition)

    switch err {
    case sql.ErrNoRows:
        err := fmt.Errorf("No rows were returned!")
        return vehicle, err
    case nil:
        return vehicle, nil
    default:
        fmt.Println("Unable to scan the row")
        return vehicle, err
    }

    return vehicle, err
}



// update vehicle in the DB
func ModifyVehicle(id int64, vehicle models.JSONVehicle) (int64, error) {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the update sql query
    sqlStatement := `UPDATE vehicles SET vin=$2, make=$3, model=$4, color=$5, type=$6, condition=$7 WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id, vehicle.VIN, vehicle.Make, vehicle.Model, vehicle.Color, vehicle.Type, vehicle.Condition)

    if err != nil {
        log.Printf("Unable to execute the query. %v", err)
        return -1, err
    }

    // check how many rows affected
    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Printf("Error while checking the affected rows. %v", err)
        return rowsAffected, err
    }

    if rowsAffected == 0 {
        return 0, fmt.Errorf("No rows existing with this id")
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected, nil
}



// delete vehicle in the DB
func RemoveVehicle(id int64) (int64, error) {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the delete sql query
    sqlStatement := `DELETE FROM vehicles WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id)

    if err != nil {
        log.Printf("Unable to execute the query. %v", err)
        return -1, err
    }

    // check how many rows affected
    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Printf("Error while checking the affected rows. %v", err)
        return -1, nil
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)
    if rowsAffected == 0 {
        return -1, fmt.Errorf("No rows exist with this id")
    }

    return rowsAffected, nil
}


func FetchAllVehicle() ([]models.JSONVehicle, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var vehicles []models.JSONVehicle

    // create the select sql query
    sqlStatement := `SELECT * FROM vehicles`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Printf("Unable to execute the query. %v", err)
        return nil, err
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var vehicle models.JSONVehicle

        // unmarshal the row object to vehicle
        err = rows.Scan(&vehicle.ID, &vehicle.VIN, &vehicle.Make, &vehicle.Model, &vehicle.Color, &vehicle.Type, &vehicle.Condition)

        if err != nil {
            log.Printf("Unable to scan the row. %v", err)
            return nil, err
        }
        vehicles = append(vehicles, vehicle)
    }

    return vehicles, nil
}

func GetSearchedVehicle(m map[string][]string) ([]models.JSONVehicle, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var vehicles []models.JSONVehicle


    statement := ""
    for key, val := range m {
	    statement = key + "= '" + val[0] + "'"
    }

    // create the select sql query
    sqlStatement :=  `SELECT * FROM vehicles WHERE ` + statement 

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Printf("Unable to execute the query. %v", err)
        return nil, err
    }


    // close the statement
    defer rows.Close()
    count := 0

    // iterate over the rows
    for rows.Next() {
        var vehicle models.JSONVehicle

        // unmarshal the row object to vehicle
        err = rows.Scan(&vehicle.ID, &vehicle.VIN, &vehicle.Make, &vehicle.Model, &vehicle.Color, &vehicle.Type, &vehicle.Condition)

        if err != nil {
            log.Printf("Unable to scan the row. %v", err)
            return nil, err
        }
        vehicles = append(vehicles, vehicle)
	count = count + 1
    }

    if count == 0 {
	err := fmt.Errorf("Vehicle for the given search criteria not found!")
        return nil, err
    }

    return vehicles, nil
}
