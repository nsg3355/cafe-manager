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

# 프로젝트 구성

# 테이블 설계
CREATE TABLE user_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    category VARCHAR(50) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    cost DECIMAL(10,2) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    barcode VARCHAR(50) UNIQUE NOT NULL,
    expiration_date DATE NOT NULL,
    size ENUM('small', 'large') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_info(id) ON DELETE CASCADE
);

CREATE TABLE access_control (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    access_token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_info(id) ON DELETE CASCADE
);