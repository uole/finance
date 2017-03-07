package finance

import (
	"errors"
	"github.com/uole/loso/arith"
	"github.com/uole/loso/convert"
	"github.com/uole/loso/http/client"
	"regexp"
	"strings"
)

type Stock struct {
	Name     string
	Code     string
	Open     float64
	Close    float64
	Quote    float64
	High     float64
	Low      float64
	Volume   float64
	Turnover float64
}

// get all stock code
func GetList() []string {
	return nil
}

// calc stock growth
//https://github.com/nzai/stockrecorder/blob/master/market/china.go
func (s *Stock) GrowthRate() float64 {
	v := (s.Quote - s.Close) / s.Close * 100
	if val, err := arith.Round(v, 3); err == nil {
		return val
	}
	return v
}

// get stock info
func StockQuote(code string) (*Stock, error) {
	if code == "" {
		return nil, errors.New("stock code can't empty")
	}
	var stockCode string
	if code[0] >= '0' && code[0] <= '9' {
		if code[0] == '0' || code[0] == '3' {
			stockCode = "sz" + code
		} else if code[0] == '6' {
			stockCode = "sh" + code
		} else {
			return nil, errors.New("unknow stock code " + code)
		}
	} else {
		stockCode = code
	}
	strUrl := "http://hq.sinajs.cn/list=" + stockCode
	buffer, err := client.Get(strUrl, nil)
	if err != nil {
		return nil, err
	}
	content := convert.ToUtf8(convert.CHARSET_GBK, string(buffer))
	reg := regexp.MustCompile(`.*?="(.*?)";`)
	if groups := reg.FindStringSubmatch(content); groups != nil {
		content = groups[1]
		data := strings.Split(content, ",")
		if len(data) >= 9 {
			stock := &Stock{
				Name:     data[0],
				Code:     code,
				Quote:    convert.ParseFloat(data[3]),
				Open:     convert.ParseFloat(data[1]),
				Close:    convert.ParseFloat(data[2]),
				High:     convert.ParseFloat(data[4]),
				Low:      convert.ParseFloat(data[5]),
				Turnover: convert.ParseFloat(data[9]),
				Volume:   convert.ParseFloat(data[8]),
			}
			return stock, nil
		} else {
			return nil, errors.New("")
		}
	} else {
		return nil, err
	}
}
