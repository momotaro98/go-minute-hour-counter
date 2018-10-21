package minutehourcounter

import "time"

type Event2 struct {
	count int
	time  int64
}

type MinuteHourCounter2 struct {
	minuteEvents []Event2
	hourEvents   []Event2
	minuteCount  int
	hourCount    int
}

func NewMinuteHourCounter2() *MinuteHourCounter2 {
	return &MinuteHourCounter2{}
}

func (c *MinuteHourCounter2) Add(count int) {
	nowSecs := time.Now().Unix()
	c.ShiftOldEvents(nowSecs)

	// 1分間のリストに流し込む
	c.minuteEvents = append(c.minuteEvents, Event2{count, nowSecs})

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

// このCounterの欠点
// 1点目: まず、この設計には柔軟性がない。例えば、直近24時間のカウントを保持したい
// とする。すると、多くのコードに修正が必要になる。ShiftOldEvents()は、わずかに
// 分と時間のデータのやり取りをしているだけの非常に密度の濃い関数である。
// 2点目: 次に、メモリの使用量が多い。高トラフィックのサーバが1秒間に100回もAdd()
// を呼び出したとしよう。直近1時間のデータをすべて保持しているので、約5MBのメモリが
// 必要になる。Add()が呼び出される頻度に関係なく、MinuteHourCounterの使用する
// メモリは一定であるほうが良い。
