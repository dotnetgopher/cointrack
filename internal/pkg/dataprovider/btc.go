package dataprovider

type BtcDataProvider struct {
	apiAddress string
	apiKey string
	proxyURL   string
	proxyPort  string

}


func New(apiAddress string, apiKey string, proxyURL string, proxyPort string) *BtcDataProvider {
	return &BtcDataProvider {
		apiAddress: apiAddress,
		apiKey: apiKey,
		proxyURL: proxyURL,
		proxyPort: proxyPort,
	}
}

func 