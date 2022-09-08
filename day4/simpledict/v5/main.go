package v5

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`validperiod=&wlanAcIp=&isLocalUser=2&passType=1&freePkgTotal=15000%25E5%2588%2586%25E9%2592%259F%25EF%25BC%2588%25E5%2585%25B6%25E4%25B8%25AD%25E6%25A0%25A1%25E5%259B%25AD15000%25E5%2588%2586%25E9%2592%259F%25EF%25BC%2589%252C0MB&logonsessid=&wlanAcName=0436.0571.571.00&userType=&wlanFees=0.00&usedFree=4020%25E5%2588%2586%25E9%2592%259F%25EF%25BC%2588%25E5%2585%25B6%25E4%25B8%25AD%25E6%25A0%25A1%25E5%259B%25AD4020%25E5%2588%2586%25E9%2592%259F%25EF%25BC%2589%252C0MB&booktime=&wlanUserIp=10.128.147.189&AUTO_LOGIN=true&ssid=&userName=19857180404&encryUser=&productName=%25E6%25A0%25A1%25E5%259B%25ADWLAN40%25E5%2585%2583%25E5%25A5%2597%25E9%25A4%2590&feesSum=0.00`)
	req, err := http.NewRequest("POST", "https://120.199.39.54:7090/zmcc/portalLoginRedirect.wlan", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", `nActions="http://111.1.47.58/school/wlan.do?areaCode=571"; ALVERIFY=0; theme_name=THEME_PC`)
	req.Header.Set("Origin", "https://120.199.39.54:7090")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://120.199.39.54:7090/zmcc/portalLogin.wlan?1653378875389")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
