package tools

import (
	"time"
)

/**
* @Author: super
* @Date: 2020-08-21 10:51
* @Description:
**/

func GetNowTime() string {
	t := time.Now()
	return t.String()
}


func GetTimeStamp() int64 {
	unix := time.Now().Unix()
	return unix
}

func GetTimeFormat() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}

func StampToTime() string {
	unix := time.Now().Unix()
	format := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	return format
}

func ParseToTime(str string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	formatTime,err:=time.ParseInLocation("2006-01-02 15:04:05",str,loc)

	if err!=nil{
		return "", err
	}
	return formatTime.String(), nil
}

func ParseToTimeStamp(str string) (int64, error) {
	loc, _ := time.LoadLocation("Local")
	formatTime,err:=time.ParseInLocation("2006-01-02 15:04:05",str,loc)

	if err!=nil{
		return 0, err
	}
	return formatTime.Unix(), nil
}