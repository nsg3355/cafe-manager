# 요구사항 분석
1. 사장님은 시스템에 휴대폰 번호와 비밀번호를 입력하여 회원 가입을 할 수 있습니다. => 회원가입 API 개발 <br>
    - 사장님의 휴대폰 번호가 올바르게 입력되었는지 확인해주세요. => 핸드폰 번호 유효성 검사 <br>
    - 비밀번호를 안전하게 보관할 수 있는 장치를 만들어주세요 . => 비밀번호 해시 암호화 DB저장 <br>
2. 사장님은 회원 가입이후, 등록된 정보로 로그인과 로그아웃을 할 수 있습니다. => 로그인/로그아웃 API 개발 <br>
3. 사장님은 로그인 이후 상품 관리와 관련하여 아래의 기능을 활용할 수 있어야합니다. => 상품관리 CRUD <br>
    1. 상품의 필수속성(required)은 다음과 같습니다. => 테이블 스키마 <br>
        1. 카테고리 <br>
        2. 가격 <br>
        3. 원가 <br>
        4. 이름 <br>
        5. 설명 <br>
        6. 바코드 <br>
        7. 유통기한  <br>
        8. 사이즈 : small or large <br>
    2. 직접 상품을 등록/삭제할 수 있어야합니다. => 상품 등록/삭제 API 개발 <br>
    3. 상품의 속성을 부분 수정할 수 있어야합니다. => 상품 수정 API 개발 <br>
    4. 등록한 상품의 리스트를 볼 수 있어야합니다. => 상품 조회 API 개발 (이름 검색: like, 초성) <br>
        - cursor based pagination 기반으로, 1page 당 10개의 상품이 보이도록 구현 => 커스텀 Response <br>
    5. 등록한 상품의 상세 내역을 볼 수 있어야합니다. => 상품 상세 API 개발 <br>
    6. 상품 이름을 기반으로 검색할 수 있어야합니다. => *4번 포함 <br>
        - 이름에 대해서 like 검색과 초성검색을 지원해야 합니다. <br>
            - ex) 슈크림 라떼 → 검색가능한 키워드 : 슈크림, 크림, 라떼, ㅅㅋㄹ, ㄹㄸ <br>
4. 로그인하지 않은 유저의 상품 관련 API에 대한 접근 제한 처리가 되어야 합니다. => 상품 관리 API token 검증 (token 만료시간 10분) <br>

# 구현
언어에 상관없이 Docker를 기반으로 서버를 실행 할 수 있도록 작성해주세요. <br>
    => Dockerfile이 포함되어 있으며 build.sh -> docker-compose up 명령어를 통해서 실행합니다. (아래 자세히) <br>
<br>
DB 관련 테이블에 대한 DDL 파일을 소스 디렉토리 안에 넣어주세요. <br>
    => services/seeds/ 해당 경로에 있습니다. <br>
<br>
테스트 케이스를 작성해주세요. <br>
    => 테스트 코드 작성 및 테스트를 완료했습니다. <br>
    => 테스트 진행 시 테스트 상세 내용을 읽어주세요. (아래 자세히) <br>
<br>
JWT토큰을 발행해서 인증을 제어하는 방식으로 구현해주세요 <br>
    => 로그인 시 발행되는 jwt 토큰을 상품관리 API 호출 할때 Header=Authorization 값에 포함하여 요청해야 합니다. (token 유효시간 10분) <br>
    => 로그아웃 시 jwt 토큰은 즉시 만료됩니다. <br>
<br>
각 API는 아래의 custom response json 형식으로 반환되야합니다.(204 No Content 제외) <br>
    => common.go 파일에 정해진 형식의 struct를 만들고 해당 객체로 반환 하도록 하였습니다. <br>


# 프로젝트 구성
common: 서비스에서 공통적으로 사용하는 코드를 넣었습니다. <br>
config: 서비스에서 사용하는 환경변수 및 DB 커넥션 로직을 넣었습니다. <br>
services: 실제 비즈니스 로직이 들어있습니다. <br>
    api/v1: API 및 test코드가 들어있고 크게 api/logic/query/types로 분리했습니다. (docs는 시간 여건상 작성하지 않았습니다.) <br>
    seeds: 서비스에서 사용되는 테이블 스키마 정보가 있습니다. <br>

# 서비스 구동방법 (Window 11 환경에서 개발)
0. [사전 준비사항]: git bash, docker 등 설치 완료된 상태 <br>
1. 적당한 디렉토리를 만들고 코드를 받아주세요. (ph-cafe-manager) <br>
    => git clone https://github.com/nsg3355/ph-cafe-manager.git <br>
2. 프로젝트 디렉토리로 이동하세요. <br>
    => cd ph-cafe-manager <br>
3. build.sh 파일을 실행시켜 주세요. (docker build) <br>
    => docker images 명령어로 이미지 생성된 것을 확인 <br>
4. docker-compose.yml 파일을 실행시켜 주세요. (docker-compose) <br>
    => docker-compose up <br>

# 개발서버 실행 방법
1. 적당한 디렉토리를 만들고 코드를 받아주세요. (ph-cafe-manager) <br>
    => git clone https://github.com/nsg3355/ph-cafe-manager.git <br>
2. go module을 설치하세요. <br>
    => go mod tidy <br>
    => go mod vendor <br>
    => go build <br>
3. run-mysql.sh 파일을 실행해 주세요. (로컬 mysql 실행) <br>
    => 볼륨마운트 부분은 개인 PC에 적당한 디렉토리를 걸어주세요. (제거해도 상관없음, 데이터 보존X) <br>
4. VScode에서 F5로 디버깅 모드로 코드를 실행하세요. <br>


# 테스트 방법 (개발: payhere, 테스트: payhere_test)
1. 테스트를 위해서 테스트용 DB를 구성하세요. <br>
    => init.go에 있는 "CREATE DATABASE payhere_test;" 쿼리를 DB툴에서 실행 (DBeaver) <br>
    => services/seeds/ddl 폴더 아래 테이블 생성쿼리 3건을 실행해 주세요. <br>
3. user(사용자관리)는 TestPostSignup -> TestPostLogin -> TestPostLogout 순서로 진행해 주세요. <br>
    => 회원가입 -> 로그인 -> 로그아웃 <br>
4. product(상품관리)는 TestPostItem -> TestPutItem -> TestGetList -> TestGetByid -> TestDeleteItem 순서로 진행해 주세요. <br>
    => 상품등록 -> 상품수정 -> 상품목록 -> 상품상세 -> 상품삭제 <br>


# API 요청 예시

**회원가입**
POST | http://localhost:8085/cafe-mgr/api/v1/user/signup
- body
{
    "phone_number": "01023051738",
    "password": "p@ssw0rd"
}
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": null
}

**로그인**
POST | http://localhost:8085/cafe-mgr/api/v1/user/login
- body
{
    "phone_number": "01023051738",
    "password": "p@ssw0rd"
}
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwMTAyMzA1MTczOCIsImV4cCI6MTcyNTQ1NjgyMn0.WvZqmpKJKRMS2Y-JB4mLoramEEDlYRI8sw8GsQTSI5s"
}

**로그아웃**
POST | http://localhost:8085/cafe-mgr/api/v1/user/logout
- body
{
    "user_id": "1"
}
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": null
}

**상품목록**
GET | http://localhost:8085/cafe-mgr/api/v1/product/list?product_id=2&keyword=ㅇㅁㄹㅇ
- header
Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwMTAyMzA1MTczOCIsImV4cCI6MTcyNTQ1NjgyMn0.WvZqmpKJKRMS2Y-JB4mLoramEEDlYRI8sw8GsQTSI5s
- query param
product_id:2
keyword:ㅇㅁㄹㅇ
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": {
        "products": [
            {
                "id": 2,
                "category": "라떼",
                "price": "4500.00",
                "name": "아메리카노",
                "size": "small"
            }
        ]
    }
}
**상품상세**
GET | http://localhost:8085/cafe-mgr/api/v1/product/byid?product_id=2
- header
Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwMTAyMzA1MTczOCIsImV4cCI6MTcyNTQ1NjgyMn0.WvZqmpKJKRMS2Y-JB4mLoramEEDlYRI8sw8GsQTSI5s
- query param
product_id:2
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": {
        "id": 2,
        "user_id": "1",
        "category": "라떼",
        "price": "4500.00",
        "cost": "450.50",
        "name": "아메리카노",
        "description": "스벅원두",
        "barcode": "455444992",
        "expiration_date": "2025-09-12",
        "size": "small",
        "created_at": "2024-09-04 13:24:41",
        "updated_at": "2024-09-04 13:24:41"
    }
}

**상품등록**
POST | http://localhost:8085/cafe-mgr/api/v1/product/item
- header
Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwMTAyMzA1MTczOCIsImV4cCI6MTcyNTQ1NjgyMn0.WvZqmpKJKRMS2Y-JB4mLoramEEDlYRI8sw8GsQTSI5s
- body
{
    "user_id": 1,
    "category": "라떼",
    "price": "4500",
    "cost": "450.50",
    "name": "아메리카노",
    "description": "스벅원두",
    "barcode": "455444992",
    "expiration_date": "2025-09-12",
    "size": "small"
}
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": null
}

**상품수정**
PUT | http://localhost:8085/cafe-mgr/api/v1/product/item
- header
Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwMTAyMzA1MTczOCIsImV4cCI6MTcyNTQ1NjgyMn0.WvZqmpKJKRMS2Y-JB4mLoramEEDlYRI8sw8GsQTSI5s
- body
{
    "product_id": 2,
    "category": "라떼",
    "price": "4600",
    "cost": "460.50",
    "name": "바닐라라떼",
    "description": "서울우유",
    "barcode": "444444991",
    "expiration_date": "2025-09-30",
    "size": "large"
}
- response
{
    "meta": {
        "code": 200,
        "message": "ok"
    },
    "data": 1
}

**상품삭제**
`
DELETE | http://localhost:8085/cafe-mgr/api/v1/product/item
- header
Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZV9udW1iZXIiOiIwMTAyMzA1MTczOCIsImV4cCI6MTcyNTQ1NjgyMn0.WvZqmpKJKRMS2Y-JB4mLoramEEDlYRI8sw8GsQTSI5s
- body
{
    "product_id": 2
}
- response
{
    "meta": {
        "code": 400,
        "message": "token is expired or invalid"
    },
    "data": null
}
`