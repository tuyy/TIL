### HTTP/1.1 동작방식
* 기본적으로 Connection 당 하나의 요청을 처리하도록 설계되었다. 그래서 동시 전송은 불가능하고 요청과 응답이 순차적으로 이루어진다.

### HTTP/1.1 단점
* 매 요청시 중복된 헤더값을 전송하므로 불필요한 메타정보를 매번 전송한다.
* 하나의 Connection 당 하나의 요청을 처리하므로 매 요청마다 불필요한 TCP 핸드쉐이킹이 반복적으로 일어나므로 네트워크 성능 저하가 발생한다.
* 순차적으로 요청과 응답이 이루어지므로 다수의 요청을 처리하기 위해 pipelining이 존재한다. 이는 요청 순서에따라 대기시간이 길어진다.

#### 이를 해결하기 위한 노력?
* Image spriting : 이미지 요청 최소화 노력
* Domain sharding : 다수의 Connection 요청으로 병렬처리
* Minify javascript/css : 공백을 없애는 노력
* Data URL scheme : 이미지를 코드에 넣어 요청 수 최소화
* Load Faster : CSS를 html 상위에 배치, script를 하단에 배치

#### 구글의 SPDY 프로토콜
* 처리량 관점이 아닌 Latency 관점에서 HTTP를 고속화한 새로운 프로토콜
* HTTP/2 초안의 참고 규격이 된다.

### HTTP/2
* 새로운 프로토콜이라기 보단 기존 HTTP에서 성능 향상에 초점을 맞춘 프로토콜

#### HTTP/2 주요 장점
* Multiplexed Streams : 하나의 Connection으로 동시에 여러개의 메시지를 주고 받을 수 있으며 응답은 순서에 상관없이 stream으로 주고 받는다.
* Stream Prioritization : 요청 리소스 간 우선순위를 설정할 수 있다.
* Server Push : 서버는 클라이언트 요청 없이 리소스를 마음대로 보내줄 수 있다.
* Header Compression : Header를 압축하여 사용하고 중복 제거

### HTTP/1.1과 HTTP/2의 성능 비교
* 물론 네트워크 환경에 따라 다르겠지만 최소 15% ~ 50% 이상의 성능 향상을 꾀할 수 있다고 한다.

### 앞으로는?
* HTTP/2는 모바일 시대에 적합한 것 같다. 현재 모바일 브라우저는 HTTP/2를 대부분 지원한다.
