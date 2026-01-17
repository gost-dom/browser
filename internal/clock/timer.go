package clock

import "time"

type Timer interface {
	Time() time.Time
	SetTime(time.Time)
}

type timer time.Time

func (t *timer) Time() time.Time {
	if t == nil {
		return time.UnixMilli(0)
	}
	return time.Time(*t)
}

func (t *timer) SetTime(newTime time.Time) {
	if t != nil {
		*t = timer(newTime)
	}
}
