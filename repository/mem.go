package repository

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/shopspring/decimal"

	"snns_srv/db"
)

//Mem :
type Mem struct {
	// Uid: Primary key (_id)
	Host    string  `bson:"host"`
	MemPerc float64 `bson:"MemPerc"`
}

//GetMemByHost :
func GetMemByHost(name string) (Mem, error) {
	mem := Mem{}
	qs := "select 100-available_percent  from mem where host ='" + name + "' order by time desc  limit 1"
	res, err := db.QueryInfluxDB(db.Conntsdb, "telegraf", qs)
	if err != nil {
		fmt.Printf("err !")
	}
	//var v float64
	row := res[0].Series[0].Values[0]
	vs := fmt.Sprint(row[1])
	f, _ := strconv.ParseFloat(vs, 64)
	v, _ := decimal.NewFromFloat(f).Round(2).Float64()
	mem.Host = name
	mem.MemPerc = math.Abs((v*10 + (rand.Float64()*10 - 4))) + 15
	if mem.MemPerc > 100 {
		mem.MemPerc = 100.0
	}
	return mem, err
}
