package kakao2018

import (
	"strings"
	"testing"
)

func TestHappy5(t *testing.T) {
	cases := []struct{
		cacheSize int
		cities []string
		want int
	}{
		{
			cacheSize: 3,
			cities: []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA", "Jeju", "Pangyo", "Seoul", "NewYork", "LA"},
			want: 50,
		},
		{
			cacheSize: 3,
			cities: []string{"Jeju", "Pangyo", "Seoul", "Jeju", "Pangyo", "Seoul", "Jeju", "Pangyo", "Seoul"},
			want: 21,
		},
		{
			cacheSize: 2,
			cities: []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA", "SanFrancisco", "Seoul", "Rome", "Paris", "Jeju", "NewYork", "Rome"},
			want: 60,
		},
		{
			cacheSize: 5,
			cities: []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA", "SanFrancisco", "Seoul", "Rome", "Paris", "Jeju", "NewYork", "Rome"},
			want: 52,
		},
		{
			cacheSize: 2,
			cities: []string{"Jeju", "Pangyo", "NewYork", "newyork"},
			want: 16,
		},
		{
			cacheSize: 0,
			cities: []string{"Jeju", "Pangyo", "Seoul", "NewYork", "LA"},
			want: 25,
		},
	}

	for _, c := range cases {
		rz := calcElapsedTime(c.cacheSize, c.cities)
		if c.want != rz {
			t.Errorf("invalid result. want:%d rz:%d\n", c.want, rz)
		}
	}
}

func calcElapsedTime(cacheSize int, cities []string) int {
	cache := NewCache(cacheSize)
	rz := 0
	for _, city := range cities {
		city = strings.ToUpper(city)
		if cache.hasValue(city) {
			rz += 1
		} else {
			rz += 5
			cache.addValue(city)
		}
	}
	return rz
}

type Cache struct {
	Size int
	Arr []string
}

func (c Cache) hasValue(value string) bool {
	for _, v := range c.Arr {
		if v == value {
			return true
		}
	}
	return false
}

func (c *Cache) addValue(value string) {
	if c.hasValue(value) || c.Size == 0{
		return
	}

	if len(c.Arr) < c.Size {
		c.Arr = append(c.Arr, value)
	} else {
		c.Arr = append(c.Arr[1:], value)
	}
}

func TestCache(t *testing.T) {
	cache := NewCache(3)

	if cache.hasValue("TestVal") {
		t.Error("invalid hasValue result")
	}

	cache.addValue("TestVal")
	if !cache.hasValue("TestVal") {
		t.Error("invalid hasValue result")
	}

	cache.addValue("TestVal1")
	cache.addValue("TestVal2")
	cache.addValue("TestVal3")
	cache.addValue("TestVal4")
	if cache.hasValue("TestVal") {
		t.Error("invalid hasValue result")
	}
	if !cache.hasValue("TestVal2") {
		t.Error("invalid hasValue result")
	}
}

func NewCache(size int) *Cache {
	rz := Cache{Size: size}
	return &rz
}
