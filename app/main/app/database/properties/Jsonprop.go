package properties

import (
	"encoding/json"
	"io/ioutil"
)

type jsonData struct {
	MySQLLogin         string `json:"mySQLLogin"`
	MySQLPassword      string `json:"mySQLPassword"`
	MysqlAddress       string `json:"mySQLAddress"`
	MemcachedAddress   string `json:"memcachedAddress"`
	PostgresqlLogin    string `json:"postgresqlLogin"`
	PostgresqlPassword string `json:"postgresqlPassword"`
	PostgresAddress    string `json:"postgresAddress"`
	DatabaseName       string `json:"databaseName"`
}

func JsonProp() jsonData {
	jData, err := ioutil.ReadFile("/home/onest/go/src/github.com/olegvn88/gostudy/app/main/app/database/properties/config/properties.json")
	if err != nil {
		panic(err)
	}
	data := jsonData{}
	err = json.Unmarshal(jData, &data)
	if err != nil {
		panic(err)
	}
	return data
}
