# 멈춰! 크롤러

## 개요
네이버 헤드라인 뉴스 크롤러

## 배포
1. GitHub에서 build-lambda-zip 도구를 다운로드
```
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

2. 빌드 및 압축파일 생성
```
set GOOS=linux
go build -o main main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -output main.zip main
```

전체 docs는 다음 링크 참고 : https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/golang-package.html#golang-package-windows

## 업데이트 로그
- 2021-04-15
  + 한 task에서 뉴스 아티클 스크랩 or 댓글 스크랩하다가 에러 발생하는 경우 전체 프로세스가 중단되지 않도록 에러 처리
  + add info logging (zap sugar)
  + 빌드 목적에 따른 메인함수 분리
    + /cmd/server/main.go : 일반목적
    + /cmd/lambda/main.go : lambda 빌드  
  + 첫 배포 (AWS Lambda 함수로 배포)
- 2021-05-24
  + 뉴스 크롤링할 때 URL도 DB에 함께 저장함
  + 뉴스 본문의 길이가 200자 이하이면 저장하지 않음