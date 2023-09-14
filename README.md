# sample-api

local에서 실행하는게 아닌, wsl2 환경 DockerFile로 이미지 생성 시 util/HangulUtil.go 파이썬 경로 주의

파이썬 설치경로는 /usr/bin/python3임.

마찬가지로 HangulUtil.go에서 CombineSplitWords 함수에서, output을 그대로 string으로 변환 후 반환해야함.
