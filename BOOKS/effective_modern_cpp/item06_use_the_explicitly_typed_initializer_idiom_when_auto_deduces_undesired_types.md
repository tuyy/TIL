## Item 6: Use the explicitly typed initializer idiom when auto deduces undesired types.
* auto가 원치 않은 형식으로 deduction된다면 명시적으로 캐스팅해야한다.
```C++
std::vector<bool> features(const Widget& w);
Widget w;
bool highPriority = features(w)[5];  // w의 우선순위가 높은가?
processWidget(w, highPriority);  // Ok!

// C++에서 bool의 참조는 불가능하므로,
// highPriority는 std::vector<bool>::reference 형식으로 deduction된다.
auto highPriority = features(w)[5];
processWidget(w, highPriority);  // 미정의 행동!!!!!!!

// auto를 사용하려면, 명시적으로 캐스팅!
auto highPriority = static_cast<bool>(features(w)[5]);
```

### 참고
* Effective Modern C++ by Scott Meyers
