package sub

// import (
// 	"fmt"

// 	"gopkg.in/ini.v1"
// )

// type ConfigList struct {
// 	Port      int
// 	DBname    string
// 	SQLDriver string
// }

// var Config ConfigList

// func init() {
// 	cfg, _ := ini.Load("config.ini")
// 	Config = ConfigList{
// 		Port:      cfg.Section("web").Key("port").MustInt(),
// 		DBname:    cfg.Section("db").Key("name").MustString("example.sql"),
// 		SQLDriver: cfg.Section("db").Key("driver").String(),
// 	}
// }

// func IniExample() {
// 	fmt.Printf("%T %v\n", Config.Port, Config.Port)
// 	fmt.Printf("%T %v\n", Config.DBname, Config.DBname)
// 	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
// }
