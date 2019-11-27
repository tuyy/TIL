#### javascript 기본타입 (7가지)
* Number
* String
* Boolean
* Object
* Null
* Undefined
* Simbol

#### javascript 연산자
* == 는 값만 비교, === 는 값 + 타입 비교
* NaN 은 isNaN(..) 으로만 비교 가능
* &&, || 굳이 연산의 결과가 미리 정해지면 그 다음 수식은 실행되지 않음
```
var val = window.val || 'hello val';
```

#### typeof
* null 을 주의하자
```
typeof 80 => "number"
typeof "80" => "string"
typeof true => "boolean"
typeof a => "undefined"
typeof functino(){} => "function"
typeof [] => "object"

typeof null => "object"     <- null조심

if (!v && typeof v === 'object') {
    console.log('v is null.');
}
```

### object
```
let obj = {name: 'hello', val: 'world'};
Object.keys(obj);  // 키를 배열로 반환
Object.values(obj);
```
