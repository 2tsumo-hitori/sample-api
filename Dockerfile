# Go 언어 이미지 사용
FROM golang:alpine

# 작업 디렉터리 설정
WORKDIR /app

# 파이썬 설치
RUN apk update && apk add python3

# 현재 디렉터리의 모든 파일을 컨테이너 내 /app 디렉터리로 복사
COPY . .

# Go 모듈을 초기화하고 필요한 종속성을 설치
RUN go mod init github.com/2tsumo-hitori/sample-api
RUN go mod tidy
# 애플리케이션 빌드
RUN go build -o main main.go

# 컨테이너 시작 시 실행될 명령 설정
CMD ["./main"]
