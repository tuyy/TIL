### vanilla js 기본 개발 환경 세팅 with jest

#### 기본 jest 환경
```javascript
$ cd ${PROJECT_HOME_DIR}
$ npm init
$ npm install jest

=>> abc.test.js 추가

$ npx jest
 PASS  ./abc.test.js
  ✓ first (1 ms)

Test Suites: 1 passed, 1 total
Tests:       1 passed, 1 total
Snapshots:   0 total
Time:        1.343 s
Ran all test suites.
```

###### abc.test.js

```javascript
function assert1(received) {
    if (received !== 1) {
        throw new Error("received는 1이 아니다. received:" + received)
    }
}

describe('Name of the group', () => {
    test("first", () => {
        assert1(0);
    });
    
    test("second", () => {
        expect(0).toBe(1);
    });
});
```

#### 기본 개발 환경 repo
https://github.com/tuyy/dev_init_for_vanillajs
