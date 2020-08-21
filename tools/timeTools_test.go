package tools

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-21 10:52
* @Description:
**/

func TestGetNowTime(t *testing.T) {
	time := GetNowTime()
	fmt.Println(time)
}

func TestGetTimeStamp(t *testing.T) {
	stamp := GetTimeStamp()
	fmt.Println(stamp)
}

func TestGetTimeFormat(t *testing.T) {
	stamp := GetTimeFormat()
	fmt.Println(stamp)
}

func TestStampToTime(t *testing.T) {
	time := StampToTime()
	fmt.Println(time)
}

func TestParseToTime(t *testing.T) {
	s, err := ParseToTime("2020-04-11 13:33:37")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(s)
}

func TestParseToTimeStamp(t *testing.T) {
	s, err := ParseToTimeStamp("2020-04-11 13:33:37")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(s)
}