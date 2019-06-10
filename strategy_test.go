package backTrace

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer(t *testing.T) {
	testLogger := logrus.WithFields(logrus.Fields{
		"function": "TestGetStock()",
	})
	stockData := GetSockData("600018")
	if len(stockData) > 0 {
		testLogger.Infof("find stock code numbers is %d", len(stockData))
		buy := BreakOutStrategyBuy{}
		sell := BreakOutStrategySell{}
		ana := Analyzer{BuyPolicies: []Strategy{&buy},
			SellPolicies: []Strategy{&sell}}
		res, err := ana.Analyse(stockData)
		if err != nil {
			testLogger.Fatal("analyse function is err!", err)
		}
		assert.Equal(t, len(res), len(stockData))
		hold_n := 0
		sell_n := 0
		buy_n := 0
		for _, r := range res {
			if r == 2 {
				hold_n += 1
			} else if r == 1 {
				sell_n += 1
			} else {
				buy_n += 1
			}
		}
		testLogger.Println(fmt.Sprintf("hold: %d, sell: %d, buy: %d", hold_n, sell_n, buy_n))
	} else {
		testLogger.Fatal("can't find the stock in the database!")
	}

}