package handlers

import (
    "fmt"
    "bytes"
    "testing"
    "net/http"
    "net/http/httptest"
    "go-postgres/models"
    "encoding/json"
    "encoding/xml"
    "io/ioutil"
    "go-postgres/handlers/vehicle"
    "github.com/gorilla/mux"
    proto "github.com/golang/protobuf/proto"
)

func TestCreateVehicleJSONData(t *testing.T) {

     payload := &models.JSONVehicle{
        VIN: "XVVKKK",
        Make: "VolksVagon",
        Model: "Jetta",
        Color: "White",
        Type: "car",
        Condition: "used",
    }
    jsonBytes, _ := json.Marshal(payload)

    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(jsonBytes))
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")

    CreateVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
	t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusOK != result.StatusCode {
	t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}


func TestCreateVehicleXMLData(t *testing.T) {

    payload := models.XMLVehicle{
        VIN: "IIII",
        Make: "Nissan",
        Model: "N3",
        Color: "Brown",
        Type: "car",
        Condition: "used",
    }
    xmlBytes, _ := xml.Marshal(&payload)

    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(xmlBytes))
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/xml")

    CreateVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusOK != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}

func TestCreateVehicleProtoData(t *testing.T) {

    payload := &vehicle.VehicleRequest{
        VIN: "IVB",
        Make: "Chevrolet",
        Model: "C6",
        Color: "Yellow",
        Type: "car",
        Condition: "used",
    }

    data, err := proto.Marshal(payload)
    if err != nil {
        t.Log("Error while marshalling the object")
    }

    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(data))

    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "protobuf")

    CreateVehicle(w, req)

    result := w.Result()

    _, err = ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusOK != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

    // Check the response body is what we expect.
    /*expected := `{"alive": true}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }*/
}

func TestCreateVehicleUnknownData(t *testing.T) {

    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", nil)
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/unknown")

    CreateVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusUnsupportedMediaType != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}

func TestGetVehicleJSONData(t *testing.T) {

    id := "1"
    url := "/api/vehicle/" + id
    req := httptest.NewRequest(http.MethodGet, url, nil)
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")

    vars := map[string]string{
        "id": "1",
    }

    req = mux.SetURLVars(req, vars)

    GetVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusOK != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}

func TestGetInvalidVehicleJSONData(t *testing.T) {

    id := "1000"
    url := "/api/vehicle/" + id
    req := httptest.NewRequest(http.MethodGet, url, nil)
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")

    vars := map[string]string{
        "id": "1000",
    }

    req = mux.SetURLVars(req, vars)


    GetVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusInternalServerError != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}

func TestGetInvalidVehicleXMlData(t *testing.T) {

    id := "1000"
    url := "/api/vehicle/" + id
    req := httptest.NewRequest(http.MethodGet, url, nil)
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/xml")

    vars := map[string]string{
        "id": "1000",
    }

    req = mux.SetURLVars(req, vars)


    GetVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusInternalServerError != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}


func TestUpdateVehicleJSONData(t *testing.T) {

     payload := &models.JSONVehicle{
        VIN: "XVVKKK",
        Make: "VolksVagon",
        Model: "Jetta",
        Color: "White",
        Type: "car",
        Condition: "used",
    }
    jsonBytes, _ := json.Marshal(payload)

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/2", bytes.NewReader(jsonBytes))
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")

    vars := map[string]string{
        "id": "2",
    }

    req = mux.SetURLVars(req, vars)

    UpdateVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusOK != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}

func TestUpdateVehicleXMLData(t *testing.T) {

    payload := models.XMLVehicle{
        VIN: "IIII",
        Make: "Nissan",
        Model: "N3",
        Color: "Brown",
        Type: "car",
        Condition: "used",
    }
    xmlBytes, _ := xml.Marshal(&payload)

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/2", bytes.NewReader(xmlBytes))
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/xml")

    vars := map[string]string{
        "id": "2",
    }

    req = mux.SetURLVars(req, vars)
    UpdateVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    if http.StatusOK != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}


func TestGetAllVehicleJsonData(t *testing.T) {

    req := httptest.NewRequest(http.MethodGet, "/api/vehicle", nil)
    w := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/json")

    GetAllVehicle(w, req)

    result := w.Result()

    _, err := ioutil.ReadAll(result.Body)
    if err != nil {
        t.Error(err)
    }
    result.Body.Close()

    t.Log(result.StatusCode)
    // Check the response body is what we expect.
    fmt.Println(w.Body)

    if http.StatusOK != result.StatusCode {
        t.Error("wanted", http.StatusOK, "got", result.StatusCode)
    }

}


func BenchmarkCreateVehicle10(b *testing.B) {
    for n:=0;n <b.N; n++ {
	payload := &models.JSONVehicle{
            VIN: "XVVKKK",
            Make: "VolksVagon",
            Model: "Jetta",
            Color: "White",
            Type: "car",
            Condition: "used",
        }
        jsonBytes, _ := json.Marshal(payload)

        req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(jsonBytes))
        w := httptest.NewRecorder()
        req.Header.Set("Content-Type", "application/json")

        CreateVehicle(w, req)

        result := w.Result()

        _, err := ioutil.ReadAll(result.Body)
        if err != nil {
            b.Error(err)
        }
        result.Body.Close()

        b.Log(result.StatusCode)
        if http.StatusOK != result.StatusCode {
            b.Error("wanted", http.StatusOK, "got", result.StatusCode)
        }
    }
}
