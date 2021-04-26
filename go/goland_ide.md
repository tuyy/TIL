# goland 명령어
 자주 사용하는거 위주로 작성. 안하다가 나중에 한번 보면 바로 손에 익힐 수 있도록 계속 추가하자

```
- cmd + option + L : 기본 줄 맞춤
- cmd + option + +shift + F   : rungo fmt
- option + shift + up or down : 해당 라인 이동
- option + up : 단어 block
    - 여러번 누르면 점점 범위가 커짐
- cmd + / : 주석
- cmd + shift + enter : move to a new line
- ctl + shift + space : smart auto complete


# 코드 자동 완성
- err.nn   => if err != nil { .. }
- err.nil
- {collection 변수}.sort  -> sort.Strings({collection 변수})
- {문자열}. + ctl + space twice


# Debugging
- ctl + shift + D : 디버깅 시작
- cmd + F2 : 디버깅 종료

- cmd + F8 : toggle break point

- option + F9 : Run to cursor -> 기본 이동
- F7 : step into -> 함수 안으로 이동
- shift + F8 : step out -> 함수 밖으로 이동
- F8 : step over
```
