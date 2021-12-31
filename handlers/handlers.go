package handlers

import (
    "database/sql"
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "encoding/xml"
    "fmt"
    "time"
    "log"
    "net/http" // used to access the request and response object of the api
    "os"       // used to read the environment variable
    "strconv"  // package used to covert string into int type

    "github.com/gorilla/mux" // used to get the params from the route

    "github.com/joho/godotenv" // package used to read the .env file
    _ "github.com/lib/pq"      // postgres golang driver
)

var (
	cacheSince = time.Now().Format(http.TimeFormat)
	cacheUntil = time.Now().AddDate(60, 0, 0).Format(http.TimeFormat)
)


// create connection with postgres db
func createConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
        log.Println("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    // return the connection
    return db
}


// CreateVehicle create a vehicle in the postgres db
func CreateVehicle(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


    contentType := r.Header.Get("Content-type")

    switch contentType {
	case "application/json" :
		res := CreateVehicleJson(r, w)
		json.NewEncoder(w).Encode(res)
		break

	case "application/xml" :
		fmt.Println("The content is XML")
		res := CreateVehicleXml(r, w)
                xml.NewEncoder(w).Encode(res)
		break

	case "protobuf" :
		fmt.Println("The content is Protobuf")
		resp := CreateVehicleProto(r, w)
		//fmt.Fprintf(w, string(resp))
		w.Write(resp)
                break

	default :
		log.Printf("Unsupported Content Type: %s", contentType)
    }

}


// GetVehicle will return a vehicle by its id
func GetVehicle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // Adding cache support for GET requests
    w.Header().Set("Cache-Control", "max-age:3600, public")
    w.Header().Set("Last-Modified", cacheSince)
    w.Header().Set("Expires", cacheUntil)

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
       log.Printf("Unable to convert the string into int.  %v", err)

    }

    contentType := r.Header.Get("Content-type")

    switch contentType {
        case "application/json" :
                res := GetVehicleJson(r, w, int64(id))
                json.NewEncoder(w).Encode(res)
                break

        case "application/xml" :
                fmt.Println("The content is XML")
		res := GetVehicleXml(r, w, int64(id))
                xml.NewEncoder(w).Encode(res)
                break

        default :
                log.Printf("Unsupported Content Type: %s", contentType)
    }

}

// UpdateVehicle updates vehicle's params in the postgres db
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    params := mux.Vars(r)

    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Printf("Unable to convert the string into int.  %v", err)
    }

    contentType := r.Header.Get("Content-type")

    switch contentType {
        case "application/json" :
                res := UpdateVehicleJson(r, w, int64(id))
                json.NewEncoder(w).Encode(res)
                break

        case "application/xml" :
                fmt.Println("The content is XML")
		res := UpdateVehicleXml(r, w, int64(id))
                xml.NewEncoder(w).Encode(res)
                break
        default :
                log.Printf("Unsupported Content Type: %s", contentType)
    }

}

// DeleteVehicle delete vehicle's detail in the postgres db
func DeleteVehicle(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // get the id from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id in string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Printf("Unable to convert the string into int.  %v", err)
    }
    contentType := r.Header.Get("Content-type")

    switch contentType {
        case "application/json" :
                res := DeleteVehicleJson(r, w, int64(id))
                json.NewEncoder(w).Encode(res)
                break

        case "application/xml" :
                fmt.Println("The content is XML")
		res := DeleteVehicleXml(r, w, int64(id))
                xml.NewEncoder(w).Encode(res)
                break

        default :
                log.Printf("Unsupported Content Type: %s", contentType)
    }

}

// GetAllVehicles will return all the vehicles
func GetAllVehicle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Cache-Control", "max-age:3600, public")
    w.Header().Set("Last-Modified", cacheSince)
    w.Header().Set("Expires", cacheUntil)

    contentType := r.Header.Get("Content-type")

    switch contentType {
        case "application/json" :
                vehicles, err := GetAllVehicleJson()
		if err != nil {
			log.Printf("Unable to get all vehicles. %v", err)
			w.WriteHeader(500)
		}
                // send the response
                json.NewEncoder(w).Encode(vehicles)
                break

        case "application/xml" :
                fmt.Println("The content is XML")
		vehicles, err := GetAllVehicleXml()
                if err != nil {
                        log.Printf("Unable to get all vehicles. %v", err)
                        w.WriteHeader(500)
                }
                // send the response
                xml.NewEncoder(w).Encode(vehicles)
                break

        default :
                log.Printf("Unsupported Content Type: %s", contentType)
    }

}


// SearchVehicle will return all the vehicles corresponding to the searching criteria
func SearchVehicle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Cache-Control", "max-age:3600, public")
    w.Header().Set("Last-Modified", cacheSince)
    w.Header().Set("Expires", cacheUntil)

    contentType := r.Header.Get("Content-type")

    switch contentType {
        case "application/json" :
                res := SearchVehicleJson(r, w)
                json.NewEncoder(w).Encode(res)
                break

        case "application/xml" :
                fmt.Println("The content is XML")
                res:= SearchVehicleXml(r, w)
                // send the response
                xml.NewEncoder(w).Encode(res)
                break

        default :
                log.Printf("Unsupported Content Type: %s", contentType)
    }

}
