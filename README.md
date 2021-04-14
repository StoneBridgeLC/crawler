# 멈춰! 크롤러

## 현재 진행상황
- main branch에 네이버 헤드라인 크롤러 프로토타입 커밋   
- 지금은 순차적으로 헤드라인 URL 스크랩 -> url마다 뉴스 내용 스크랩 -> 댓글 스크랩 하는중   
- 로컬에서 간단한 테스트 결과 뉴스 자체는 겹쳐서 크롤링 되는 경우는 없음   
- 다만, 간혹 댓글의 경우에는 내용 중복으로 크롤링 되는 경우가 발견됐는데, 각 댓글마다 create_time이 다른걸 보면 애초에 다른 댓글이 맞는것 같기도 함 => 확인 필요
- 한 task에서 뉴스 아티클 스크랩 or 댓글 스크랩하다가 에러 발생하는 경우 전체 프로세스가 중단되지 않도록 에러 처리
- add info logging (zap sugar)
- 빌드 목적에 따른 메인함수 분리
  + /cmd/server/main.go : 일반목적
  + /cmd/lambda/main.go : lambda 빌드
- AWS Lambda 함수로 배포 완료
- lambda 빌드 관련 참고 : https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/golang-package.html#golang-package-windows

## todo
1. comments.pid 대응?
1. 지금처럼 순차적으로 처리할것인지, 적어도 각 뉴스에 대해선 동시성을 지원할 것인지