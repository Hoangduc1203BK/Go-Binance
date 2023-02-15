package bnc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    = os.Getenv("API_KEY")
	secretKey = os.Getenv("SECRET_KEY")
)
var client = binance.NewClient(apiKey, secretKey)

func ServiceGetTokenPrice(symbol *string) (interface{}, error) {
	url := fmt.Sprintf("https://www.binance.com/api/v3/ticker/price?symbol=%s", *symbol)
	res, err := http.Get(url)

	if err != nil {
		return Token{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	parseByte := string(body)
	var result Token
	json.Unmarshal([]byte(parseByte), &result)

	return result, nil
}

func ServiceGetKline(symbol *string, period *string) (Kline, error) {
	klines, err := client.NewKlinesService().Symbol(*symbol).Interval(*period).Do(context.Background())
	if err != nil {
		return Kline{}, err
	}

	latestKline := klines[len(klines)-1]
	result := Kline{
		openTime:  int(latestKline.OpenTime),
		open:      latestKline.Open,
		high:      latestKline.High,
		low:       latestKline.Low,
		close:     latestKline.Close,
		closeTime: int(latestKline.CloseTime),
	}

	return result, nil
}

func checkPrice(symbol *string, price *string, payload *ListTokenPriceDTO) (bool, error) {
	latest, err := ServiceGetKline(symbol, &payload.time)

	if err != nil {
		return false, nil
	}

	p1, err1 := strconv.ParseFloat(*price, 64)
	p2, err2 := strconv.ParseFloat(*&latest.close, 64)

	if err1 != nil || err2 != nil {
		return false, fmt.Errorf("Fail to check price")
	}

	parsePercent, err := strconv.ParseFloat(payload.percent[:len(payload.percent)-1], 64)

	if err != nil {
		return false, fmt.Errorf("Fail to check price")
	}
	ratio := math.Pow(10, float64(8))
	rounded := math.Round((p1-p2)*ratio) / ratio
	percent := math.Round(parsePercent*ratio) / ratio
	var result bool

	if payload.trend == "increase" {
		if p1 >= p2 && rounded == percent/100 {
			result = true
		}
	} else {
		if p1 <= p2 && -rounded == percent/100 {
			result = false
		}
	}

	return result, nil
}

func ServiceListTokenPrice(payload *ListTokenPriceDTO) ([]Token, error) {
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	tokens := []Token{}
	for i, p := range prices {
		if i <= 100 {
			check, _ := checkPrice(&p.Symbol, &p.Price, payload)
			if check == true {
				token := Token{
					Price:  p.Price,
					Symbol: p.Symbol,
				}
				tokens = append(tokens, token)
			}
		}
	}

	return tokens, nil
}
