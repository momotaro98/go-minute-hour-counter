package main

import (
	"testing"
)

func BenchmarkMinuteHourCounter1_Add(b *testing.B) {
	counter := NewMinuteHourCounter1()
	loadCounterAdd(b, counter)
}

func BenchmarkMinuteHourCounter2_Add(b *testing.B) {
	counter := NewMinuteHourCounter2()
	loadCounterAdd(b, counter)
}

func BenchmarkMinuteHourCounter3_Add(b *testing.B) {
	counter := NewMinuteHourCounter3()
	loadCounterAdd(b, counter)
}

func loadCounterAdd(b *testing.B, counter MinuteHourCounter) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.Add(1)
	}
}

func BenchmarkMinuteHourCounter1_Count(b *testing.B) {
	counter := NewMinuteHourCounter1()
	loadCounterCount(b, counter)
}

func BenchmarkMinuteHourCounter2_Count(b *testing.B) {
	counter := NewMinuteHourCounter2()
	loadCounterCount(b, counter)
}

func BenchmarkMinuteHourCounter3_Count(b *testing.B) {
	counter := NewMinuteHourCounter3()
	loadCounterCount(b, counter)
}

func loadCounterCount(b *testing.B, counter MinuteHourCounter) {
	for i := 0; i < 10000; i++ {
		counter.Add(1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.MinuteCount()
	}
	for i := 0; i < b.N; i++ {
		counter.HourCount()
	}
}
