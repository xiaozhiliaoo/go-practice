package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type ZhSpotEm struct {
	Id                int     `json:"序号"`
	Code              string  `json:"代码"`
	Name              string  `json:"名称"`
	LatestPrice       float64 `json:"最新价"`
	RisePrice         float64 `json:"涨跌幅"`
	RiseAmount        float64 `json:"涨跌额"`
	Amount            float64 `json:"成交量"`
	Money             float64 `json:"成交额"`
	Amplitude         float64 `json:"振幅"`
	Max               float64 `json:"最高"`
	Min               float64 `json:"最低"`
	TodayOpen         float64 `json:"今开"`
	YesterdayClose    float64 `json:"昨收"`
	AmountRatio       float64 `json:"量比"`
	ChangeHandsRate   float64 `json:"换手率"`
	PE                float64 `json:"市盈率-动态"`
	PB                float64 `json:"市净率"`
	TotalMarketValue  float64 `json:"总市值"`
	LiquidMarketValue float64 `json:"流通市值"`
	RiseSpeed         float64 `json:"涨速"`
	MinuteRise        float64 `json:"5分钟涨跌"`
	DayRise           float64 `json:"60日涨跌幅"`
	YearRise          float64 `json:"年初至今涨跌幅"`
}

type UsDaily struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

var str = `[{"序号":1,"代码":"831445","名称":"龙竹科技","最新价":4.97,"涨跌幅":29.77,"涨跌额":1.14,"成交量":34693,"成交额":16598786.95,"振幅":27.68,"最高":4.97,"最低":3.91,"今开":3.91,"昨收":3.83,"量比":8.23,"换手率":3.75,"市盈率-动态":101.11,"市净率":2,"总市值":741255744,"流通市值":459508676,"涨速":0,"5分钟涨跌":0,"60日涨跌幅":-7.45,"年初至今涨跌幅":-22.22}]`
var usStr = `{"date":"1986-03-13T00:00:00.000","open":28,"high":29.25,"low":25.5,"close":28,"volume":3582600}`

func main() {
	var em UsDaily
	err := jsoniter.UnmarshalFromString(usStr, &em)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(em.Date)

	fmt.Println("---------------")

	toString, _ := jsoniter.MarshalToString(em)

	fmt.Println(toString)

}
