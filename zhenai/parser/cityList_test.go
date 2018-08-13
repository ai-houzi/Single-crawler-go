package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	body, err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		panic(err)
	}

	//ParseCityList(body)

	result := ParseCityList(body)

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"city:阿坝", "city:阿克苏", "city:阿拉善盟",
	}

	const resultsize = 470

	if len(result.Requsets) != resultsize {
		t.Errorf("result should have %d request,but hand %d", resultsize, len(result.Requsets))
	}

	for i, url := range expectedUrls {
		if result.Requsets[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requsets[i].Url)
		}
	}

	if len(result.Iterms) != resultsize {
		t.Errorf("result should have %d Iterms,but hand %d", resultsize, len(result.Iterms))
	}

	for i, city := range expectedCities {
		if result.Iterms[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Iterms[i].(string))
		}
	}

}
