package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// Conntsdb :
var Conntsdb client.Client

// ConnInflux :
func ConnInflux(Addr, User, Pwd string) client.Client {
	var hc client.HTTPConfig
	hc.Addr = Addr
	hc.Username = User
	hc.Password = Pwd
	cli, err := client.NewHTTPClient(hc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hc)
	Conntsdb = cli
	return cli
}

// QueryInfluxDB :
func QueryInfluxDB(cli client.Client, db, cmd string) (res []client.Result, err error) {
	q := client.NewQuery(cmd, db, "")
	fmt.Println("Enter QueryInfluxDB()")
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
		//fmt.Print("----------------")
		//fmt.Println(res)
	} else {
		return res, err
	}
	return res, nil
}

// WritesPoints :
func WritesPoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "snns",
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   20.1,
		"system": 43.3,
		"user":   86.6,
	}

	pt, err := client.NewPoint(
		"cpu_usage",
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		log.Fatal(err)
	}
}
