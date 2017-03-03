package finance

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/uole/loso/convert"
)

type Fund struct {
	Name    string
	Code    string
	Predict float64
	Price   float64
	Type    string
}

func NewFund(code string) (*Fund, error) {
	strUrl := "http://stocks.sina.cn/fund/?vt=4&code=" + code

	//http://fundgz.1234567.com.cn/js/210009.js?rt=1488528176128
	doc, err := goquery.NewDocument(strUrl)
	if err != nil {
		return nil, err
	}
	f := &Fund{}
	f.Name = doc.Find(".fund_name").Text()
	f.Code = code
	f.Predict = convert.ParseFloat(doc.Find(".j_fund_value.fund_value").Text())
	f.Type = doc.Find(".fund_type ").Text()
	f.Price = convert.ParseFloat(doc.Find(".j_premium_rate.premium_rate").Text())
	return f, nil
}
