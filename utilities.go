package darktrace

import "time"

func UnixMicroSecondToTime(unixMS int64) time.Time {
	tUnix := unixMS / int64(time.Millisecond)
	tUnixNanoRemainder := (unixMS % int64(time.Millisecond)) * int64(time.Microsecond)
	return time.Unix(tUnix, tUnixNanoRemainder)
}

func UnixMilliSecondToTime(unixMS int64) time.Time {
	tUnix := unixMS / int64(time.Microsecond)
	tUnixNanoRemainder := (unixMS % int64(time.Microsecond)) * int64(time.Millisecond)
	return time.Unix(tUnix, tUnixNanoRemainder)
}

func UnixNanoSecondToTime(unixNS int64) time.Time {
	tUnix := unixNS / int64(time.Second)
	tUnixNanoRemainder := unixNS % int64(time.Second)
	return time.Unix(tUnix, tUnixNanoRemainder)
}

func TimeToUnixMicroSecond(t time.Time) int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Microsecond)
}

func TimeToUnixMilliSecond(t time.Time) int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func TimeToUnixNanoSecond(t time.Time) int64 {
	return time.Now().UnixNano()
}
