package kakao2018

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func makeDecodedMap(n int, arr1 []int, arr2 []int) []string {
	rz := make([]string, n)
	for i := 0; i < n; i++ {
		rz[i] = decodeBinStr(makeBinStr(n, arr1[i]|arr2[i]))
	}
	return rz
}

func TestHappy(t *testing.T) {
	caces := []struct {
		n    int
		arr1 []int
		arr2 []int
		want []string
	}{
		{5, []int{9, 20, 28, 18, 11}, []int{30, 1, 21, 17, 28},
			[]string{"#####", "# # #", "### #", "#  ##", "#####"}},
		{6, []int{46, 33, 33, 22, 31, 50}, []int{27, 56, 19, 14, 14, 10},
			[]string{"######", "###  #", "##  ##", " #### ", " #####", "### # "}},
	}

	for _, c := range caces {
		rz := makeDecodedMap(c.n, c.arr1, c.arr2)
		if !reflect.DeepEqual(c.want, rz) {
			t.Errorf("invalid result. \nex:%s\nrz:%s", c.want, rz)
		}
	}
}

func makeBinStr(n, num int) string {
	rz := strconv.FormatInt(int64(num), 2)
	size := n - len(rz)
	for i := 0; i < size; i++ {
		rz = "0" + rz
	}
	return rz
}

func TestBinaryNum(t *testing.T) {
	cases := []struct {
		num  int
		want string
	}{
		{num: 9, want: "01001"},
		{num: 28, want: "11100"},
		{num: 1, want: "00001"},
		{num: 4, want: "00100"},
	}

	for _, c := range cases {
		got := makeBinStr(5, c.num)
		if got != c.want {
			t.Errorf("invalid binstr:%s(%s) \n", got, c.want)
		}
	}
}

func TestReplaceStr(t *testing.T) {
	str := "01001"
	expected := " #  #"
	rz := decodeBinStr(str)
	if expected != rz {
		t.Error("failed to replace str")
	}
}

func decodeBinStr(str string) string {
	rz := strings.ReplaceAll(str, "0", " ")
	rz = strings.ReplaceAll(rz, "1", "#")
	return rz
}

func TestHappy2(t *testing.T) {
	cases := []struct {
		n    int
		arr1 []int
		arr2 []int
		want []string
	}{
		{
			n: 5,
			arr1: []int{9, 20, 28, 18, 11},
			arr2: []int{30, 1, 21, 17, 28},
			want: []string{"#####", "# # #", "### #", "#  ##", "#####"},
		},
		{
			n: 6,
			arr1: []int{46, 33, 33 ,22, 31, 50},
			arr2: []int{27 ,56, 19, 14, 14, 10},
			want: []string{"######", "###  #", "##  ##", " #### ", " #####", "### # "},
		},
	}

	for _, c := range cases {
		rz := makeSecretMap(c.n, c.arr1, c.arr2)
		if !reflect.DeepEqual(c.want, rz) {
			t.Errorf("invalid result.\nwant:%s\n  rz:%s\n", c.want, rz)
		}
	}
}

func makeSecretMap(n int, arr1 []int, arr2 []int) []string {
	rz := make([]string, n)
	for i := 0; i < n; i++ {
		format := "%0" + strconv.Itoa(n) + "b"
		row := fmt.Sprintf(format, arr1[i] | arr2[i])
		rz[i] = strings.ReplaceAll(strings.ReplaceAll(row, "0", " "), "1", "#")
	}
	return rz
}
