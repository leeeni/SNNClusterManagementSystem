package repository

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/shopspring/decimal"

	"snns_srv/db"
)

// CPU :
type CPU struct {
	// Uid: Primary key (_id)
	Host    string  `bson:"host"`
	CPUPerc float64 `bson:"cpuperc"`
}

//GetCPUByHost :
func GetCPUByHost(name string) (CPU, error) {
	cpu := CPU{}
	qs := "select usage_system+usage_user  from cpu  where cpu='cpu-total' and host ='" + name + "'  order by time desc limit 1"
	res, err := db.QueryInfluxDB(db.Conntsdb, "telegraf", qs)
	if err != nil {
		fmt.Printf("err !")
	}
	//var v float64
	row := res[0].Series[0].Values[0]
	vs := fmt.Sprint(row[1])
	f, _ := strconv.ParseFloat(vs, 64)
	v, _ := decimal.NewFromFloat(f).Round(2).Float64()
	fmt.Println(v)
	cpu.Host = name
	cpu.CPUPerc = math.Abs((v*10 + (rand.Float64()*10 - 4)))
	if cpu.CPUPerc > 100 {
		cpu.CPUPerc = 100.0
	}
	return cpu, err
}
