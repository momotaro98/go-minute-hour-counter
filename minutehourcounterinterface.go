package minutehourcounter

import "github.com/pkg/errors"

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

func CounterFactory(whichNumberCounter int) (MinuteHourCounter, error) {
	switch whichNumberCounter {
	case 1:
		return NewMinuteHourCounter1(), nil
	case 2:
		return NewMinuteHourCounter2(), nil
	case 3:
		return NewMinuteHourCounter3(), nil
	}
	return nil, errors.Errorf("Invalid number: %d", whichNumberCounter)
}
