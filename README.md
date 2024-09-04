# 요구사항 분석
1. 사장님은 시스템에 휴대폰 번호와 비밀번호를 입력하여 회원 가입을 할 수 있습니다. => 회원가입 API 개발
    - 사장님의 휴대폰 번호가 올바르게 입력되었는지 확인해주세요. => 핸드폰 번호 유효성 검사
    - 비밀번호를 안전하게 보관할 수 있는 장치를 만들어주세요 . => 비밀번호 해시 암호화 DB저장
2. 사장님은 회원 가입이후, 등록된 정보로 로그인과 로그아웃을 할 수 있습니다. => 로그인/로그아웃 API 개발
3. 사장님은 로그인 이후 상품 관리와 관련하여 아래의 기능을 활용할 수 있어야합니다. => 상품관리 CRUD
    1. 상품의 필수속성(required)은 다음과 같습니다. => 테이블 스키마
        1. 카테고리
        2. 가격
        3. 원가 
        4. 이름 
        5. 설명
        6. 바코드
        7. 유통기한 
        8. 사이즈 : small or large
    2. 직접 상품을 등록/삭제할 수 있어야합니다. => 상품 등록/삭제 API 개발
    3. 상품의 속성을 부분 수정할 수 있어야합니다. => 상품 수정 API 개발
    4. 등록한 상품의 리스트를 볼 수 있어야합니다. => 상품 조회 API 개발 (이름 검색: like, 초성)
        - cursor based pagination 기반으로, 1page 당 10개의 상품이 보이도록 구현 => 커스텀 Response
    5. 등록한 상품의 상세 내역을 볼 수 있어야합니다. => 상품 상세 API 개발
    6. 상품 이름을 기반으로 검색할 수 있어야합니다. => *4번 포함
        - 이름에 대해서 like 검색과 초성검색을 지원해야 합니다.
            - ex) 슈크림 라떼 → 검색가능한 키워드 : 슈크림, 크림, 라떼, ㅅㅋㄹ, ㄹㄸ
4. 로그인하지 않은 유저의 상품 관련 API에 대한 접근 제한 처리가 되어야 합니다. => 상품 관리 API token 검증 (token 만료시간 10분)

# 구현
언어에 상관없이 Docker를 기반으로 서버를 실행 할 수 있도록 작성해주세요.
    => Dockerfile이 포함되어 있으며 build.sh -> docker-compose up 명령어를 통해서 실행합니다. (아래 자세히)

DB 관련 테이블에 대한 DDL 파일을 소스 디렉토리 안에 넣어주세요.
    => services/seeds/ 해당 경로에 있습니다.

테스트 케이스를 작성해주세요.
    => 테스트 코드 작성 및 테스트를 완료했습니다.
    => 테스트 진행 시 테스트 상세 내용을 읽어주세요. (아래 자세히)

JWT토큰을 발행해서 인증을 제어하는 방식으로 구현해주세요
    => 로그인 시 발행되는 jwt 토큰을 상품관리 API 호출 할때 Header=Authorization 값에 포함하여 요청해야 합니다. (token 유효시간 10분)
    => 로그아웃 시 jwt 토큰은 즉시 만료됩니다.

각 API는 아래의 custom response json 형식으로 반환되야합니다.(204 No Content 제외)
    => common.go 파일에 정해진 형식의 struct를 만들고 해당 객체로 반환 하도록 하였습니다.


# 프로젝트 구성
common: 서비스에서 공통적으로 사용하는 코드를 넣었습니다.
config: 서비스에서 사용하는 환경변수 및 DB 커넥션 로직을 넣었습니다.
services: 실제 비즈니스 로직이 들어있습니다.
    api/v1: API 및 test코드가 들어있고 크게 api/logic/query/types로 분리했습니다. (docs는 시간 여건상 작성하지 않았습니다.)
    seeds: 서비스에서 사용되는 테이블 스키마 정보가 있으며, (스키마/데이터 초기화는 시간 여건상 구현하지 않았습니다.)

# 서비스 구동방법 (Window 11 환경에서 개발)
0. [사전 준비사항]: git bash, docker 등 설치 완료된 상태
1. 적당한 디렉토리를 만들고 코드를 받아주세요. (ph-cafe-manager)
    => git clone https://github.com/nsg3355/ph-cafe-manager.git
2. 프로젝트 디렉토리로 이동하세요.
    => cd ph-cafe-manager
3. build.sh 파일을 실행시켜 주세요. (docker build)
    => docker images 명령어로 이미지 생성된 것을 확인
4. docker-compose.yml 파일을 실행시켜 주세요. (docker-compose)
    => docker-compose up

# 개발서버 실행 방법
1. 적당한 디렉토리를 만들고 코드를 받아주세요. (ph-cafe-manager)
    => git clone https://github.com/nsg3355/ph-cafe-manager.git
2. go module을 설치하세요.
    => go mod tidy
    => go mod vendor
    => go build
3. run-mysql.sh 파일을 실행해 주세요. (로컬 mysql 실행)
    => 볼륨마운트 부분은 개인 PC에 적당한 디렉토리를 걸어주세요. (제거해도 상관없음, 데이터 보존X)
4. VScode에서 F5로 디버깅 모드로 코드를 실행하세요.


# 테스트 방법 (개발: payhere, 테스트: payhere_test)
1. 테스트를 위해서 테스트용 DB를 구성하세요.
    => init.go에 있는 "CREATE DATABASE payhere_test;" 쿼리를 DB툴에서 실행 (DBeaver)
    => services/seeds/ddl 폴더 아래 테이블 생성쿼리 3건을 실행해 주세요.
3. user(사용자관리)는 TestPostSignup -> TestPostLogin -> TestPostLogout 순서로 진행해 주세요.
    => 회원가입 -> 로그인 -> 로그아웃
4. product(상품관리)는 TestPostItem -> TestPutItem -> TestGetList -> TestGetByid -> TestDeleteItem 순서로 진행해 주세요.
    => 상품등록 -> 상품수정 -> 상품목록 -> 상품상세 -> 상품삭제