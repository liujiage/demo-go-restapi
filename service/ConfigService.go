package service

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
}


/****
Get current app config properites
Return: return Config struct
***/

func GetConfig() Config {
	p := properties.MustLoadFile("./resource/app.properties", properties.UTF8)
	port := p.GetInt("app.port", 8080)
	ip := p.GetString("app.ip", "localhost")
	fmt.Println(port, ip)
	return Config{Port: port, Ip: ip}
}

