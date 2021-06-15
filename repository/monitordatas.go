package repository

import (
	"fmt"
	"math"
	"strconv"

	client "github.com/influxdata/influxdb1-client/v2"

	"snns_srv/db"
)

// MonitorDatas
type MonitorDatas struct {
	Host string    `bson:"host"`
	CPUs []float64 `bson:"cpus"`
	Mems []float64 `bson:"mems"`
}

//GetMonitorDatas :
func GetMonitorDatas(host string, len int) (MonitorDatas, error) {
	var mds MonitorDatas

	var res []client.Result
	var qs string
	var err error
	var ll int
	// md = MonitorData{}
	mds.Host = host
	if len < 1 {
		ll = 60
	} else {
		ll = len
	}
	for i := 0; i < ll; i++ {
		mds.CPUs = append(mds.CPUs, 0)
		mds.Mems = append(mds.Mems, 0)
	}
	//get Cpus
	qs = "select usage_system+usage_user  from cpu  where cpu='cpu-total' and host='" + host + "' order by time desc limit " + strconv.Itoa(ll)
	print(qs)
	res, err = db.QueryInfluxDB(db.Conntsdb, "telegraf", qs)
	if err != nil {
		fmt.Printf("err !")
	}
	for j, row := range res[0].Series[0].Values {
		fc, _ := strconv.ParseFloat(fmt.Sprint(row[1]), 64)
		mds.CPUs[ll-j-1] = (math.Ceil(fc * 100)) / 100
	}
	//get Mems
	qs = "select 100-available_percent  from mem where host='" + host + "' order by time desc limit " + strconv.Itoa(ll)
	res, err = db.QueryInfluxDB(db.Conntsdb, "telegraf", qs)
	if err != nil {
		fmt.Printf("err !")
	}
	for k, row := range res[0].Series[0].Values {
		fm, _ := strconv.ParseFloat(fmt.Sprint(row[1]), 64)
		mds.Mems[ll-k-1] = (math.Ceil(fm * 100)) / 100
	}

	//var v float64

	return mds, err
}
