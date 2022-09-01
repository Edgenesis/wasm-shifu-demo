package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type DeviceInfo struct {
	StatusCode   string   `json:"statusCode"`
	Message      string   `json:"message"`
	Entity       []Entity `json:"entity"`
	DeviceId     int      `json:"deviceId"`
	DeviceName   string   `json:"deviceName"`
	DeviceRemark string   `json:"deviceRemark"`
}

type Entity struct {
	DateTime string `json:"dateTime"`
	EUnit    string `json:"eUnit"`
	Evalue   string `json:"eValue"`
	Ekey     string `json:"eKey"`
	Ename    string `json:"eName"`
	Enum     string `json:"eNum"`
}

var (
	eunits   = []string{"℃", "%RH"}
	ename    = []string{"大气温度", "大气湿度"}
	baseData = []float64{30, 80}
	port     = "0.0.0.0:8099"
)

func main() {
	http.HandleFunc("/get_info", getInfo)
	log.Println("listening ", port)
	err := http.ListenAndServe(port, nil)
	log.Println(err)
}

func getInfo(w http.ResponseWriter, req *http.Request) {

	info := generateData()
	data, err := json.Marshal(info)
	if err != nil {
		log.Println(err)
		io.WriteString(w, "error when generate Json error:"+err.Error())
		return
	}
	io.WriteString(w, string(data))
}

func generateData() *DeviceInfo {
	deviceId := rand.Intn(1 << 20)
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	return &DeviceInfo{
		StatusCode:   "200",
		Message:      "success",
		DeviceId:     deviceId,
		DeviceName:   fmt.Sprintf("%d", deviceId),
		DeviceRemark: timeNow,
		Entity: []Entity{
			{
				DateTime: timeNow,
				EUnit:    eunits[0],
				Evalue:   fmt.Sprintf("%.2f", rand.Float64()+baseData[0]+float64(rand.Intn(5))),
				Ekey:     "e1",
				Ename:    ename[0],
				Enum:     "101",
			},
			{
				DateTime: timeNow,
				EUnit:    eunits[1],
				Evalue:   fmt.Sprintf("%.2f", rand.Float64()+baseData[1]+float64(rand.Intn(5))),
				Ekey:     "e2",
				Ename:    ename[1],
				Enum:     "102",
			},
		},
	}
}
