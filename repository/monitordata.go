package repository

import (
	"fmt"
	"strconv"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shopspring/decimal"

	"snns_srv/db"
)

//MonitorData :
type MonitorData struct {
	// Uid: Primary key (_id)
	Host string  `bson:"host"`
	CPU  float64 `bson:"cpu"`
	Mem  float64 `bson:"mem"`
}

//GetMonitorData :
func GetMonitorData() ([9]MonitorData, error) {
	var mds [9]MonitorData
	// var md MonitorData
	var res []client.Result
	var qs string
	var err error
	var row []interface{}
	var f, v float64
	// md = MonitorData{}
	for i := 0; i < 9; i++ {
		mds[i].Host = "SNN0" + strconv.Itoa(i+1)

		//get Cpu
		qs = "select usage_system+usage_user  from cpu  where cpu='cpu-total' where host='" + mds[i].Host + "' order by time desc limit 1"
		res, err = db.QueryInfluxDB(db.Conntsdb, "telegraf", qs)
		if err != nil {
			fmt.Printf("err !")
		}
		row = res[0].Series[0].Values[0]
		f, _ = strconv.ParseFloat(fmt.Sprint(row[1]), 64)
		v, _ = decimal.NewFromFloat(f).Round(1).Float64()
		mds[i].CPU = v

		//get Mem
		qs = "select 100-available_percent  from mem order by time desc  limit 1"
		res, err = db.QueryInfluxDB(db.Conntsdb, "telegraf", qs)
		if err != nil {
			fmt.Printf("err !")
		}
		row = res[0].Series[0].Values[0]
		f, _ = strconv.ParseFloat(fmt.Sprint(row[1]), 64)
		v, _ = decimal.NewFromFloat(f).Round(1).Float64()
		mds[i].Mem = v
	}

	//var v float64

	return mds, err
}
