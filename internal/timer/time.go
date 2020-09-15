package timer

import "time"

/**
* @Author: super
* @Date: 2020-09-14 21:39
* @Description:
**/

func GetNowTime() time.Time{
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		return time.Now()
	}
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error){
	duration, err := time.ParseDuration(d)
	if err !=  nil{
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}

