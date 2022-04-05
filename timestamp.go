package utils

import "time"

//MakeTimeStamp func
func MakeTimeStamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

//SetTimeStamp func
func SetTimeStamp(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
