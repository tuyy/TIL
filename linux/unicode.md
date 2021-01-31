## 유니코드(Unicode)
* 전세계 모든 문자들을 컴퓨터로 표현할 수 있도록 설계된 표준이 바로 유니코드(Unicode)이며 유니코드는 모든 문자들에 고유의 값을 부여한다.
* 즉 코드와 문자가 1:1로 매핑된 형태이다. 단순 코드표이다.
* 전세계 약 14만개의 문자를 표현한다.
* 최대 4바이트로 표현된다.

## 인코딩
* 모든 문자들을 유니코드 4바이트씩 지정해서 표현하는 것은 매우 비효율적입니다. 예를 들어 전체 텍스트가 모두 영어라면 어차피 영문자는 값의 범위가 0 부터 127 사이 이므로 1 바이트만 있어도 되기 때문이다.
* 그래서 등장한 것이 바로 인코딩 방식이며 인코딩 방식에 따라 컴퓨터에서 문자를 표현하기 위해 동일하게 4 바이트를 사용하는 대신에 어떤 문자는 1 바이트, 어떤 건 2 바이트 등의 길이로 저장하게 된다.

```
# 인코딩 방식
UTF-8 : 문자를 최소 1 부터 최대 4 바이트로 표현한다. (문자마다 길이가 다름)
UTF-16 : 문자를 2 혹은 4 바이트로 표현한다.
UTF-32 : 문자를 4 바이트로 표현한다.
```

### UTF-8
* 코드표를 보자 (https://ko.wikipedia.org/wiki/UTF-8)
```C++
#include <string>
#include <iostream>

static void printBinary(char c)
{
    for (int i = 7; i >= 0; --i)
    {
        std::cout << ((c & (1 << i))? '1' : '0');
    }
}

int32_t main(void)
{
    // string은 char(1byte)단위로 저장하므로, UTF-8에서 한글은 3바이트로 표현하기 때문에 size는 9이다.
    for (auto s : std::vector<std::string>{"헬로요"})
    {
        std::cout << "str:" << s << std::endl;
        std::cout << "size:" << s.size() << std::endl;

        for (int i = 0; i < s.size(); i += 3) {
            printBinary(s[i]);
            std::cout << " ";
            printBinary(s[i + 1]);
            std::cout << " ";
            printBinary(s[i + 2]);
            std::cout << std::endl;
        }
        std::cout << std::endl;
    }
    return 0;
}

$ ./main
str:헬로요
size:9
11101101 10010111 10101100
11101011 10100001 10011100
11101100 10011010 10010100
```
