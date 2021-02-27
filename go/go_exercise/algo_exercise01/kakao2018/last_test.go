package kakao2018

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestHappy999(t *testing.T) {
	cases := []struct {
		input []string
		want int
	}{
		{[]string{"2016-09-15 01:00:04.001 2.0s", "2016-09-15 01:00:07.000 2s"}, 1},
		{[]string{"2016-09-15 01:00:04.002 2.0s", "2016-09-15 01:00:07.000 2s"}, 2},
		{[]string{"2016-09-15 20:59:57.421 0.351s", "2016-09-15 20:59:58.233 1.181s", "2016-09-15 20:59:58.299 0.8s", "2016-09-15 20:59:58.688 1.041s", "2016-09-15 20:59:59.591 1.412s", "2016-09-15 21:00:00.464 1.466s", "2016-09-15 21:00:00.741 1.581s", "2016-09-15 21:00:00.748 2.31s", "2016-09-15 21:00:00.966 0.381s", "2016-09-15 21:00:02.066 2.62s"}, 7},
	}

	for _, c := range cases {
		rz := solution(c.input)
		if rz != c.want {
			t.Fatalf("invalid result. rz:%v want:%v", rz, c.want)
		}
	}
}

func solution(input []string) int {
	var arr []int
	ranges := make([][]int, len(input))
	for i, s := range input {
		rangeWithMillis := GetRangeWithMillis(s)
		ranges[i] = rangeWithMillis
		arr = append(arr, rangeWithMillis[0], rangeWithMillis[1])
	}
	sort.Ints(arr)

	max := 0
	for i := arr[0]; i < arr[len(arr)-1]; i++ {
		j := i + 999
		if j > arr[len(arr)-1] {
			j = arr[len(arr)-1]
		}

		cnt := 0
		for _, r := range ranges {
			if r[0] <= i && i <= r[1] {
				cnt++
			} else if r[0] <= j && j <= r[1] {
				cnt++
			}
		}

		if max < cnt {
			max = cnt
		}
	}
	return max
}

func TestGetRangeWithMillis(t *testing.T) {
	input := "2016-09-15 01:00:04.001 2s"
	want := []int{3602002, 3604001}

	rz := GetRangeWithMillis(input)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid rz. rz:%v want:%v", rz, want)
	}
}

func GetRangeWithMillis(input string) []int {
	tokens := strings.Split(input, " ")
    endMillis := GetMillis(tokens[1])
    duration, _ := strconv.ParseFloat(tokens[2][:len(tokens[2])-1], 3)
    durationMillis := int(duration * 1000)
	startMilis := endMillis - durationMillis + 1

	return []int{startMilis, endMillis}
}

func TestGetMillis(t *testing.T) {
	input := "2016-09-15 01:00:04.001"
	want := 3604001

	rz := GetMillis(strings.Split(input, " ")[1])
	if rz != want {
		t.Fatalf("invalid rz. rz:%d want:%d", rz, want)
	}
}

func GetMillis(s string) int {
	tokens := strings.Split(s, ":")
	h, _ := strconv.Atoi(tokens[0])
	result := h * 3600 * 1000

	m, _ := strconv.Atoi(tokens[1])
	result += m * 60 * 1000

	tmp := strings.Split(tokens[2], ".")
	sec, _ := strconv.Atoi(tmp[0])
	result += sec * 1000

	millis, _ := strconv.Atoi(tmp[1])
	result += millis

	return result
}
