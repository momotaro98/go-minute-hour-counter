package minutehourcounter

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

func BenchmarkMinuteHourCounter1_MinuteCount(b *testing.B) {
	counter := NewMinuteHourCounter1()
	loadCounterMinuteCount(b, counter)
}

func BenchmarkMinuteHourCounter2_MinuteCount(b *testing.B) {
	counter := NewMinuteHourCounter2()
	loadCounterMinuteCount(b, counter)
}

func BenchmarkMinuteHourCounter3_MinuteCount(b *testing.B) {
	counter := NewMinuteHourCounter3()
	loadCounterMinuteCount(b, counter)
}

func loadCounterMinuteCount(b *testing.B, counter MinuteHourCounter) {
	for i := 0; i < 10000; i++ {
		counter.Add(1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.MinuteCount()
	}
}

func BenchmarkMinuteHourCounter1_HourCount(b *testing.B) {
	counter := NewMinuteHourCounter1()
	loadCounterHourCount(b, counter)
}

func BenchmarkMinuteHourCounter2_HourCount(b *testing.B) {
	counter := NewMinuteHourCounter2()
	loadCounterHourCount(b, counter)
}

func BenchmarkMinuteHourCounter3_HourCount(b *testing.B) {
	counter := NewMinuteHourCounter3()
	loadCounterHourCount(b, counter)
}

func loadCounterHourCount(b *testing.B, counter MinuteHourCounter) {
	for i := 0; i < 10000; i++ {
		counter.Add(1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.HourCount()
	}
}
