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
