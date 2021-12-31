package handlers

import (
    "bytes"
    "testing"
    "net/http"
    "net/http/httptest"
    "go-postgres/models"
    "encoding/json"
)

func TestFetchVehicle(t *testing.T) {
    var id int64 = 1000

    req := httptest.NewRequest(http.MethodGet, "/api/vehicle/1000", nil)
    w := httptest.NewRecorder()

    res := GetVehicleJson(req, w, id)

    if res.Error == "" {
	t.Log("Error is nil")
	t.Error("GetVehicleJson(1000) expects error")
    }

}


func TestFetchValidVehicle(t *testing.T) {
    var id int64 = 1

    req := httptest.NewRequest(http.MethodGet, "/api/vehicle/1", nil)
    w := httptest.NewRecorder()

    res := GetVehicleJson(req, w, id)

    if res.Error != ""  {
        t.Log("Error should be nil")
	t.Log(res.Error)
        t.Error("GetVehicleJson(1) do not expect error")
    }

}

func TestCreateVehicleJson(t *testing.T) {

    payload := &models.JSONVehicle{
	VIN: "XVDD",
	Make: "Mercedes",
	Model: "S",
	Color: "Black",
	Type: "car",
	Condition: "used",
    }
    jsonBytes, _ := json.Marshal(payload)

    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(jsonBytes))
    w := httptest.NewRecorder()

    res := CreateVehicleJson(req, w)

    if res.ID == -1  {
        t.Log("Error should be nil")
        t.Log(res.Message)
        t.Error("CreateVehicleJson() should create the respource successfuly")
    }

}


func TestUpdateVehicleiJson(t *testing.T) {

    payload := &models.JSONVehicle{
        VIN: "NXZ8",
        Make: "audi",
        Model: "q5",
        Color: "Black",
        Type: "car",
        Condition: "used",
    }
    jsonBytes, _ := json.Marshal(payload)

    var id int64 = 1

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/1", bytes.NewReader(jsonBytes))
    w := httptest.NewRecorder()

    res := UpdateVehicleJson(req, w, id)

    if res.ID == -1  {
        t.Log(res.Message)
        t.Error("CreateVehicleJson() should create the respource successfuly")
    }

}


func TestUpdateInvalidVehicleJson(t *testing.T) {

    payload := &models.JSONVehicle{
        VIN: "NXZ8",
        Make: "audi",
        Model: "q5",
        Color: "Black",
        Type: "car",
        Condition: "used",
    }
    jsonBytes, _ := json.Marshal(payload)

    var id int64 = 1000

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/1000", bytes.NewReader(jsonBytes))
    w := httptest.NewRecorder()

    res := UpdateVehicleJson(req, w, id)

    if res.ID != -1  {
        t.Log(res.Message)
        t.Error("Expected ID=-1 as given ID is invalid")
    }

}

func TestDeleteVehicleJson(t *testing.T) {

    var id int64 = 3

    req := httptest.NewRequest(http.MethodDelete, "/api/deletevehicle/2", nil)
    w := httptest.NewRecorder()

    res := DeleteVehicleJson(req, w, id)

    if res.ID == -1  {
        t.Log(res.Message)
        t.Error("Resource should be deleted successfully")
    }

}


func TestDeleteInvalidVehicleJson(t *testing.T) {

    var id int64 = 2000

    req := httptest.NewRequest(http.MethodDelete, "/api/deletevehicle/2000", nil)
    w := httptest.NewRecorder()

    res := DeleteVehicleJson(req, w, id)

    if res.ID != -1  {
        t.Log(res.Message)
        t.Error("Invalid Resource deletion should not be successful")
    }

}

func TestGetAllVehicleJson(t *testing.T) {


    _, err := GetAllVehicleJson()

    if err != nil  {
        t.Error("All resources should be fetched")
    }

}

func TestInsertVehicleJson(t *testing.T) {
    payload := models.JSONVehicle{
        VIN: "NSS8",
        Make: "Suzuki",
        Model: "HH",
        Color: "green",
        Type: "motorcycle",
        Condition: "new",
    }

    _, err := InsertVehicle(payload)

    if err != nil {
	    t.Error("Insert should not fail")
    }
}


func TestSearchVehicleJson(t *testing.T) {


    req := httptest.NewRequest(http.MethodGet, "/api/search/vehicle?color=black", nil)
    w := httptest.NewRecorder()

    res := SearchVehicleJson(req, w)

    if res.Error != ""  {
        t.Log(res.Error)
        t.Error("Vehicle info should be fetched successfully")
    }

}

func TestSearchInvalidVehicleJson(t *testing.T) {


    req := httptest.NewRequest(http.MethodGet, "/api/search/vehicle?color=maroon", nil)
    w := httptest.NewRecorder()

    res := SearchVehicleJson(req, w)

    if res.Error == ""  {
        t.Log(res.Error)
        t.Error("Error should be returned for invalid search criteria")
    }

}

func TestFetchVehicleJson(t *testing.T) {
    var id int64 = 1

    _, err := FetchVehicle(id)

    if err != nil  {
        t.Log("Error should be nil")
        t.Error("FetchVehicle() do not expect failure")
    }

}

func TestFetchInvalidVehicleJson(t *testing.T) {
    var id int64 =  1000

    _, err := FetchVehicle(id)

    if err == nil  {
        t.Log("Error should not be nil")
        t.Error("FetchVehicle() expect failure")
    }

}
