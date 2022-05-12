# KiwoomService
Kiwoom OpenApi 연결하여 현재 가격을 받아 연결된 Client 에 제공한다.
단순히 최근 체결가만 얻기 위해 간단히 만듬.
개인적인 용도로 만들어 추후 관리는 예정 되어 있지 않음.

.Net Framework 4.7.2 기반이며 자신의 OS 가 담배를 피운다면 Micorosoft 에서 다운 받아 설치하시길 바랍니다.

연결 요청은 Http 이며 python, c#, Go, javascript 로 손가락 아프게 끄적여 놨음.
다른 언어도 별로 어렵지 않아요.

## 사용
+ 키움 증권 사이트에서 OpenAPI 설치와 사용 신청.
https://www.kiwoom.com/h/customer/download/VOpenApiInfoView

+ OpenAPI 로그인 후 KiwoomService.exe 실행

+ 종목 등록 요청 및 현재가 요청 사용

+ 계좌 변동 내역 이벤트 요청

+ Port 가 사용중일 경우 'KiwoomService.exe -port 다른포트' 명령어로 실행

## Python Test Code
일단 멀지도 가깝지도 않은 Python test code.
다른 언어는 위 파일로 확인.

``` python
from logging import raiseExceptions
import urllib.request
import json
import time
from threading import Thread

p = print

def main():
    p("start")

    #########################################
    ### 종목 등록
    try:
        ## 코드 ';' 구분 등록. POST 사용가능.
        srcHndl = urllib.request.urlopen("http://127.0.0.1:9797/RegStockCodes?code=000660;005930")	
        srcByte = srcHndl.read()
        res = json.loads(srcByte.decode("utf8"))
        p(res["Result"])
    except Exception as e:
        p(e)
        return

    #########################################
    ### 계좌 이벤트
    thd = Thread(target=GetEvtAccount)
    thd.start()


    #########################################
    ### 현재가 요청
    while True:
        data,err = GetRealPrice()
        if err != None :
            p(err)
            time.sleep(1)
            continue

        # p(data)
        for code, price in data.items() :
            p(code, price)

        time.sleep(2)
    pass


def GetEvtAccount() :
    while True:
        try:
            srcHndl = urllib.request.urlopen("http://127.0.0.1:9797/GetEvtAccount")	
            srcByte = srcHndl.read()
            res = json.loads(srcByte.decode("utf8"))
            # evt.Data.Code  는 6자리가 넘을경우 뒤 6자리 사용.
            # 종목코드 앞에 A는 장내주식, j는 ELW종목, Q는 ETN종목을 의미 라고함.

            p(res["Data"])
        except Exception as e:
            p(e)
            time.sleep(1)
    pass

def GetRealPrice() :
    try:
        srcHndl = urllib.request.urlopen("http://127.0.0.1:9797/GetRealPrice")	
        srcByte = srcHndl.read()
        res = json.loads(srcByte.decode("utf8"))
        # p(res)
        if res["Result"] != 0 :
            raise Exception(res["Message"])

        return res["Data"], None
    except Exception as e:
        # p(e)
        return None, str(e)

    pass

if __name__ == '__main__':	
    main()

```

## 여담
돌아가지 않는 머리와 손가락 관절의 괴음을 참아 내며 
하루 하루 오병이어로 연명 하며 대충 만듬.
그러니 많은걸 기대 하지 마시길...
