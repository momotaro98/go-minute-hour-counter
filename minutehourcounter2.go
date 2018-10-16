package main

import "time"

type MinuteHourCounter2 struct {
	minuteEvents []Event
	hourEvents   []Event
	minuteCount  int
	hourCount    int
}

func (c *MinuteHourCounter2) Add(count int) {
	nowSecs := time.Now().Unix()
	c.ShiftOldEvents(nowSecs)

	// 1分間のリストに流し込む
	c.minuteEvents = append(c.minuteEvents, Event{count, nowSecs})

	c.minuteCount += count
	c.hourCount += count
}

func (c *MinuteHourCounter2) MinuteCount() int {
	c.ShiftOldEvents(time.Now().Unix())
	return c.minuteCount
}

func (c *MinuteHourCounter2) HourCount() int {
	c.ShiftOldEvents(time.Now().Unix())
	return c.hourCount
}

func (c *MinuteHourCounter2) ShiftOldEvents(nowSecs int64) {
	minuteAgo := nowSecs - 60
	hourAgo := nowSecs - 3600

	// 1分以上経過したイベントを'minuteCount'から'hourCount'へと移動する。
	// (1時間以上経過した古いイベントは次のループで削除する)
	for len(c.minuteEvents) > 0 && c.minuteEvents[0].time <= minuteAgo {
		c.hourEvents = append(c.hourEvents, c.minuteEvents[0])

		c.minuteCount -= c.minuteEvents[0].count
		c.minuteEvents = c.minuteEvents[:len(c.minuteEvents)-1]
	}

	// 1時間以上経過した古いイベントを'hour_events'から削除する
	for len(c.hourEvents) > 0 && c.hourEvents[0].time <= hourAgo {
		c.hourCount -= c.hourEvents[0].count
		c.hourEvents = c.hourEvents[:len(c.hourEvents)-1]
	}
}
