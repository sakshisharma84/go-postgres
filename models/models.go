package models

// schema of the cars table
type JSONVehicle struct {
    ID		int64  `json:"id"`
    VIN		string  `json:"vin"`
    Make	string `json:"make"`
    Model	string `json:"model"`
    Color	string `json:"color"`
    Type	string `json:"type"`
    Condition	string `json:"condition"`
}

type XMLVehicle struct {
    ID          int64  `xml:"id"`
    VIN         string `xml:"vin"`
    Make        string `xml:"make"`
    Model       string `xml:"model"`
    Color       string `xml:"color"`
    Type        string `xml:"type"`
    Condition   string `xml:"condition"`

}

