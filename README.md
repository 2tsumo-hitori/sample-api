# sample-api

local에서 실행하는게 아닌, wsl2 환경 DockerFile로 이미지 생성 시 util/HangulUtil.go 파이썬 경로 주의

파이썬 설치경로는 /usr/bin/python3임.

마찬가지로 HangulUtil.go에서 CombineSplitWords 함수에서, output을 그대로 string으로 변환 후 반환해야함.

# 시연

![image](https://github.com/backend-study-space/search-api/assets/96719735/1363d6cf-81e6-493b-af64-96a5c4a37116)
![image](https://github.com/backend-study-space/search-api/assets/96719735/1508b1f7-193c-4579-995d-09a913ccc6ca)
