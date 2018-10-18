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

type RealTrailingBucketCounter struct {
	buckets         ConveyorQueue
	secsPerBucket   int64
	lastUpdatedTime int64
}

// NewTrailingBucketCounter3(30, 60)は、直近30分の時間バケツを追跡する。
func NewRealTrailingBucketCounter(numBuckets int, secsPerBucket int) *RealTrailingBucketCounter {
	b := NewRealConveyorQueue(numBuckets)
	spb := int64(secsPerBucket)
	now := time.Now().Unix()
	return &RealTrailingBucketCounter{buckets: b, secsPerBucket: spb, lastUpdatedTime: now}
}

func (tbc *RealTrailingBucketCounter) Add(count int, now int64) {
	tbc.update(now)
	tbc.buckets.AddToBack(count)
}

func (tbc *RealTrailingBucketCounter) TrailingCount(now int64) int {
	tbc.update(now)
	return tbc.buckets.TotalSum()
}

func (tbc *RealTrailingBucketCounter) update(now int64) {
	diffTime := now - tbc.lastUpdatedTime
	numShift := diffTime / tbc.secsPerBucket

	tbc.buckets.Shift(int(numShift))
	tbc.lastUpdatedTime = now

	// Send option
	//currentBuckets := now / tbc.secsPerBucket
	//lastUpdateBuckets := tbc.lastUpdatedTime / tbc.secsPerBucket
	//tbc.buckets.Shift(int(currentBuckets - lastUpdateBuckets))
	//tbc.lastUpdatedTime = now
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

type RealConveyorQueue struct {
	queue []int
}

func NewRealConveyorQueue(numQueue int) *RealConveyorQueue {
	queue := make([]int, numQueue)
	return &RealConveyorQueue{queue: queue}
}

func (q *RealConveyorQueue) AddToBack(count int) {
	panic("implement me")
}

func (q *RealConveyorQueue) Shift(numShift int) {
	panic("implement me")
}

func (q *RealConveyorQueue) TotalSum() int {
	panic("implement me")
}
