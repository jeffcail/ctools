package time

import (
	"time"
)

const layout = "2006-01-02 15:04:05"

// GetYear
func GetYear() int {
	return time.Now().Year()
}

// GetMonth
func GetMonth() time.Month {
	return time.Now().Month()
}

// GetDay
func GetDay() int {
	return time.Now().Day()
}

// GetHour
func GetHour() int {
	return time.Now().Hour()
}

// GetMinute
func GetMinute() int {
	return time.Now().Minute()
}

// GetSecond
func GetSecond() int {
	return time.Now().Second()
}

// GetNanosecond
func GetNanosecond() int {
	return time.Now().Nanosecond()
}

// GetCurrentTimeStamp
func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// GetCurrentTimeUnixNano
func GetCurrentTimeUnixNano() int64 {
	return time.Now().UnixNano()
}

// GetCurrentTimeString
func GetCurrentTimeString() string {
	return time.Now().Format(layout)
}

// TimestampToString
func TimestampToString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// TimeStringToTime
func TimeStringToTime(str string) time.Time {
	time, _ := time.ParseInLocation(layout, str, time.Local)
	return time
}

// TimeToTimeString
func TimeToTimeString(time time.Time) string {
	return time.Format(layout)
}

// GetWeekday
func GetWeekday(time time.Time) int {
	return int(time.Weekday())
}

// GetCurrentAfterOrBeforeDays
// 获取当前 之前几天的日期 如: -2 当前时间两天前
// 获取当前 之后几天的日子 如: 2  当前时间两天后
func GetCurrentAfterOrBeforeDays(days int) time.Time {
	return time.Now().AddDate(0, 0, days)
}

// GetCurrentZeroTime
// 获取当前零点的时间
func GetCurrentZeroTime() string {
	currentTime := time.Now()
	zeroTime := time.Date(GetYear(), GetMonth(), GetDay(), 0, 0, 0, 0, currentTime.Location())
	return zeroTime.Format(layout)
}

// SubDays
// 计算日期相差多少天
// 返回值day>0, t1晚于t2; day<0, t1早于t2
func SubDays(t1, t2 time.Time) (day int) {
	swap := false
	if t1.Unix() < t2.Unix() {
		t_ := t1
		t1 = t2
		t2 = t_
		swap = true
	}

	day = int(t1.Sub(t2).Hours() / 24)

	// 计算在被24整除外的时间是否存在跨自然日
	if int(t1.Sub(t2).Milliseconds())%86400000 > int(86400000-t2.Unix()%86400000) {
		day += 1
	}

	if swap {
		day = -day
	}

	return
}
