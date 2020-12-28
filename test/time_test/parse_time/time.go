package parse_time

import (
	"time"
)

var bjLocation *time.Location

func init() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		bjLocation = time.FixedZone("CST", 8*3600)
	} else {
		bjLocation = loc
	}
}

func StringToTime(layout, strTime string) (time.Time, error) {
	return time.ParseInLocation(layout, strTime, bjLocation)
}

// TimestampToDate 时间戳转日期字符串
func TimestampToDate(timeStamp int64, format string) string {
	//返回time对象
	t := time.Unix(timeStamp, 0)
	return t.Format(format)
}
