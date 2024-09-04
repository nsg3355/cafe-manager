# 빌드 스테이지
FROM golang:1.20 AS builder

# 작업 디렉토리 설정
WORKDIR /app

# 모듈 파일을 복사하고 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드를 복사
COPY . .

# 프로젝트 빌드 (바이너리 이름: app)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app .

# 실행 스테이지 (Ubuntu 사용)
FROM ubuntu:22.04

# 필수 패키지 설치
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && update-ca-certificates

# 빌드된 바이너리 복사
COPY --from=builder /app/app /app/app

# .env 파일 복사
COPY .env /app/.env
COPY ./services/seeds/ddl /app/services/seeds/ddl

# 실행할 디렉토리 설정
WORKDIR /app

# 실행 가능한 바이너리 설정
CMD ["/app/app"]
