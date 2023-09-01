package cryptocompare_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/forbole/bdjuno/v4/client/cryptocompare"
)

func TestConvertCryptocomparePrices(t *testing.T) {
	result := `
{
	"CUDOS": {
		"USD": {
			"TYPE": "5",
			"MARKET": "CCCAGG",
			"FROMSYMBOL": "CUDOS",
			"TOSYMBOL": "USD",
			"FLAGS": "2050",
			"PRICE": 0.00237,
			"LASTUPDATE": 1680508413,
			"MEDIAN": 0.002369,
			"LASTVOLUME": 795.02960819,
			"LASTVOLUMETO": 1.8842201714103,
			"LASTTRADEID": "3347838618",
			"VOLUMEDAY": 1974357.8239218986,
			"VOLUMEDAYTO": 4679.2280426949,
			"VOLUME24HOUR": 6204684.682352951,
			"VOLUME24HOURTO": 14705.10269717649,
			"OPENDAY": 0.00239,
			"HIGHDAY": 0.002399,
			"LOWDAY": 0.002346,
			"OPEN24HOUR": 0.002392,
			"HIGH24HOUR": 0.002439,
			"LOW24HOUR": 0.0023417658,
			"LASTMARKET": "CoinEx",
			"VOLUMEHOUR": 144926.93789457,
			"VOLUMEHOURTO": 343.4768428101309,
			"OPENHOUR": 0.002359,
			"HIGHHOUR": 0.002373,
			"LOWHOUR": 0.002351,
			"TOPTIERVOLUME24HOUR": 6204684.682352951,
			"TOPTIERVOLUME24HOURTO": 14705.10269717649,
			"CHANGE24HOUR": -0.000021999999999999884,
			"CHANGEPCT24HOUR": -0.9197324414715671,
			"CHANGEDAY": -0.000020000000000000052,
			"CHANGEPCTDAY": -0.8368200836820106,
			"CHANGEHOUR": 0.000011000000000000159,
			"CHANGEPCTHOUR": 0.4662992793556659,
			"CONVERSIONTYPE": "multiply",
			"CONVERSIONSYMBOL": "USDT",
			"CONVERSIONLASTUPDATE": 1680508460,
			"SUPPLY": 8868274663,
			"MKTCAP": 21017810.95131,
			"MKTCAPPENALTY": 0,
			"CIRCULATINGSUPPLY": 5075184880,
			"CIRCULATINGSUPPLYMKTCAP": 12028188.1656,
			"TOTALVOLUME24H": 7250964.214818185,
			"TOTALVOLUME24HTO": 17184.785189119095,
			"TOTALTOPTIERVOLUME24H": 6346831.11155808,
			"TOTALTOPTIERVOLUME24HTO": 15041.989734392646,
			"IMAGEURL": "/media/38553724/cudos.png"
		},
		"ETH": {
			"TYPE": "5",
			"MARKET": "CCCAGG",
			"FROMSYMBOL": "CUDOS",
			"TOSYMBOL": "WETH",
			"FLAGS": "2052",
			"PRICE": 0.00000132,
			"LASTUPDATE": 1680500255,
			"MEDIAN": 0.00000131755152958313,
			"LASTVOLUME": 4643.826941689025,
			"LASTVOLUMETO": 0.006118481290141717,
			"LASTTRADEID": "0xc8c14d2acf38c0ff15efcd0a4630a4596fbe1f6cdd0d61b095c14efb93a9d4ed-0",
			"VOLUMEDAY": 166851.9458252716,
			"VOLUMEDAYTO": 0.2192929710691818,
			"VOLUME24HOUR": 904133.103246099,
			"VOLUME24HOURTO": 1.1892326371083,
			"OPENDAY": 0.00000132,
			"HIGHDAY": 0.00000132,
			"LOWDAY": 0.00000131,
			"OPEN24HOUR": 0.00000129,
			"HIGH24HOUR": 0.00000132,
			"LOW24HOUR": 0.00000129,
			"LASTMARKET": "uniswapv2",
			"VOLUMEHOUR": 0,
			"VOLUMEHOURTO": 0,
			"OPENHOUR": 0.00000132,
			"HIGHHOUR": 0.00000132,
			"LOWHOUR": 0.00000132,
			"TOPTIERVOLUME24HOUR": 0,
			"TOPTIERVOLUME24HOURTO": 0,
			"CHANGE24HOUR": 3.0000000000000136e-8,
			"CHANGEPCT24HOUR": 2.325581395348848,
			"CHANGEDAY": 0,
			"CHANGEPCTDAY": 0,
			"CHANGEHOUR": 0,
			"CHANGEPCTHOUR": 0,
			"CONVERSIONTYPE": "direct",
			"CONVERSIONSYMBOL": "",
			"CONVERSIONLASTUPDATE": 1680500255,
			"SUPPLY": 8868274663,
			"MKTCAP": 11706.12255516,
			"MKTCAPPENALTY": 0,
			"CIRCULATINGSUPPLY": 5075184880,
			"CIRCULATINGSUPPLYMKTCAP": 6699.2440416,
			"TOTALVOLUME24H": 7250964.214818185,
			"TOTALVOLUME24HTO": 9.567049704383454,
			"TOTALTOPTIERVOLUME24H": 6346831.11155808,
			"TOTALTOPTIERVOLUME24HTO": 8.377817067256666,
			"IMAGEURL": "/media/38553724/cudos.png"
		}
	}
}
`

	var apisPrices map[string]map[string]cryptocompare.MarketTicker
	err := json.Unmarshal([]byte(result), &apisPrices)
	require.NoError(t, err)
	cfg := cryptocompare.Config{
		Config: struct {
			CryptoCompareProdAPIKey string "yaml:\"crypto_compare_prod_api_key\""
			CryptoCompareFreeAPIKey string "yaml:\"crypto_compare_free_api_key\""
		}{
			CryptoCompareProdAPIKey: "test",
			CryptoCompareFreeAPIKey: "test",
		},
	}

	ccc := cryptocompare.NewClient(&cfg)
	prices := ccc.ConvertCryptocompare(apisPrices)
	require.Equal(t, 2, len(prices))
	priceValues := make([]float64, len(prices))
	marketCapValues := make([]int64, len(prices))
	for i, p := range prices {
		priceValues[i] = p.Price
		marketCapValues[i] = p.MarketCap
	}
	require.Contains(t, priceValues, float64(0.00237), "priceValues should contain USD VALUE")
	require.Contains(t, priceValues, float64(0.00000132), "priceValues should contain ETH VALUE")
	require.Contains(t, marketCapValues, int64(21017810), "marketCapValues should contain USD market cap")
	require.Contains(t, marketCapValues, int64(11706), "marketCapValues should contain ETH market cap")
}