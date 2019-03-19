package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/olegvn88/gostudy/app/main/app/database/properties"
)

var (
	memcacheClient *memcache.Client
	prop           = properties.JsonProp()
)

//Memcache - мапка ,которая живет в памяти
//Там хранят данные, которые не жалко потерять, но к которым нужен быстрый доступ
func main() {
	memcacheAddresses := []string{prop.MemcachedAddress}
	memcacheClient = memcache.New(memcacheAddresses...)

	mKey := "record_21"
	mKey2 := "record_22"
	setRecord(mKey2)
	//addItem(mKey)
	item := getRecord(mKey)

	fmt.Printf("first get %+v ", item.Value)

	afterIncrement, err := memcacheClient.Increment(mKey, uint64(2))
	PanicOnError(err)
	fmt.Println("afterincrement by 2", mKey, "is", afterIncrement)

	afterDerement, err := memcacheClient.Decrement(mKey, uint64(1))
	PanicOnError(err)
	fmt.Println("afterdecrement by 1", mKey, "is", afterDerement)

	//получаем несколько записей за раз
	mkeys := []string{mKey, "record_22"}
	println("get multiple", mkeys)
	multipleItems, err := memcacheClient.GetMulti(mkeys)
	PanicOnError(err)
	fmt.Println("multipleItems", multipleItems)
	fmt.Println(getRecord(mKey2).Value)

	return
}

func getRecord(mKey string) *memcache.Item {
	fmt.Println("get ", mKey)
	//получаем одиночную запись
	item, err := memcacheClient.Get(mKey)
	//если записи нет, то для этого есть существует специальная ошибка, ее надо обрабатывать отдельно, это почти штатная ситуация, а не что-то страшное
	if err == memcache.ErrCacheMiss {
		fmt.Println("record is not found in memcache")
		return nil
	} else if err != nil {
		PanicOnError(err)
	}
	return item
}

var ttl = 5

func setRecord(mKey string) {
	err := memcacheClient.Set(&memcache.Item{
		Key:   mKey,
		Value: []byte("1"),
		//Expiration: int32(ttl),
	})
	PanicOnError(err)

	//time.Sleep(time.Microsecond)
	//time.Sleep(4 * time.Second)
}

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func addItem(mKey string) {
	err := memcacheClient.Add(&memcache.Item{
		Key:   mKey,
		Value: []byte("1"),
		//Expiration: int32(ttl),
	})
	// если запись не была одобрена, то вернется соответствующая ошибка
	if err == memcache.ErrNotStored {
		fmt.Println("Record is not stored")
	} else if err != nil {
		PanicOnError(err)
	}
}
