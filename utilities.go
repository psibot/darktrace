package darktrace

import "time"

func unixMicroSecondToTime(unixMS int) time.Time {
	tUnixMicro := int64(unixMS)
	tUnix := tUnixMicro / int64(time.Millisecond)
	tUnixNanoRemainder := (tUnixMicro % int64(time.Millisecond)) * int64(time.Microsecond)
	return time.Unix(tUnix, tUnixNanoRemainder)
}

func unixMilliSecondToTime(unixMS int) time.Time {
	tUnixMilli := int64(unixMS)
	tUnix := tUnixMilli / int64(time.Microsecond)
	tUnixNanoRemainder := (tUnixMilli % int64(time.Microsecond)) * int64(time.Millisecond)
	return time.Unix(tUnix, tUnixNanoRemainder)
}

func unixNanoSecondToTime(unixNS int) time.Time {
	tUnixNano := int64(unixNS)
	tUnix := tUnixNano / int64(time.Second)
	tUnixNanoRemainder := tUnixNano % int64(time.Second)
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
