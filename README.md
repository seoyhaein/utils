# utils

### 종속성 문제

- jsoniter "github.com/json-iterator/go" 이 녀석 및 관련 라이브러리는 윈도우에서는 안됨. 문제해결 못해서 우분투로 넘어옴.

## git 사용법 계속 잊어버려서
```bash

# git tag 확인
git tag

# tag 생성, -m 삭제 가능.
git tag -a v1.0.0 -m "Release version 1.0.0"

# 그후 github 에 업데이트 할때 아래 처럼 테그를 명시해줘야함.
git push origin v1.0.0

```
 