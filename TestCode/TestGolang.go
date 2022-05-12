package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*///////////////////////////////////////////////////////////
Go 1.18 로 돌렸으나 별로 쓴게 없으니 이전 버전도 될거라 사료됨.
안되면 아마도 전원 내렸다가 다시 키면 될거라 확신은 하지 않음.
///////////////////////////////////////////////////////////*/

var p = fmt.Println

type ResponseData struct {
	Result  int
	Message string
}
type KiwoomPrice struct {
	ResponseData
	Data map[string]int
}
type EvtAccount struct {
	ResponseData
	Data struct {
		Account   string // 계좌번호
		Code      string // 종목 코드
		Name      string // 종목 명
		Price     int    // 현재가
		UnitPrice int    // 평균단가
		Quantity  int    // 보유수량
		SellOrBuy int    // 1:매도, 2:매수
	}
}
type KiwoomName struct {
	ResponseData
	Data map[string]string
}

func main() {
	var err error
	p("start")

	///////////////////////////////////////////
	/// 종목 등록
	/// 코드 ';' 구분 등록. POST 사용가능.
	res, err := http.Get("http://127.0.0.1:9797/RegStockCodes?code=000660;005930")
	if err != nil {
		p(err)
		return
	}
	_ = res // 생략....

	///////////////////////////////////////////
	/// 종목 명 요청
	nameData, err := GetStockName()
	if err == nil {
		for code, name := range nameData {
			p(code, name)
		}
	}

	///////////////////////////////////////////
	/// 계좌 이벤트
	go func() {
		for {
			evt, err := GetEvtAccount()
			if err != nil {
				p(err)
				time.Sleep(time.Second * 1)
				continue
			}
			/// evt.Data.Code  는 6자리가 넘을경우 뒤 6자리 사용.
			/// 종목코드 앞에 A는 장내주식, j는 ELW종목, Q는 ETN종목을 의미 라고함.

			fmt.Printf("%+v", evt)
		}
	}()

	///////////////////////////////////////////
	/// 현재가 요청
	for {
		data, err := GetRealPrice()
		if err != nil {
			p(err)
			time.Sleep(time.Second * 1)
			continue
		}

		for code, price := range data {
			p(code, price)
		}

		time.Sleep(time.Millisecond * 2000)
	}

}

func GetStockName() (map[string]string, error) {
	var err error
	var res *http.Response
	var resByte []byte

	res, err = http.Get("http://127.0.0.1:9797/GetStockName")
	if err != nil {
		return nil, err
	}
	resByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	kn := KiwoomName{}

	err = json.Unmarshal(resByte, &kn)
	if err != nil {
		return nil, err
	}

	if kn.Result != 0 {
		return nil, errors.New(kn.Message)
	}

	return kn.Data, nil

}

func GetEvtAccount() (EvtAccount, error) {
	var err error
	var res *http.Response
	var resByte []byte
	var resEvt EvtAccount

	res, err = http.Get("http://127.0.0.1:9797/GetEvtAccount")
	if err != nil {
		return resEvt, err
	}
	resByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return resEvt, err
	}

	evt := EvtAccount{}
	err = json.Unmarshal(resByte, &evt)
	if err != nil {
		return resEvt, err
	}

	if evt.Result != 0 {
		return resEvt, errors.New(evt.Message)
	}

	return resEvt, nil

}
func GetRealPrice() (map[string]int, error) {
	var err error
	var res *http.Response
	var resByte []byte

	res, err = http.Get("http://127.0.0.1:9797/GetRealPrice")
	if err != nil {
		return nil, err
	}
	resByte, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	kp := KiwoomPrice{}
	err = json.Unmarshal(resByte, &kp)
	if err != nil {
		return nil, err
	}

	if kp.Result != 0 {
		return nil, errors.New(kp.Message)
	}

	return kp.Data, nil
}
