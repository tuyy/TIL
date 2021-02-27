package example2

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestHappy2223(t *testing.T) {
	record := []string{"Enter uid1234 Muzi", "Enter uid4567 Prodo","Leave uid1234","Enter uid1234 Prodo","Change uid4567 Ryan"}
	want := []string{"Prodo님이 들어왔습니다.", "Ryan님이 들어왔습니다.", "Prodo님이 나갔습니다.", "Prodo님이 들어왔습니다."}

	rz := solution(record)
	if !reflect.DeepEqual(rz, want) {
	    t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

type ActAndUuid struct {
	Act, Uuid string
}

func solution(record []string) []string {
	m := make(map[string]string)
	var arr []ActAndUuid
	for _, v := range record {
		tokens := strings.Split(v, " ")
		act := tokens[0]
		uuid := tokens[1]
		if len(tokens) == 3 {
			name := tokens[2]
			m[uuid] = name
		}

		if act == "Enter" || act == "Leave" {
			arr = append(arr, ActAndUuid{act, uuid})
		}
	}

	result := make([]string, len(arr))
	for i, v := range arr {
		switch v.Act {
		case "Enter":
			result[i] = fmt.Sprintf("%s님이 들어왔습니다.", m[v.Uuid])
		case "Leave":
			result[i] = fmt.Sprintf("%s님이 나갔습니다.", m[v.Uuid])
		}
	}

	return result
}
