package constance

import (
	"time"
)

type PAGING struct {
	Page  int
	Limit int
}

func DefaultPaging() PAGING {
	p := PAGING{
		Page:  1,
		Limit: 10,
	}

	return p
}

const JWT_TIME = time.Second * 60
const REFRESH_TIME = time.Hour * 24 * 30

var TOKEN = map[string]int{
	"AccessToken":  1,
	"RefreshToken": 2,
}
