## JAVA 8 FEATURES

### Lambda Expression
```java
List<String> list = Arrays.asList("first", "second", "333");
list.forEach(v -> System.out.println(v));

first
second
333
```

### Method References
```java
List<String> list = Arrays.asList("first", "second", "333");
list.forEach(System.out::println);

first
second
333
```

### Default Method
```java
public interface vehicle {

   default void print() {
      System.out.println("I am a vehicle!");
   }
}
```

### Streams
```java
List<String> list = Arrays.asList("first", "second", "666", "333", "444", "555");
List<String> result = list.stream()
                          .limit(10)
                          .filter(str -> !str.equals("first"))
                          .map(str -> str + 3)
                          .sorted()
                          .collect(Collectors.toList());
result.forEach(System.out::println);

String result2 = list.stream()
                     .map(str -> str + 1)
                     .collect(Collectors.joining(", "));
System.out.println("JOINING:" + result2);
```

### 기타
 Optional, LocalDate, parallel array sorting, StringJoiner 등 추가되었다.
이미 다른 언어를 통해 사용해 본 feature들이며, 자세한 문법은 필요할때 학습하면서 정리해나가도록 한다.
