## 스타일
* 프로그래밍 스타일의 근본원리
    * 코드는 간단명료 해야한다.
    * 직관적인 논리, 자연스러운 표현, 일반적인 언어의 사용, 의미 있는 이름, 깔끔한 형식, 쓸모 있는 주석을 가져야 한다.
    * 일관성도 매우 중요하다. (팀원과 코드 컨벤션을 잘 따르는 것이 중요하다.)
    
    
* 다음 코드의 문제점은?
```C++
// 1. 
if ( (country == SING) || (country == BRANI)) ||
     (country == POL) || (country == ITALY))
{
    // country가 싱가포르거나 브루나이 또는 폴란드라면
    // 통화가 연결된 시간이 아니라 응답 시간이 현재 시간이 된다.
    // 응답 시간을 초기화하고 요일을 설정한다.
    ..
    ..
}

// 2.
#define ONE 1
#define TEN 10
#define TWENTY 20
```

### 1.1 이름
* 우리는 문맥과 유효범위에서 많은 정보를 얻을 수 있으며, 변수의 유효범위가 넓을수록 이름은 더 많은 정보를 전달해야 한다.

#### 전역변수에는 서술적인 이름을, 지역변수에는 짧은 이름을 붙여라
* 프로그래머들은 때로 문맥과 관계없이 긴 변수 이름을 쓰려고 하지만, 이는 잘못된 것이다.
* 코드의 명료성은 간결함을 통해 얻을 수 있는 경우가 많다.
* 이름은 세부적인 규칙의 내용보다는 정해진 규칙을 일관되게 지키는 것이 훨씬 중요하다.
```C++
// 1. 전역변수
int npending = 0; // 전역변수를 선언할 때 간단한 주석을 달아주는 것도 좋다.

// 2. 지역변수 : 짧은 이름을 써도 충분하다.
// 나쁜예
for (theElementIndex = 0; theElementIndex < numberOfElements; theElementIndex++)
    elementArray[theElementIndex] = theElementIndex;

// 좋은예
for (i = 0; i < nelems; i++)
    elem[i] = i;
```

#### 일관성을 지키라
```C++
// 나쁜예
class UserQueue
{
public:
    int32_t noOfItemsInQ, frontOfTheQueue, queueCapacity;
    int noOfUsersInQueue() {...}
};

// 좋은예 : 문맥을 고려한 이름
class UserQueue
{
public:
    int32_t noOfItems, front, capacity
    int nusers() {...}
};
```

#### 함수 이름에는 능동형을 써라
```C++
// 나쁜예
if (checkoctal(c)) ...

// 좋은예
if (isoctal(c)) ...
```

#### 정확한 이름을 써라
* 혼란을 주는 이름은 잘못된 추상화의 예이다.
```C++
// 나쁜예 : 함수명의 결과가 반대다.
bool inTable(T obj)
{
    int j = this.getIndex(obj);
    return (j == nTable);
}
```

#### 연습
* 아래 코드의 명명과 값의 선택에 대해 평가해보라
```C++
#define TRUE 0
#define FALSE 1

if ((ch = getchar()) == EOF)
    not_eof = FALSE;
```

* 아래 함수를 더 좋게 고쳐라
```C++
int smaller(char *s, char *t)
{
    if (strcmp(s, t) < 1)
        return 1;
    else
        return 0;
}
```

* 이 코드를 소리 내어 읽어보라
```C++
if ((falloc(SMRHSHSCRTCH, S_IFEXT|0644. MAXRODDHSH)) < 0) ..
```

### 1.2 표현식과 문장
* 표현식과 문장 또한 이름과 마찬가지로 그 의미가 최대한 명료해야한다.
* 연산 단위가 눈에 잘 띄게 연산자 주변에 공백을 두자. 즉, 읽기 좋게 코드를 작성해야한다.

#### 들여쓰기로 구조를 알아 보기 쉽게 하라
```C++
// 나쁜예
for(n++;n<100;field[n++]='\0');
*i = '\0'; return('\n')

// 개선
for (n++; n < 100; n++)
    field[n] = '\0';
*i = '\0';
return '\n';
```

#### 표현식을 자연스럽게 써라
* 부정을 포함하는 조건식은 언제나 이해하기 어렵다.
```C++
// 나쁜예
if (!(block_id < actblks) || !(block_id >= unblocks)) ...

// 개선
if ((block_id >= actblks) || (block_id < unblocks)) ...
```

#### 괄호를 써서 애매함을 해소하라
* 연산자 우선순위 문제 뿐 아니라 애매하면 괄호를 써주는게 좋다.
```C++
// 나쁜예
leap_year = y % 4 == 0 && y % 100 != 0 || y % 400 == 0;

// 개선
leap_year = ((y%4 == 0) && (y%100 != 0)) || (y%400 == 0);
```

#### 복잡한 표현은 잘게 쪼개라
```C++
// 나쁜예
*x += (*xp=(2*k < (n-m) ? c[k+1] : d[k--]));

// 개선
if (2*k < n-m)
    *xp = c[k+1];
else
    *xp = d[k--];
*x += *xp;
```

#### 명료하게 써라
* 명료성은 간결성과 동일하지 않다.
* 기준은 짧거나 긴 것이 아니라, 얼마나 코드를 이해하기 쉬운지가 되어야한다.
```C++
// 나쁜예
subkey = subkey >> (bitoff - ((bitoff >> 3) << 3));

// 개선
subkey >>= bitoff & 0x7;

// 삼항 연산자는 짧은 식에 적합하다.
max = (a > b) ? a : b;
```

#### 부수효과를 조심하라
* ++,-- 연산자는 값을 반환하는 것 외에도 변수 값 자체를 바꿔버리는 부수효과가 있다.
* scanf 같은 함수도 마찬가지다. 조심! 또 조심하자.

#### 연습
* 아래 코드를 개선하라
```C++
if (!(c=='y' || c == 'Y'))
    return;
    
length = (length < BUFSIZE) ? length : BUFSIZE;

flag = flag ? 0 : 1;

if (val & 1)
    bit = 1;
else
    bit = 0;
```

### 일관성과 관용 표현

#### 들여쓰기와 중괄호 {}를 쓰는 스타일에서는 일관성을 지켜라
#### 일관성을 위해 관용 표현을 사용하라
```C++
// 나쁜예
i = 0;
while (i <= n-1)
    array[i++] = 1.0

// 개선
for (i = 0; i < n; i++)
    array[i] = 1.0;
    
// 나쁜예
do {
    c = getchar();
    putchar(c);
} while (c != EOF);

// 개선 : do-while 은 반드시 한번은 실행되어야할때만 쓰자!
while ((c = getchar()) != EOF)
    putchar(c);
```

#### 다중결정이 필요할 때는 else-if를 사용하라
* 항상 명료한 코드 depth를 고민하자
* 각각의 조건문과 그 조건문에서 유발되는 동작들은 가능한 한 가까이 붙여서 써야 하는 것이 원칙이다.
```C++
// 나쁜예
switch (c) {
case '-':
    sign = -1;
    /* fall through */
case '+':
    c = getchar();
    break;
   
case '.':
    break;
default:
    break;
}

// 개선 : 상황에 맞게 switch와 if-else문을 사용하라
if (c == '-') {
    // ..
} else if (..) {
    // ..
} else {
    // ..
}

```

#### 연습
``` C++
// AS-IS
if (retval != SUCCESS)
{
    return (retval);
}

return SUCCESS;


// TO-BE
return retval;
```

### 1.4 매크로 함수
* 요즘 같은 시대엔 매크로 함수는 득보다 실이 더 많다.

#### 매크로 함수를 멀리하라
// TODO
