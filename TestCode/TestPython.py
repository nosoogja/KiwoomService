from logging import raiseExceptions
import urllib.request
import json
import time
from threading import Thread

'''
오랜만에 python을 하여 문법이 생각이 안난다.
누군가가 뇌를 다림질 하였다. 누구냐. 너.
하여 대~~~충 돌려본다.

PYthon 3.8
'''

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