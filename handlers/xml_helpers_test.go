package handlers

import (
    "bytes"
    "testing"
    "net/http"
    "net/http/httptest"
    "go-postgres/models"
    "encoding/xml"
)

func TestFetchXmlVehicle(t *testing.T) {
    var id int64 = 1000

    req := httptest.NewRequest(http.MethodGet, "/api/vehicle/1000", nil)
    w := httptest.NewRecorder()

    res := GetVehicleXml(req, w, id)

    if res.Error == "" {
	t.Log("Error is nil")
	t.Error("GetVehicleXml(1000) expects error")
    }

}


func TestFetchValidXmlVehicle(t *testing.T) {
    var id int64 = 1

    req := httptest.NewRequest(http.MethodGet, "/api/vehicle/1", nil)
    w := httptest.NewRecorder()

    res := GetVehicleXml(req, w, id)

    if res.Error != ""  {
        t.Log("Error should be nil")
	t.Log(res.Error)
        t.Error("GetVehicleXml(1) do not expect error")
    }

}

func TestCreateXmlVehicle(t *testing.T) {

    payload := models.XMLVehicle{
	VIN: "XVDD",
	Make: "Mercedes",
	Model: "S",
	Color: "Black",
	Type: "car",
	Condition: "used",
    }
    xmlBytes, _ := xml.Marshal(&payload)


    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(xmlBytes))
    w := httptest.NewRecorder()

    res := CreateVehicleXml(req, w)

    if res.ID == -1  {
        t.Log("Error should be nil")
        t.Log(res.Message)
        t.Error("CreateVehicleXml() should create the respource successfuly")
    }

}


func TestUpdateXmlVehicle(t *testing.T) {

    payload := models.XMLVehicle{
        VIN: "NXZ8",
        Make: "audi",
        Model: "q5",
        Color: "Black",
        Type: "car",
        Condition: "used",
    }
    xmlBytes, _ := xml.Marshal(&payload)

    var id int64 = 1

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/1", bytes.NewReader(xmlBytes))
    w := httptest.NewRecorder()

    res := UpdateVehicleXml(req, w, id)

    if res.ID == -1  {
        t.Log(res.Message)
        t.Error("CreateVehicleXml() should create the respource successfuly")
    }

}


func TestUpdateInvalidXmlVehicle(t *testing.T) {

    payload := models.XMLVehicle{
        VIN: "NXZ8",
        Make: "audi",
        Model: "q5",
        Color: "Black",
        Type: "car",
        Condition: "used",
    }
    xmlBytes, _ := xml.Marshal(&payload)

    var id int64 = 1000

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/1000", bytes.NewReader(xmlBytes))
    w := httptest.NewRecorder()

    res := UpdateVehicleXml(req, w, id)

    if res.ID != -1  {
        t.Log(res.Message)
        t.Error("Expected ID=-1 as given ID is invalid")
    }

}

func TestDeleteXmlVehicle(t *testing.T) {

    var id int64 = 5

    req := httptest.NewRequest(http.MethodDelete, "/api/deletevehicle/5", nil)
    w := httptest.NewRecorder()

    res := DeleteVehicleXml(req, w, id)

    if res.ID == -1  {
        t.Log(res.Message)
        t.Error("Resource should be deleted successfully")
    }

}


func TestDeleteInvalidXmlVehicle(t *testing.T) {

    var id int64 = 2000

    req := httptest.NewRequest(http.MethodDelete, "/api/deletevehicle/2000", nil)
    w := httptest.NewRecorder()

    res := DeleteVehicleXml(req, w, id)

    if res.ID != -1  {
        t.Log(res.Message)
        t.Error("Invalid Resource deletion should not be successful")
    }

}

func TestGetAllXmlVehicle(t *testing.T) {


    _, err := GetAllVehicleXml()

    if err != nil  {
        t.Error("All resources should be fetched")
    }

}

func TestInsertXmlVehicle(t *testing.T) {
    payload := models.XMLVehicle{
        VIN: "NSS8",
        Make: "Suzuki",
        Model: "HH",
        Color: "green",
        Type: "motorcycle",
        Condition: "new",
    }

    _, err := InsertXVehicle(payload)

    if err != nil {
	    t.Error("Insert should not fail")
    }
}


func TestSearchXmlVehicle(t *testing.T) {


    req := httptest.NewRequest(http.MethodGet, "/api/search/vehicle?color=black", nil)
    w := httptest.NewRecorder()

    res := SearchVehicleXml(req, w)

    if res.Error != ""  {
        t.Log(res.Error)
        t.Error("Vehicle info should be fetched successfully")
    }

}

func TestSearchInvalidXmlVehicle(t *testing.T) {


    req := httptest.NewRequest(http.MethodGet, "/api/search/vehicle?color=maroon", nil)
    w := httptest.NewRecorder()

    res := SearchVehicleXml(req, w)

    if res.Error == ""  {
        t.Log(res.Error)
        t.Error("Error should be returned for invalid search criteria")
    }

}

func TestFetchVehicleXml(t *testing.T) {
    var id int64 = 1

    _, err := FetchXVehicle(id)

    if err != nil  {
        t.Log("Error should be nil")
        t.Error("FetchVehicle() do not expect failure")
    }

}

func TestFetchInvalidVehicleXml(t *testing.T) {
    var id int64 =  1000

    _, err := FetchXVehicle(id)

    if err == nil  {
        t.Log("Error should not be nil")
        t.Error("FetchVehicle() expect failure")
    }

}
