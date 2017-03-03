package finance

import (
	"github.com/uole/loso/convert"
	"github.com/uole/loso/http/client"
	"log"
	"strings"
)

var (
	shanhaiASymbolsURL string = "http://query.sse.com.cn/security/stock/downloadStockListFile.do?csrcCode=&stockCode=&areaName=&stockType=1"
	shanhaiBSymbolsURL string = "http://query.sse.com.cn/security/stock/downloadStockListFile.do?csrcCode=&stockCode=&areaName=&stockType=2"
)

type Symbol struct {
	Name         string
	Code         string
	Organization string
	Company      string
	MarketTime   string
}

type SymbolAdapter interface {
	Find() ([]Symbol, error)
}

type ShenZhenAdapter struct {
	Urls []string
}

type ShangHaiAdapter struct {
	Urls []string
}

func (a *ShenZhenAdapter) Parse(buffer []byte) ([]Symbol, error) {
	str := convert.ToUtf8(convert.CHARSET_GBK, string(buffer))
	lines := strings.Split(str, "\n")
	data := make([]Symbol,0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.Replace(line,"  ", "",  -1)
		line = strings.Replace(line,"\t", " ", -1)
		row := strings.Split(line," ")
		if len(row) == 7{
			data = append(data,Symbol{
				Name:row[1],
				Code:row[0],
				Organization:"上海交易所",
				Company:row[3],
				MarketTime:row[4],
			})
		}
	}
	return data, nil
}

func (m *ShenZhenAdapter) Find() ([]Symbol, error) {
	dataset := make([]Symbol, 0)
	for _, urlStr := range m.Urls {
		log.Println(urlStr)
		buffer, err := client.Request(urlStr, nil, client.METHOD_GET, nil)
		if err == nil {
			data, err := m.Parse(buffer)
			log.Println(data)
			if err == nil {
				old := dataset
				dataset = make([]Symbol,len(dataset) + len(data))
				copy(dataset,old)
				copy(dataset[len(old):],data)
			}else{
				log.Println(err)
			}
		}else{
			log.Println(err)
		}
	}
	log.Println(len(dataset))
	log.Println(dataset)
	return dataset, nil
}

func (a ShangHaiAdapter) Find() ([]Symbol, error) {

	return nil, nil
}

func FindAll() ([]Symbol, error) {

	return nil, nil
}
