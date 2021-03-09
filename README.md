# go_basic
- Golang 기본기 학습 및 컴퓨터지식 기초 공부
- 기반youtube: [Tucker Programming](https://www.youtube.com/channel/UCZp_ftx6UB_32VfVmlS3o_A)

### 고수준언어와 저수준언어와 기계어
- 고수준언어(통상적인 프로그래밍 언어) -> 저수준언어(어셈블리어) <--> 기계어
- 저수준언어와 기계어 사이에는 양방향 변환이 가능하지만, 고수준언어와 어셈블리어는 단절되어 있다.

### 컴파일언어 vs 동적언어
- 컴파일언어
  - 종류 : c, c++, golang 등...
  - 컴파일 : 코딩 -> 빌드 -> 실행파일(기계어)
  - 특징 
    - 빌드를 하게되면 코딩된 부분들이 모두 기계어로 변환되기 때문에 속도가 매우 빠르다.
    - 엔진같은 경우 컴파일언어로 많이 만듬
  - 문제점 
    - 플랫폼 별 다른 변환 필요(실행파일이 모두 달랐음, window/mac/linux 등등..)
    - 경우의 수가 너무 많았다(각 브랜드별 cpu칩셋에 운영체재까지 통일이 안되어있어서 조합이 너무 다양했음)
  
- 동적언어
  - 종류 : java, c# / python, javascript
  - 컴파일 :(java, c#) 코딩 -> 빌드 -> 중간상태 -> 실행 시 필요한 부분만 그때그때 기계어로 변환
  - 특징
    - 그때그때 필요한 부분들을 기계어로 변환하여 해석하기때문에 상대적으로 컴파일언어보다는 속도가 느림.
    - 프로그램을 실행할 때, 플랫폼을 파악해서 그거에 맞게 기계어로 만들어서 너무 효율적이게 됐음.
  - 문제점
    - 초기 > 속도가 떨어졌다
    - low한 프로그래밍을 하기엔 제한이 있음

## GoLang 시작
- 목적 : 시스템 프로그래밍을 위해 탄생한 언어
- 특징 : 기본 utf-8을 지원함.

### package main
- main 은 시작점으로의 약속이다.

### 변수의 속성
- 메모리 주소 + 사이즈 = 끝점
- 이름
- 값
- 메모리 주소
- 사이즈

### int : 4/8byte --> 컴퓨터 사양이 4/8 인지에 따라 기본형이 정해짐
- int32 : 4byte > 0 ~ 42억
- uint32 : 4byte > -21억 ~ 21억
- int64 : 8byte
- int8 : 1byte > -128 ~ 127
- uint8 : 1byte > 0 ~ 255
- int16 : 2byte > -32768 ~ 32767
- uint16 : 2byte > 0 ~ 65535 

### float(정밀도가 관건)
- float32 : 4byte > 숫자부분 7개
- float64 : 8byte > 숫자부분 15개

### bool(boolean) 
- 사이즈 x

### String
- 문자열
- 사이즈 x
- 사이즈는 값에 결정됨


### 연산자(operator)
- link : golang.org/ref/spec#Operator
- 산술 
  - +, -, x, /
- 비트
  - &(and), |(or), ^(xor)
  - ^ 의 경우 이항연산자일 때는 xor, 단항연산자일때는 not
- 논리

### Garbage Collector(GC)
- 메모리에 쌓인 쓰레기(메모리)를 청소한다
- c언어에서는 malloc을 이용하게될 경우 memory leak을 경험할 수 있다. (가령 if문 안에 변수를 선언하고 그 변수에 malloc 을 할당했는데, if 문을 빠져나가게 되면 c언어 특성상 변수가 삭제된다.( if 뿐만 아니라 {} scope를 빠져나가게되면 삭제되게끔 정책이 되어 있다. ) 그렇게 되면 malloc 은 살아있게 되고 접근할 수 없는 garbage 로 남게된다.(malloc 과 free 는 반드시 한세트로 사용해야한다!!)
- reference count 횟수
  - 이 횟수를 통해 garbage 인지 여부를 알 수 있음
  - var a int
  - var p int
  - a = *p
  - 이런 경우엔 a 자체도 값을 가지고 있고, p 또한 a 를 가리키므로
  - a 의 rc 는 2가 된다.
  - 여기서 만약 a 가 삭제되고 p 는 살아있게 된다면 이를 '댕글링' 이라고 한다.(실체없이 댕글댕글 매달려 있다고 해서...) 이런 경우엔 access violation 으로 프로그램이 강제 종료됨.
  - 횟수가 0이 되면 쓰레기로 간주할 수 있다.

- c언어 vs Golang
  - go에선 stack, heap 메모리의 구분이 없음.(c는 있다)
  - gc가 있음(golang), c는 gc가 없음
- gc 는 멀티스레드로 작업을 한다.
  - 초창기에는 프로그램을 멈춰두고 gc가 작업을 하다보니 굉장히 속도에 제한사항이 있었고, 때문에 java가 욕을 많이 먹었음.
- GC가 있는 언어 : Managed 언어 / 없는 언어 : Unmnaged 언어

### slice vs linkedlist
- add
  - slice 는 capacity 가 넘친다는 가정하에 데이터를 추가할 경우 O(N)
  - linkedList는 O(1)
- remove
  - slice : 맨앞/맨끝일경우 O(1), 중간일경우 O(N)
    - 맨앞인 경우는 시작점pointer만 바꿔주면되기때문에 O(1)이다
  - linkedlist : doubleLinkedList 인 경우 O(1), single은 O(N)
- randomAccess(내부 구성요소 찾기)
  - slice : O(1) / index 를 직접 기입해서 접근하면 되기때문
  - linkedList : O(N)
- 배열(slice)는 연속된 메모리, linkedList 는 산재된 메모리
  - cpu가 캐시에 데이터를 가져갈때 배열은 매우 유리함
  - linkedList 는 '캐시미스' 가 계속 날 확률이 높음
