package dao

import (
	"log"
	"time"
)

func SaveCode(phoneNumber, code string) (err error) {
	err = rdb.Set(phoneNumber, code, time.Minute*30).Err()
	return
}

func GetCode(phoneNumber string) (code string, err error) {
	code, err = rdb.Get(phoneNumber).Result()
	return
}

func DelCode(phoneNumber string) {
	err := rdb.Del(phoneNumber).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

// SavePhoneNumberByRT 存储rt
func SavePhoneNumberByRT(RT, phoneNumber string) (err error) {
	err = rdb.Set(RT, phoneNumber, time.Hour*30).Err()
	return
}

// GetPhoneNumberByRT 获取rt
func GetPhoneNumberByRT(RT string) (phoneNumber string, err error) {
	phoneNumber, err = rdb.Get(RT).Result()
	return
}
