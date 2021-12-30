package handlers

import (
    "database/sql"
    "fmt"
    "net/url"
    "net/http"
    "encoding/xml"
    "go-postgres/models"
    "log"
    _ "github.com/lib/pq"
)

type XmlResponse struct {
    ID      int64  `xml:"id,omitempty"`
    Message string `xml:"message,omitempty"`
}

type XmlResponseVehicle struct {
    Error   string `xml:"error,omitempty"`
    VehicleDetails models.XMLVehicle `xml:"details,omitempty"`
}

type XmlResponseList struct {
    Error   string `xml:"error,omitempty"`
    VehicleDetails []models.XMLVehicle `xml:"details,omitempty"`
}


func createVehicleXml(r *http.Request, w http.ResponseWriter) XmlResponse{
    fmt.Println("The content is XML")
    var vehicle models.XMLVehicle
    var res XmlResponse

    // decode the xml request to vehicle
    err := xml.NewDecoder(r.Body).Decode(&vehicle)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
    insertID, err := insertXVehicle(vehicle)

    if err != nil {
        w.WriteHeader(400)
        res = XmlResponse{
            ID:      -1,
            Message: err.Error(),
        }
    } else {
         // format a response object
        res = XmlResponse{
            ID:      insertID,
            Message: "Vehicle created successfully",
        }
    }
    return res
}

func getVehicleXml(r *http.Request, w http.ResponseWriter, id int64) XmlResponseVehicle {

    var res XmlResponseVehicle
    // call the getVehicle function using id to retrieve a vehicle
    vehicle, err := getXVehicle(int64(id))

    if err != nil {
        w.WriteHeader(500)
        res = XmlResponseVehicle{
                Error: err.Error(),
                VehicleDetails: models.XMLVehicle{},
        }

    } else {
         res = XmlResponseVehicle{
                Error: "",
                VehicleDetails: vehicle,
        }
    }
    return res
}

func updateVehicleXml(r *http.Request, w http.ResponseWriter, id int64) XmlResponse {
    var vehicle models.XMLVehicle
    var res XmlResponse

    // decode the xml request to vehivle
    err := xml.NewDecoder(r.Body).Decode(&vehicle)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call updateVehicle to update the vehicle
    updatedRows, err := updateXVehicle(id, vehicle)

    if err != nil {
        w.WriteHeader(500)
        res = XmlResponse{
            ID:      -1,
            Message: err.Error(),
        }
    } else {
        // format the message string
        msg := fmt.Sprintf("Vehicle updated successfully. Total rows/record affected %v", updatedRows)

        // format the response message
        res = XmlResponse{
            ID:      id,
            Message: msg,
        }
    }

    return res
}

func deleteVehicleXml(r *http.Request, w http.ResponseWriter, id int64) XmlResponse {

    var res XmlResponse
    deletedRows, err := deleteXVehicle(int64(id))

    if err != nil {
        w.WriteHeader(500)
        res = XmlResponse{
            ID:      -1,
            Message: err.Error(),
        }
    } else {
        // format the message string
        msg := fmt.Sprintf("Vehicle deleted successfully. Total rows/record affected %v", deletedRows,)

        // format the response message
        res = XmlResponse{
            ID:      int64(id),
            Message: msg,
        }
    }

    return res
}

func getAllVehicleXml() ([]models.XMLVehicle, error) {
    // get all the vehicles in the db
    vehicles, err := getAllXVehicle()
    return vehicles, err

}

func searchVehicleXml(r *http.Request, w http.ResponseWriter) XmlResponseList {
    var res XmlResponseList
    fmt.Println(r.URL.RawQuery)

    m, _ := url.ParseQuery(r.URL.RawQuery)

    vehicles, err := getSearchedXVehicle(m)

    if err != nil {
        w.WriteHeader(500)
        res = XmlResponseList {
                Error: err.Error(),
                VehicleDetails: vehicles,
        }

    } else {
         res = XmlResponseList {
                Error: "",
                VehicleDetails: vehicles,
        }
    }

    return res

}

// insert vehicle in the DB
func insertXVehicle(vehicle models.XMLVehicle) (int64, error) {

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



func getXVehicle(id int64) (models.XMLVehicle, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var vehicle models.XMLVehicle

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
func updateXVehicle(id int64, vehicle models.XMLVehicle) (int64, error) {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the update sql query
    sqlStatement := `UPDATE vehicles SET vin=$2, make=$3, model=$4, color=$5, type=$6, condition=$7 WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id, vehicle.VIN, vehicle.Make, vehicle.Model, vehicle.Color, vehicle.Type, vehicle.Condition)

    if err != nil {
        log.Println("Unable to execute the query. %v", err)
        return -1, err
    }

    // check how many rows affected
    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Println("Error while checking the affected rows. %v", err)
        return rowsAffected, err
    }

    if rowsAffected == 0 {
        return 0, fmt.Errorf("No rows existing with this id")
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)

    return rowsAffected, nil
}



// delete vehicle in the DB
func deleteXVehicle(id int64) (int64, error) {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the delete sql query
    sqlStatement := `DELETE FROM vehicles WHERE id=$1`

    // execute the sql statement
    res, err := db.Exec(sqlStatement, id)

    if err != nil {
        log.Println("Unable to execute the query. %v", err)
        return -1, err
    }

    // check how many rows affected
    rowsAffected, err := res.RowsAffected()

    if err != nil {
        log.Println("Error while checking the affected rows. %v", err)
        return -1, nil
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)
    if rowsAffected == 0 {
        return -1, fmt.Errorf("No rows exist with this id")
    }

    return rowsAffected, nil
}


func getAllXVehicle() ([]models.XMLVehicle, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var vehicles []models.XMLVehicle

    // create the select sql query
    sqlStatement := `SELECT * FROM vehicles`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Println("Unable to execute the query. %v", err)
        return nil, err
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var vehicle models.XMLVehicle

        // unmarshal the row object to vehicle
        err = rows.Scan(&vehicle.ID, &vehicle.VIN, &vehicle.Make, &vehicle.Model, &vehicle.Color, &vehicle.Type, &vehicle.Condition)

        if err != nil {
            log.Println("Unable to scan the row. %v", err)
            return nil, err
        }
        vehicles = append(vehicles, vehicle)
    }

    return vehicles, nil
}



func getSearchedXVehicle(m map[string][]string) ([]models.XMLVehicle, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var vehicles []models.XMLVehicle


    statement := ""
    for key, val := range m {
	    statement = key + "= '" + val[0] + "'"
    }

    // create the select sql query
    sqlStatement :=  `SELECT * FROM vehicles WHERE ` + statement 

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Println("Unable to execute the query. %v", err)
        return nil, err
    }


    // close the statement
    defer rows.Close()
    count := 0

    // iterate over the rows
    for rows.Next() {
        var vehicle models.XMLVehicle

        // unmarshal the row object to vehicle
        err = rows.Scan(&vehicle.ID, &vehicle.VIN, &vehicle.Make, &vehicle.Model, &vehicle.Color, &vehicle.Type, &vehicle.Condition)

        if err != nil {
            log.Println("Unable to scan the row. %v", err)
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
