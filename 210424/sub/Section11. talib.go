package sub

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func TalibExample() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	// Rsi란 종가의 변화로 추세 강도를 측정하는 선행지표
	// 헌재 추세강도가 어떠한지를 0~100퍼센트의 수치로 보여줌
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
	wma := talib.Wma(spy.Close, 14)
	fmt.Print(wma)
}
