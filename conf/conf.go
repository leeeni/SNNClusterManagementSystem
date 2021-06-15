package conf

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

//config 配置文件数据结构定义
type config struct {
	Name     string
	Hostid   string
	Hostname string
	Tsdb     struct {
		Server string
		Port   int
		User   string
		Pwd    string
	}
}

// Conf :  配置数据变量
var Conf config

// LoadConf :
// @title			LoadConf
// @description		保存配置文件
// @auth			Gene             时间（2020/6/18   10:57 ）
// @param			fn string 文件名
// @return			nil
func LoadConf(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := toml.NewDecoder(f).Decode(&Conf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Conf)
}

// SaveConf :
// @title			SaveConf
// @description		保存配置文件
// @auth			Gene             时间（2020/6/18   10:57 ）
// @param			fn string 文件名
// @return			nil
func SaveConf(fn string) {
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR, 6)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := toml.NewEncoder(f).Encode(&Conf); err != nil {
		fmt.Println(err)
		return
	}
}
