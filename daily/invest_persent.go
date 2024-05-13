package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	HuShen300Cost := 26750.00
	HuShen300Market := 29040.35
	HuShen300Earn := HuShen300Market - HuShen300Cost

	ChinaStockCost := 4574.00
	ChinaStockMarket := 5110.03
	ChinaStockEarn := ChinaStockMarket - ChinaStockCost

	ChinaInternet50Cost := 13000.00
	ChinaInternet50Market := 15277.37
	ChinaInternet50Earn := ChinaInternet50Market - ChinaInternet50Cost

	SP500Cost := 13203.00
	SP500Market := 14312.42

	SP500Earn := SP500Market - SP500Cost

	NASDAQCost := 10750.00
	NASDAQMarket := 11560.7
	NASDAQEarn := NASDAQMarket - NASDAQCost

	GoldStockCost := 3500.00
	GoldStockMarket := 3550.75
	GoldStockEarn := GoldStockMarket - GoldStockCost

	TotalMarket := HuShen300Market + NASDAQMarket + SP500Market + ChinaStockMarket + ChinaInternet50Market + GoldStockMarket
	TotalCost := HuShen300Cost + NASDAQCost + SP500Cost + ChinaStockCost + ChinaInternet50Cost + GoldStockCost
	TotalEarn := TotalMarket - TotalCost
	TotalEarnRate := Percent(TotalEarn, TotalCost)

	fmt.Printf("总市值：%s\n", cast.ToString(TotalMarket))
	fmt.Printf("总成本：%s\n", cast.ToString(TotalCost))
	fmt.Printf("总收益：%s\n", cast.ToString(TotalEarn))
	fmt.Printf("收益率：%s\n", cast.ToString(TotalEarnRate))

	fmt.Printf("沪深300收益：%.4f\n", HuShen300Earn)
	fmt.Printf("A股收益：%.4f\n", ChinaStockEarn)
	fmt.Printf("互联网收益：%.4f\n", ChinaInternet50Earn)
	fmt.Printf("标普500收益：%.4f\n", SP500Earn)
	fmt.Printf("纳斯达克收益：%.4f\n", NASDAQEarn)
	fmt.Printf("黄金收益：%.4f\n", GoldStockEarn)

	fmt.Printf("沪深300投入比：%s\n", Percent(HuShen300Cost, TotalCost))
	fmt.Printf("A股投入比：%s\n", Percent(ChinaStockCost, TotalCost))
	fmt.Printf("互联网收益投入比：%s\n", Percent(ChinaInternet50Cost, TotalCost))
	fmt.Printf("标普500投入比：%s\n", Percent(SP500Cost, TotalCost))
	fmt.Printf("纳斯达克投入比：%s\n", Percent(NASDAQMarket, TotalCost))
	fmt.Printf("黄金投入比：%s\n", Percent(GoldStockMarket, TotalCost))

	fmt.Printf("沪深300盈亏比：%s\n", Percent(HuShen300Earn, TotalEarn))
	fmt.Printf("A股盈亏比：%s\n", Percent(ChinaStockEarn, TotalEarn))
	fmt.Printf("互联网收益盈亏比：%s\n", Percent(ChinaInternet50Earn, TotalEarn))
	fmt.Printf("标普500盈亏比：%s\n", Percent(SP500Earn, TotalEarn))
	fmt.Printf("纳斯达克盈亏比：%s\n", Percent(NASDAQEarn, TotalEarn))
	fmt.Printf("黄金盈亏比：%s\n", Percent(GoldStockEarn, TotalEarn))

}

// Percent a除以b的百分比
func Percent(a, b float64) string {
	return fmt.Sprintf("%.2f%%", (a/b)*100)
}
