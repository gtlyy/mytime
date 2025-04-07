// 功能：提供时间处理函数
package mytime

import (
	"fmt"
	"time"

	"github.com/gtlyy/myfun"
)

// ts -> Str(UTC): 1633603737443 -> 2021-10-07T10:48:57.443Z (or other format)
func TsToStr(ts string, format string) string {
	sec := ts[0:10]
	msec := ts[10:len(ts)]
	tt := time.Unix(myfun.StringToInt64(sec), 1000000*myfun.StringToInt64(msec)).UTC()
	return tt.Format(format)
}

// ts -> Str(CST): 1633603737443 -> 2021-10-07T18:48:57.443Z (or other format)
func TsToStrCST(ts string, format string) string {
	sec := ts[0:10]
	msec := ts[10:len(ts)]
	tt := time.Unix(myfun.StringToInt64(sec), 1000000*myfun.StringToInt64(msec)).UTC()
	// tt = tt.Add(8 * time.Hour)
	tt = tt.In(time.FixedZone("CST", 8*3600))
	return tt.Format(format)
}

// ts -> ISO Str(UTC): 1633603737443 -> 2021-10-07T10:48:57.443Z
func TsToISO(ts string) string {
	return TsToStr(ts, "2006-01-02T15:04:05.999Z")
}

// ts -> ISO Str(CST): 1633603737443 -> 2021-10-07T18:48:57.443Z
func TsToISOCST(ts string) string {
	return TsToStrCST(ts, "2006-01-02T15:04:05.000Z") //.999Z .000Z not the same
}

// ts -> time(UTC): 1633603737443 -> 2021-10-07 10:48:57.443 +0000 UTC
func TsToTime(ts string) time.Time {
	sec := ts[0:10]
	msec := ts[10:len(ts)]
	tt := time.Unix(myfun.StringToInt64(sec), 1000000*myfun.StringToInt64(msec)).UTC()
	return tt
}

// ts -> time(CST): 1633603737443 -> 2021-10-07 18:48:57.443 +0000 CST
func TsToTimeCST(ts string) time.Time {
	s := TsToISOCST(ts)
	zone := time.Local
	t, _ := time.ParseInLocation("2006-01-02T15:04:05.999Z", s, zone)
	return t
}

// iso -> ts: 2021-10-07T10:48:57.443Z -> 1633603737443
func ISOToTs(strISO string) string {
	format := "2006-01-02T15:04:05.999Z"
	tt, err := time.Parse(format, strISO)
	if err != nil {
		fmt.Println("Error ISOToTs().")
	}
	return myfun.Int64ToString(tt.UnixNano() / 1000000)
}

// iso(CST) -> ts: 2021-10-07T18:48:57.443Z -> 1633603737443
func ISOCSTToTs(strISOCST string) string {
	format := "2006-01-02T15:04:05.999Z"
	tt, err := time.ParseInLocation(format, strISOCST, time.Local)
	if err != nil {
		fmt.Println("Error ISOToTs().")
	}
	return myfun.Int64ToString(tt.UnixNano() / 1000000)
}

// iso -> time: 2021-10-07T10:48:57.443Z -> time(2021-10-07 10:48:57.443 +0000 UTC)
func ISOToTime(strISO string) time.Time {
	tt, _ := time.Parse("2006-01-02T15:04:05.999Z", strISO)
	return tt
}

// iso(CST) -> time: 2021-10-07T18:48:57.443Z -> time(2021-10-07 18:48:57.443 +0000 CST)
func ISOCSTToTime(strISOCST string) time.Time {
	format := "2006-01-02T15:04:05.999Z"
	tt, err := time.ParseInLocation(format, strISOCST, time.Local)
	if err != nil {
		fmt.Println("Error ISOToTs().")
	}
	return tt
}

// iso to iso_cst: 2021-10-07T10:48:57.443Z  -> 2021-10-07T18:48:57.443Z
func ISOToISOCST(strISO string) string {
	ts1 := ISOToTs(strISO)
	iso_cst1 := TsToISOCST(ts1)
	return iso_cst1
}

// iso_cst to iso: 2021-10-07T18:48:57.443Z  -> 2021-10-07T10:48:57.443Z
func ISOCSTToISO(strISOCST string) string {
	ts1 := ISOCSTToTs(strISOCST)
	iso1 := TsToISO(ts1)
	return iso1
}

// time -> ts: time(2021-10-07 10:48:57.443 +0000 UTC) -> 1633603737443
func TimeToTs(tt time.Time) string {
	return myfun.Int64ToString(tt.UnixNano() / 1000000)
}

// time(CST) -> ts: time(2021-10-07 18:48:57.443 +0000 CST) -> 1633603737443
func TimeCSTToTs(tt time.Time) string {
	return myfun.Int64ToString(tt.UnixNano()/1000000 - 8*60*60*1000)
}

// time(UTC) -> str: time(2021-10-07 10:48:57.443 +0000 UTC) -> 2021-10-07T10:48:57.443Z
func TimeToStr(tt time.Time, format string) string {
	return tt.Format(format)
}

// time(UTC) -> str: time(2021-10-07 10:48:57.443 +0000 UTC) -> 2021-10-07T10:48:57.443Z
func TimeToISO(tt time.Time) string {
	return TimeToStr(tt, "2006-01-02T15:04:05.999Z")
}

// time(UTC) -> str: time(2021-10-07 10:48:57.443 +0000 UTC) -> 2021-10-07T10:48:57.443Z
func TimeToISOCST(tt time.Time) string {
	return TimeToStr(tt, "2006-01-02T15:04:05.999Z")
}

// Get a iso time. eg: 2018-03-16T18:02:48.284Z
func ISONow() string {
	utcTime := time.Now().UTC()
	// iso := utcTime.String()
	// isoBytes := []byte(iso)
	// iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	// return iso
	return utcTime.Format("2006-01-02T15:04:05.999Z")
}

// Get iso CST now.
func ISONowCST() string {
	return ISOToISOCST(ISONow())
}

// Get now ts:
func TsNow() string {
	return myfun.Int64ToString(time.Now().UnixNano() / 1000000)
}

// 返回 hour, min, sec, nanoSec  : UTC
func GetHmsnNowUtc() (h, m, s, n int) {
	time_utc := time.Now().UTC()
	return time_utc.Hour(), time_utc.Minute(), time_utc.Second(), time_utc.Nanosecond()
}

// 返回 hour, min, sec, nanoSec  : CST
func GetHmsnNowCst() (h, m, s, n int) {
	time_utc := time.Now().UTC()
	time_cst := time_utc.Add(8 * time.Hour)
	return time_cst.Hour(), time_cst.Minute(), time_cst.Second(), time_cst.Nanosecond()
}

// 将 yyyymmdd 转为 yyyy-mm-ddT09:30:00Z
// 这个不算通用的，以后可以删去。
func ConvertDate(date string) string {
	if len(date) != 8 {
		return ""
	}
	year := date[0:4]
	month := date[4:6]
	day := date[6:8]
	formattedDate := year + "-" + month + "-" + day + "T09:30:00Z"
	return formattedDate
}
