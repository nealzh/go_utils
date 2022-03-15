package schedule_utils

import (
	"bytes"
	"math"
	"strconv"
	"time"
)

type ScheduleTime struct {
	Year     string
	Month    string
	Day      string
	Hour     string
	Minute   string
	Second   string
	Duration string
}

var asterisk = "*"
var dateTimeStr = "2006-01-02 15:04:05"

func ScheduleRunningTimeTrigger(st *ScheduleTime) (bool, error) {

	currentTime := time.Now()
	var scheduleDateTime time.Time

	runningTimeTigger := false

	var scheduleYear, scheduleMonth, scheduleDay, scheduleHour, scheduleMinute, scheduleSecond, scheduleDuration int
	var err error

	//fmt.Println("at ScheduleRunningTimeTrigger scheduleTime:", st)

	// Year
	if st.Year == asterisk {
		scheduleYear = currentTime.Year()
	} else {
		scheduleYear, err = strconv.Atoi(st.Year)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after year.")

	// Month
	if st.Month == asterisk {
		scheduleMonth = int(currentTime.Month())
	} else {
		scheduleMonth, err = strconv.Atoi(st.Month)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after month.")

	// Day
	if st.Day == asterisk {
		scheduleDay = currentTime.Day()
	} else {
		scheduleDay, err = strconv.Atoi(st.Day)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after day.")

	// Hour
	if st.Hour == asterisk {
		scheduleHour = currentTime.Hour()
	} else {
		scheduleHour, err = strconv.Atoi(st.Hour)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after hour.")

	// Minute
	if st.Minute == asterisk {
		scheduleMinute = currentTime.Minute()
	} else {
		scheduleMinute, err = strconv.Atoi(st.Minute)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after minure.")

	// Second
	if st.Second == asterisk {
		scheduleSecond = currentTime.Second()
	} else {
		scheduleSecond, err = strconv.Atoi(st.Second)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after second.")

	// Duration
	if st.Duration == asterisk {
		scheduleDuration = math.MaxInt64
	} else {
		scheduleDuration, err = strconv.Atoi(st.Duration)
	}

	if err != nil {
		return runningTimeTigger, err
	}
	//fmt.Println("at ScheduleRunningTimeTrigger after duration.")

	var buffer bytes.Buffer

	buffer.WriteString(strconv.Itoa(scheduleYear))

	buffer.WriteString("-")

	monthStr := strconv.Itoa(scheduleMonth)
	if len(monthStr) == 1 {
		buffer.WriteString("0")
	}
	buffer.WriteString(monthStr)

	buffer.WriteString("-")

	dayStr := strconv.Itoa(scheduleDay)
	if len(dayStr) == 1 {
		buffer.WriteString("0")
	}
	buffer.WriteString(dayStr)

	buffer.WriteString(" ")

	hourStr := strconv.Itoa(scheduleHour)
	if len(hourStr) == 1 {
		buffer.WriteString("0")
	}
	buffer.WriteString(hourStr)

	buffer.WriteString(":")

	minuteStr := strconv.Itoa(scheduleMinute)
	if len(minuteStr) == 1 {
		buffer.WriteString("0")
	}
	buffer.WriteString(minuteStr)

	buffer.WriteString(":")

	secondStr := strconv.Itoa(scheduleSecond)
	if len(secondStr) == 1 {
		buffer.WriteString("0")
	}
	buffer.WriteString(secondStr)

	scheduleDateTimeStr := buffer.String()

	//fmt.Println("at ScheduleRunningTimeTrigger after scheduleDateTimeStr.", scheduleDateTimeStr)

	scheduleDateTime, err = time.ParseInLocation(dateTimeStr, scheduleDateTimeStr, time.Local)

	//fmt.Println("at ScheduleRunningTimeTrigger after scheduleDateTime.", scheduleDateTime, currentTime, err)

	if err != nil {
		return runningTimeTigger, err
	}

	if scheduleDateTime.After(currentTime) {
		return runningTimeTigger, err
	}

	startSeonds := int(currentTime.Sub(scheduleDateTime).Seconds())

	if startSeonds < scheduleDuration {
		runningTimeTigger = true
	}

	//fmt.Printf("startSeonds: %v, scheduleDuration: %v, runningTimeTigger: %v.\n", strconv.Itoa(startSeonds), strconv.Itoa(scheduleDuration), runningTimeTigger)
	//fmt.Println()

	return runningTimeTigger, err
}
