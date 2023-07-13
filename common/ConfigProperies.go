package common

import (
	"fmt"
	"github.com/magiconair/properties"
)


/****
Config struct
****/
type Config struct{
	Port int
	Ip string
	ImgrateMode string //up and down
}


/****
Get current app config properites
Return: return Config struct
***/

func GetConfig() Config {
	p := properties.MustLoadFile("./resource/app.properties", properties.UTF8)
	port := p.GetInt("app.port", 8080)
	ip := p.GetString("app.ip", "localhost")
	imgrateMode := p.GetString("imgrate.mode", "up")
	fmt.Println(port, ip, imgrateMode)
	return Config{Port: port, Ip: ip, ImgrateMode: imgrateMode}
}

 