# 멈춰! 크롤러

## 현재 진행상황
- main branch에 네이버 헤드라인 크롤러 프로토타입 커밋   
- 지금은 순차적으로 헤드라인 URL 스크랩 -> url마다 뉴스 내용 스크랩 -> 댓글 스크랩 하는중   
- 로컬에서 간단한 테스트 결과 뉴스 자체는 겹쳐서 크롤링 되는 경우는 없음   
- 다만, 간혹 댓글의 경우에는 내용 중복으로 크롤링 되는 경우가 발견됐는데, 각 댓글마다 create_time이 다른걸 보면 애초에 다른 댓글이 맞는것 같기도 함 => 확인 필요   

## todo
1. 로그 개판
2. 에러처리 개판
3. comments.pid 대응?
4. 지금처럼 순차적으로 처리할것인지, 적어도 각 뉴스에 대해선 동시성을 지원할 것인지
5. 계속 크롤링 하면 네이버측에서 서버 IP 차단 가능성 농후. 요청간 대기시간 추가할 예정
