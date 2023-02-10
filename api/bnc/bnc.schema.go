package bnc

type Token struct {
	Symbol string
	Price  string
}

type Kline struct {
	openTime  int
	open      string
	high      string
	low       string
	close     string
	closeTime int
}

type ListTokenPriceDTO struct {
	trend   string `validate:"required, string"`
	percent string `validate:"required, float"`
	time    string `validate:"required, string"`
}
