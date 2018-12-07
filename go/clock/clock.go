package clock

import "strconv"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	hour = dealWithHour(hour)

	if minute >= 60 {
		hour = dealWithHour(hour + minute/60)
		minute = minute % 60
	}
	if minute < 0 && minute >= -60 {
		hour = dealWithHour(hour - 1)
		minute += 60
	} else if minute < -60 {
		if minute%60 == 0 {
			hour = dealWithHour(hour + minute/60)
			minute = 0
		} else {
			hour = dealWithHour(hour - 1 + minute/60)
			minute = 60 + minute%60
		}
	}

	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

func dealWithHour(hour int) int {
	if hour < 0 && hour > -24 {
		hour += 24
	} else if hour == -24 {
		hour = 0
	} else if hour < -24 {
		hour = 24 + hour%24
	}

	if hour == 24 {
		hour = 0
	} else if hour > 24 {
		hour %= 24
	}

	return hour
}

func (c Clock) String() string {
	res := ""
	if c.Hour < 10 {
		res += "0" + strconv.Itoa(c.Hour)
	} else {
		res += strconv.Itoa(c.Hour)
	}

	res += ":"

	if c.Minute < 10 {
		res += "0" + strconv.Itoa(c.Minute)
	} else {
		res += strconv.Itoa(c.Minute)
	}

	return res
}

func (c Clock) Add(minutes int) Clock {

	c.Minute += minutes

	if c.Minute >= 60 {
		c.Hour = dealWithHour(c.Hour + c.Minute/60)
		c.Minute = c.Minute % 60
	}
	if c.Minute < 0 && c.Minute >= -60 {
		c.Hour = dealWithHour(c.Hour - 1)
		c.Minute += 60
	} else if c.Minute < -60 {
		if c.Minute%60 == 0 {
			c.Hour = dealWithHour(c.Hour + c.Minute/60)
			c.Minute = 0
		} else {
			c.Hour = dealWithHour(c.Hour - 1 + c.Minute/60)
			c.Minute = 60 + c.Minute%60
		}
	}

	return c
}
