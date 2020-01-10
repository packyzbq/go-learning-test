package mock

import (
	"bytes"
	"reflect"
	"testing"
)

// SpySleep 监控 Sleep 函数调用次数
type SpySleep struct {
	Calls int
}

func (s *SpySleep) Sleep() {
	s.Calls += 1
}

// 用于监控 countdown 和 sleep 调用顺序
type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// 输出倒计时 3 2 1 Go
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	//sleeper := ConfigurableSleep{duration: 2 * time.Second}
	t.Run("sleep count", func(t *testing.T) {
		sleeper := &SpySleep{Calls: 0}
		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if sleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
		}
	})

	// 注意 这种测试会暴露算法内部
	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
