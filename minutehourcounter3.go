package main

import "time"

type MinuteHourCounter3 struct {
	minuteCounts TrailingBucketCounter
	hourCounts   TrailingBucketCounter
}

func NewMinuteHourCounter3() *MinuteHourCounter3 {
	m := NewRealTrailingBucketCounter( /*numBuckets*/ 60 /*numBuckets*/, 1)
	h := NewRealTrailingBucketCounter( /*numBuckets*/ 60 /*numBuckets*/, 60)
	return &MinuteHourCounter3{minuteCounts: m, hourCounts: h}
}

func (c *MinuteHourCounter3) Add(count int) {
	now := time.Now().Unix()
	c.minuteCounts.Add(count, now)
	c.hourCounts.Add(count, now)
}

func (c *MinuteHourCounter3) MinuteCount() int {
	now := time.Now().Unix()
	return c.minuteCounts.TrailingCount(now)
}

func (c *MinuteHourCounter3) HourCount() int {
	now := time.Now().Unix()
	return c.hourCounts.TrailingCount(now)
}

// TrailingBucketCounter は時間バケツN個のカウントを保持する。
type TrailingBucketCounter interface {
	Add(count int, now int64)

	// TrailingCount は最新の合計バケツ分の合計カウントを返す。
	TrailingCount(now int64) int
}

type RealTrailingBucketCounter struct{}

// NewTrailingBucketCounter3(30, 60)は、直近30分の時間バケツを追跡する。
func NewRealTrailingBucketCounter(numBuckets int, secsPerBucket int) *RealTrailingBucketCounter {
	return nil
}

func (RealTrailingBucketCounter) Add(count int, now int64) {
	panic("implement me")
}

func (RealTrailingBucketCounter) TrailingCount(now int64) int {
	panic("implement me")
}

// ConveyorQueue は上限数を持ったキュー。古いデータは端から落ちる。
type ConveyorQueue interface {
	// AddToBack はキューの最後の値を増加する。
	AddToBack(count int)

	// Shift はキューの値を'numShift'分だけシフトする。
	// 新しい項目は0で初期化する。
	// 最古の項目はmax_items以下なら削除する。
	Shift(numShift int)

	// TotalSum は現在のキューに含まれる項目の合計値を返す。
	TotalSum() int
}
