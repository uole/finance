package finance

import (
	"encoding/json"
	"errors"
	"github.com/uole/gokit/arith"
	"github.com/uole/gokit/convert"
	"github.com/uole/gokit/http/client"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type Fund struct {
	Name        string
	Code        string
	Quote       float64
	QuoteTime   time.Time
	Predict     float64
	PredictTime time.Time
	Type        string
}

func (s *Fund) GrowthRate() float64 {
	v := (s.Predict - s.Quote) / s.Quote * 100
	if val, err := arith.Round(v, 3); err == nil {
		return val
	}
	return v
}

type fundData struct {
	Fundcode string `json:"fundcode"` //基金代码
	Name     string `json:"name"`     //基金名称
	Jzrq     string `json:"jzrq"`     //净值日期
	Dwjz     string `json:"dwjz"`     //单位净值
	Gsz      string `json:"gsz"`      //估算值
	Gszzl    string `json:"gszzl"`    //估算涨幅
	Gztime   string `json:"gztime"`   //估值时间
}

func FundQuote(code string) (*Fund, error) {
	urlStr := "http://fundgz.1234567.com.cn/js/" + code + ".js?rt=" + strconv.Itoa(rand.Int())
	buffer, err := client.Get(urlStr, nil)
	if err != nil {
		return nil, err
	}
	reg := regexp.MustCompile(`({.*?})`)
	matches := reg.FindStringSubmatch(string(buffer))
	if len(matches) > 0 {
		jsonFund := &fundData{}
		if err = json.Unmarshal([]byte(matches[0]), jsonFund); err == nil {
			preTime, err := time.Parse("2006-01-02 15:04", jsonFund.Gztime)
			if err != nil {
				preTime = time.Now()
			}
			quoteTime, err := time.Parse("2006-01-02", jsonFund.Jzrq)
			if err != nil {
				quoteTime = time.Now()
			}
			return &Fund{
				Name:        jsonFund.Name,
				Code:        jsonFund.Fundcode,
				Predict:     convert.ParseFloat(jsonFund.Gsz),
				PredictTime: preTime,
				Quote:       convert.ParseFloat(jsonFund.Dwjz),
				QuoteTime:   quoteTime,
			}, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("not match")
	}
}
