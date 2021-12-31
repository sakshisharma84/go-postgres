package handlers

import (
    "bytes"
    "testing"
    "net/http"
    "net/http/httptest"
    "go-postgres/handlers/vehicle"
    proto "github.com/golang/protobuf/proto"
)

func TestCreateVehicleProto(t *testing.T) {

    payload := &vehicle.VehicleRequest{
	    VIN: "BCJD",
	    Make: "GMC",
	    Model: "G2",
	    Color: "black",
	    Type: "car",
	    Condition: "used",
    }

    data, err := proto.Marshal(payload)
    if err != nil {
	t.Log("Error while marshalling the object")
    }

    req := httptest.NewRequest(http.MethodPost, "/api/newvehicle", bytes.NewReader(data))
    w := httptest.NewRecorder()

    resp := CreateVehicleProto(req, w)

    if  resp == nil  {
        t.Log(resp)
        t.Error("CreateVehicle() should create the resource successfuly")
    }
}

func TestUpdateVehicleProto(t *testing.T) {

    payload := &vehicle.VehicleRequest{
            VIN: "NXZ8",
            Make: "Audi",
            Model: "q5",
            Color: "black",
            Type: "car",
            Condition: "used",
    }

    data, err := proto.Marshal(payload)
    if err != nil {
        t.Log("Error while marshalling the object")
    }

    var id int64 = 1

    req := httptest.NewRequest(http.MethodPut, "/api/vehicle/1", bytes.NewReader(data))
    w := httptest.NewRecorder()

    resp := UpdateVehicleProto(req, w, id)

    if  resp == nil  {
        t.Log(resp)
        t.Error("UpdateVehicle() should create the resource successfuly")
    }
}

func TestGetVehicleProto(t *testing.T) {

    var id int64 = 1

    req := httptest.NewRequest(http.MethodGet, "/api/vehicle/1", nil)
    w := httptest.NewRecorder()

    resp := GetVehicleProto(req, w, id)

    if  resp == nil  {
        t.Log(resp)
        t.Error("GetVehicle() should create the resource successfuly")
    }
}

func TestDeleteVehicleProto(t *testing.T) {

    var id int64 = 2

    req := httptest.NewRequest(http.MethodDelete, "/api/deletevehicle/2", nil)
    w := httptest.NewRecorder()

    resp := DeleteVehicleProto(req, w, id)

    if  resp == nil  {
        t.Log(resp)
        t.Error("DeleteVehicle() should create the resource successfuly")
    }
}

func TestSearchVehicleProto(t *testing.T) {


    req := httptest.NewRequest(http.MethodGet, "/api/search/vehicle?color=black", nil)
    w := httptest.NewRecorder()

    resp := SearchVehicleProto(req, w)

    if resp == nil  {
        t.Log(resp)
        t.Error("Vehicle info should be fetched successfully")
    }

}


