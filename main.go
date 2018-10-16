package main

import (
	"fmt"
)

// MinuteHourCounter は直近1分間および直近1時間の累積カウントを記録する。
// 例えば、帯域幅の使用状況を確認するのに使える。
type MinuteHourCounter interface {
	// Add は新しいデータ点を追加する(count >= 0)。
	// Add がコールされてそれから1分間は、MinuteCount()の返す値が+countだけ増える。
	// Add がコールされてそれから1分間は、HourCount()の返す値が+countだけ増える。
	Add(count int)

	// MinuteCount は直近60秒間の累積カウントを返す。
	MinuteCount() int

	// HourCount は直近3600秒間の累積カウントを返す。
	HourCount() int
}

type Event struct {
	count int
	time  int64
}

func CounterFactory() MinuteHourCounter {
	var c MinuteHourCounter
	// c = &MinuteHourCounter1{}
	c = &MinuteHourCounter2{}
	return c
}

func main() {
	fmt.Println("vim-go")

	counter := CounterFactory()
	for i := 0; i < 10; i++ {
		counter.Add(1)
	}
	fmt.Println(counter.MinuteCount())
	fmt.Println(counter.HourCount())
}
