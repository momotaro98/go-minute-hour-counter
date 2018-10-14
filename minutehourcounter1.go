package main

import "time"

type Event struct {
	count int
	time  int64
}

type MinuteHourCounter1 struct {
	events []Event
}

func (c *MinuteHourCounter1) Add(count int) {
	c.events = append(c.events, Event{count, time.Now().Unix()})
}

func (c *MinuteHourCounter1) MinuteCount() int {
	return c.CountSince(time.Now().Unix() - 60)
}

func (c *MinuteHourCounter1) HourCount() int {
	return c.CountSince(time.Now().Unix() - 3600)
}

func (c *MinuteHourCounter1) CountSince(cutoff int64) (count int) {
	for _, event := range c.events {
		if event.time <= cutoff {
			break
		}
		count += event.count
	}
	return count
}

// この実装ではパフォーマンスの問題がある。
// 1. これからも大きくなっていく。
// このクラスはすべてのイベントを保持している。つまり、メモリを無限に使
// 用してしまうのだ！ MinuteHourCounterは、1時間よりも不要なイベント
// を自動的に削除するべきだ
// 2. MinuteCountとHourCountが遅すぎる。
// CountSince()メソッドの処理時間はO(n)でありパフォーマンスが悪い。
// MinuteHourCounterは、Add()の呼び出しに対する値をminute_countと
// hour_countとで別々に保持するべきだ。
