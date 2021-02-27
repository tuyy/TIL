package example1

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestHappy3(t *testing.T) {
	cases := []struct{
		dartResult string
		want int
	}{
		{dartResult: "1S2D*3T", want: 37},
		{dartResult: "1D2S0T", want: 3},
		{dartResult: "1D2S#10S", want: 9},
		{dartResult: "1S*2T*3S", want: 23},
		{dartResult: "1D#2S*3S", want: 5},
		{dartResult: "1D2S3T*", want: 59},
		{dartResult: "1T2D3D#", want: -4},
	}

	for _, c := range cases {
		rz := calcDartPoint(c.dartResult)
		if c.want != rz {
			t.Errorf("invalid result. want:%d rz:%d\n", c.want, rz)
		}
	}
}

func calcDartPoint(dartResult string) int {
	points := [3]int{}
	i := 0
	numStr := ""
	for _, ch := range dartResult {
		switch ch {
		case 'S':
			points[i] = calcByNumStr(&numStr, 1)
			i++
		case 'D':
			points[i] = calcByNumStr(&numStr, 2)
			i++
		case 'T':
			points[i] = calcByNumStr(&numStr, 3)
			i++
		case '*':
			if i == 1 {
				points[i - 1] *= 2
			} else {
				points[i - 1] *= 2
				points[i - 2] *= 2
			}
		case '#':
			points[i - 1] *= -1
		default:
			numStr += string(ch)
		}
	}

	return addPoints(points)
}

func addPoints(points [3]int) int {
	sum := 0
	for _, point := range points {
		sum += point
	}
	return sum
}

func calcByNumStr(numStr *string, squareNum int) int {
	rz, err := strconv.Atoi(*numStr)
	if err != nil {
		panic("invalid numstr" + *numStr)
	}

	rz = int(math.Pow(float64(rz), float64(squareNum)))
	*numStr = ""
	return rz
}

func TestHappy4(t *testing.T) {
	str := "1B3S4H"
	for _, ch := range str {
		switch ch {
		case '1':
			fmt.Println(ch)
		case 'B':
			fmt.Println(ch)
		case '3':
			fmt.Println(ch)
		default:
			fmt.Println("d", ch)
		}
	}
}
