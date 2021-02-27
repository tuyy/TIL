package kakao2018

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestHappy7(t *testing.T) {
	cases := []struct{
		n, t, m int
		timetable []string
		want string
	}{
		{1, 1, 5, []string{"08:00","08:01","08:02","08:03"}, "09:00"},
		{2, 10, 2, []string{"09:10", "09:09", "08:00"}, "09:09"},
		{2, 1, 2, []string{"09:00", "09:00", "09:00", "09:00"}, "08:59"},
		{1, 1, 5, []string{"00:01", "00:01", "00:01", "00:01", "00:01"}, "00:00"},
		{1, 1, 1, []string{"23:59"}, "09:00"},
		{10, 60, 45, []string{"23:59","23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59", "23:59"}, "18:00"},
	}

	for _, c := range cases {
		rz := getLastBusTime(c.n, c.t, c.m, c.timetable)
		if rz != c.want {
			t.Fatalf("invalid result. want:%s rz:%s\n", c.want, rz)
		}
	}
}

func getLastBusTime(n, t, m int, timetable []string) string {
	waitingTimes := ConvertToBusTimeSecs(timetable)
	sort.Ints(waitingTimes)

	remainWaitingTimes := waitingTimes
	busTimes := GetBusTimeSecs(n, t)
	sort.Ints(busTimes)
	var lastWaitingTimes []int
	for i, busTime := range busTimes {
		successCnt := 0
		for _, waitingTime := range remainWaitingTimes {
			if waitingTime <= busTime {
				successCnt++

				// 가장 마지막 버스시간인 경우
				if i == len(busTimes)-1 {
					lastWaitingTimes = append(lastWaitingTimes, waitingTime)
				}

				if successCnt == m {
					break
				}
			}
		}

		remainWaitingTimes = waitingTimes[successCnt:]
	}

	if len(lastWaitingTimes) == m {
		return ConvertToBusTimeStr(lastWaitingTimes[len(lastWaitingTimes) - 1] - secPerMin)
	}

	return ConvertToBusTimeStr(busTimes[len(busTimes) - 1])
}

func ConvertToBusTimeStr(timeSec int) string {
	return fmt.Sprintf("%02d:%02d", timeSec / 3600, (timeSec % 3600) / 60)
}

const (
	secPerHour = 3600
	secPerMin = 60
	startTime = secPerHour * 9
)

func TestGetBusTimes(t *testing.T) {
	num := 3
	times := 10
	want := []int{startTime, startTime + (secPerMin * times), startTime + (secPerMin * times * 2)}
	rz := GetBusTimeSecs(num, times)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. want:%v rz:%v\n", want, rz)
	}
}

func GetBusTimeSecs(n, t int) []int {
	rz := make([]int, n)
	for i := 0; i < n; i++ {
		rz[i] = startTime + (secPerMin * t * i)
	}
	return rz
}

func TestConvertToBusTimeSecs(t *testing.T) {
	timetable := []string{"08:00","08:01","08:02"}
	want := []int{secPerHour * 8, secPerHour * 8 + secPerMin, secPerHour * 8 + secPerMin * 2}
	rz := ConvertToBusTimeSecs(timetable)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. want:%v rz:%v\n", want, rz)
	}
}

func ConvertToBusTimeSecs(timetable []string) []int {
	rz := make([]int, len(timetable))
	for i, timeStr := range timetable {
		rz[i] = ConvertToBusTimeSec(timeStr)
	}
	return rz
}

func ConvertToBusTimeSec(timeStr string) int {
	tokens := strings.Split(timeStr, ":")
	hours, _ := strconv.Atoi(tokens[0])
	min, _ := strconv.Atoi(tokens[1])
	return hours * secPerHour + min * secPerMin
}
