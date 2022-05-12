using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
//using System.Linq;
using System.Net;
//using System.Text;
using System.Threading;
using System.Threading.Tasks;

/*/////////////////////////////////////////////////////////////
 * 
 * Framework 4.7.2 로 확인함.
 * Newtonsoft.Json 는 nuget에서 Package 다운 받아 쓰시길. c# 쓸만한게 없다.
 * 아니면 .NET 6 사용 하면 있을듯함.
 * 대~~~충 고쳐서 쓰면 당신은 챔피온.
 * 
 * *//////////////////////////////////////////////////////////

namespace TestCsharp
{
  

    class Program
    {
        static void Main(string[] args)
        {
            p("start");
            new Program().MainProcess();
            Console.ReadKey();
        }


        public void MainProcess()
        {
            ///////////////////////////////////////////
            /// 종목 등록
            try
            {
                /// 코드 ';' 구분 등록. POST 사용가능.
                var resText = HttpGet("http://127.0.0.1:9797/RegStockCodes?code=000660;005930");
                p(resText);
            }
            catch (Exception err)
            {
                p(err.Message);
                return;
            }


            ///////////////////////////////////////////
            /// 계좌 이벤트
            Task.Run(GetEvtAccount);

            ///////////////////////////////////////////
            /// 현재가 요청
            while (true)
            {
                try
                {
                    var data = GetRealPrice();
                    foreach(var item in data)
                    {
                        p($"code:{item.Key} price:{item.Value}");
                    }
                }
                catch(Exception err)
                {
                    p(err.Message);
                }


                Thread.Sleep(2000);
            }
        }


        public void GetEvtAccount() 
        {
            string resText = "";
            EvtAccount bin = null;

            while (true)
            {
                try
                {
                    resText = HttpGet("http://127.0.0.1:9797/GetEvtAccount");
                    bin = JsonConvert.DeserializeObject<EvtAccount>(resText);
                    if (bin.Result != 0)
                    {
                        throw new Exception($"{bin.Message}");
                    }
                    /// evt.Data.Code  는 6자리가 넘을경우 뒤 6자리 사용.
                    /// 종목코드 앞에 A는 장내주식, j는 ELW종목, Q는 ETN종목을 의미 라고함.

                    p(resText);
                    p($"{bin.Data.Account} {bin.Data.Quantity}");

                }
                catch(Exception err)
                {
                    p(err.Message);
                    Thread.Sleep(1000);
                }
            }

        }

        public Dictionary<string,int> GetRealPrice()
        {
            KiwoomPrice bin = null;

            string resText = "";
            resText = HttpGet("http://127.0.0.1:9797/GetRealPrice");
            bin = JsonConvert.DeserializeObject<KiwoomPrice>(resText);
            if(bin.Result != 0)
            {
                throw new Exception($"{bin.Message}");
            }


            return bin.Data;
        }

        public string HttpGet(string url)
        {
            string resText = "";

            WebRequest request = WebRequest.Create(url);
            var res = request.GetResponse();
            using (Stream dataStream = res.GetResponseStream())
            {
                StreamReader reader = new StreamReader(dataStream);
                resText = reader.ReadToEnd();
            }

            return resText;
        }


        public static void p(string msg)
        {
            Console.WriteLine(msg);
        }
    }

    public class ResponseData
    {
        public int Result { get; set; }
        public string Message { get; set; }
    }
    public class KiwoomPrice : ResponseData
    {
        public Dictionary<string, int> Data { get; set; }   // 코드:가격
    }
    public class EvtAccount : ResponseData
    {

        public class DataModel
        {
            public string Account { get; set; } // 계좌번호
            public string Code { get; set; }
            public string Name { get; set; }
            public int Price { get; set; }      // 현재가
            public int UnitPrice { get; set; } // 매입단가
            public int Quantity { get; set; }   // 보유 수량
            public int SellOrBuy { get; set; }  // 이벤트 발생 1:매도, 2:매수


        }
        public DataModel Data { get; set; }
    }
}
