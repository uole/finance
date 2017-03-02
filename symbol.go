package finance

var (
	shanhaiASymbolsURL string = "http://query.sse.com.cn/security/stock/downloadStockListFile.do?csrcCode=&stockCode=&areaName=&stockType=1"
	shanhaiBSymbolsURL string = "http://query.sse.com.cn/security/stock/downloadStockListFile.do?csrcCode=&stockCode=&areaName=&stockType=2"
)

type Symbol struct {
	Name         string
	Code         string
	Organization string
	Company      string
	marketTime   string
}
