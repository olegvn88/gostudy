package properties

import (
	"github.com/magiconair/properties"
)

type PropList struct {
	mySQLLogin, mySQLPassword, mysqlAddress, memcachedAddress, postgresqlLogin, postgresqlPassword, postgresAddress, databaseName string
}

var (
	ok       bool
	propData = PropList{
		"mySQLLogin",
		"mySQLPassword",
		"mysqlAddress",
		"memcachedAddress",
		"postgresqlLogin",
		"postgresqlPassword",
		"postgresAddress",
		"databaseName",
	}
)

func getProperties() *properties.Properties {
	return properties.MustLoadFile("/home/onest/go/src/github.com/olegvn88/gostudy/app/main/app/database/properties/prop.config", properties.UTF8)
}

func (p PropList) GetPostgressLogin() string {
	p.postgresqlLogin, ok = getProperties().Get(propData.postgresqlLogin)
	if ok != true {
		return "Key is not found"
	}
	return p.postgresqlLogin
}

func (p PropList) GetPostgressPassword() string {
	p.postgresqlPassword, ok = getProperties().Get(propData.postgresqlPassword)
	if ok != true {
		return "Key is not found"
	}
	return p.postgresqlPassword
}

func (p PropList) GetMySQLLogin() string {
	p.mySQLLogin, ok = getProperties().Get(propData.mySQLLogin)
	if ok != true {
		return "Key is not found"
	}
	return p.mySQLLogin
}

func (p PropList) GetMySQLPassword() string {
	p.mySQLPassword, ok = getProperties().Get(propData.mySQLPassword)
	if ok != true {
		return "Key is not found"
	}
	return p.mySQLPassword
}

func (p PropList) GetMySQLAddress() string {
	p.mysqlAddress, ok = getProperties().Get(propData.mysqlAddress)
	if ok != true {
		return "Key is not found"
	}
	return p.mysqlAddress
}

func (p PropList) GetPostgresAddress() string {
	p.postgresAddress, ok = getProperties().Get(propData.postgresAddress)
	if ok != true {
		return "Key is not found"
	}
	return p.postgresAddress
}

func (p PropList) GetMemcachedAddress() string {
	p.memcachedAddress, ok = getProperties().Get(propData.memcachedAddress)
	if ok != true {
		return "Key is not found"
	}
	return p.memcachedAddress
}

func (p PropList) GetDatabaseName() string {
	p.databaseName, ok = getProperties().Get("databaseName")
	if ok != true {
		return "Key is not found"
	}
	return p.databaseName
}
