package domains

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestExtractDomains(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "Alibaba",
			url:      "https://www.alibaba.com/",
			expected: "1688.com,activities.alibaba.com,activity.alibaba.com,ads.alibaba.com,alibaba.com,alibabacloud.com,alibabagroup.com,aliexpress.com,alimama.com,alios.cn,aliqin.cn,app.alibaba.com,arabic.alibaba.com,autonavi.com,buyer.alibaba.com,dingtalk.com,dutch.alibaba.com,facebook.com,fliggy.com,french.alibaba.com,german.alibaba.com,global.alipay.com,hebrew.alibaba.com,hindi.alibaba.com,i.alibaba.com,i.alicdn.com,idinfo.zjaic.gov.cn,img.alicdn.com,importer.alibaba.com,indonesian.alibaba.com,inspection.alibaba.com,ipp.alibabagroup.com,italian.alibaba.com,japanese.alibaba.com,ju.taobao.com,korean.alibaba.com,linkedin.com,logistics.alibaba.com,m.alibaba.com,onetouch.alibaba.com,portuguese.alibaba.com,rfq.alibaba.com,rule.alibaba.com,russian.alibaba.com,seller.alibaba.com,spanish.alibaba.com,supplier.alibaba.com,surveymonkey.com,taobao.com,taobao.lazada.sg,thai.alibaba.com,tmall.com,tradeassurance.alibaba.com,trademanager.alibaba.com,turkish.alibaba.com,twitter.com,ucweb.com,umeng.com,vietnamese.alibaba.com,wholesaler.alibaba.com,world.taobao.com,xiami.com,youtube.com",
		}, {
			name:     "Starter Studio",
			url:      "https://www.starterstudio.org/",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := http.Get(test.url)
			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)

			lines := strings.Split(string(bytes), "\n")

			if res := ExtractDomains(lines); res != test.expected {
				t.Fatalf("%s returned unexpected result: %q", test.name, res)
			}
		})
	}
}
