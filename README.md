# KiwoomService
Kiwoom OpenApi 연결하여 현재 가격을 받아 연결된 Client 에 제공한다.
단순히 최근 체결가만 얻기 위해 간단히 만듬.
개인적인 용도로 만들어 추후 관리는 예정 되어 있지 않음.

.Net Framework 4.7.2 기반이며 자신의 OS 가 담배를 피운다면 Micorosoft 에서 다운 받아 설치하시길 바랍니다.
연결 요청은 Http 이며 python, c#, Go, javascript 로 손가락 아프게 끄적여 놨음.
다른 언어도 별로 어렵지 않아요.

+ **사용**   
+키움 증권 사이트에서 OpenAPI 설치와 사용 신청.
https://www.kiwoom.com/h/customer/download/VOpenApiInfoView

+OpenAPI 로그인 후 KiwoomService.exe 실행

+종목 등록 요청 및 현재가 요청 사용

+계좌 변동 내역 이벤트 요청

+Port 가 사용중일 경우 'KiwoomService.exe -port 다른포트' 명령어로 실행

``` python
try:
    ## 코드 ';' 구분 등록. POST 사용가능.
    srcHndl = urllib.request.urlopen("http://127.0.0.1:9797/RegStockCodes?code=000660;005930")	
    srcByte = srcHndl.read()
    res = json.loads(srcByte.decode("utf8"))
    p(res["Result"])
except Exception as e:
    p(e)
    return

```

## 여담
돌아가지 않는 머리와 손가락 관절의 괴음을 참아 내며 
하루 하루 오병이어로 연명 하며 대충 만듬.
그러니 많은걸 기대 하지 마시길...
