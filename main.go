package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getCookie() *http.Client {
	client := http.Client{
		/*
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        100,              // 最大连接数,默认0无穷大
				MaxIdleConnsPerHost: 100,              // 对每个host的最大连接数量(MaxIdleConnsPerHost<=MaxIdleConns)
				IdleConnTimeout:     90 * time.Second, // 多长时间未使用自动关闭连接
			},
		*/
	}

	req, err := http.NewRequest("GET", "https://yfd.imau.edu.cn/nonlogin/yiban/authentication/ff8080816ef7cbdf016ef7e70a650000.htm?verify_request=91bc00809d1864febba0726b2ac7d2b14e4ab6af7016bd88ab3425269e50e60b629a80bca46b98bf5cbd0e9303aa7d81bdbb227fd34ef7ccf274978fe4c8126f7813c5566feb92e66f0c5a3e56743d7f33467b6b5ce08d3c83291602c04c493eb887e38bf7cbc962d9959bc62db7616c84ed14aee95249dcae46a3069411606198aa446b81c6afc416274490989ef040f28087bb9e014cc82369fd3b14265a2d4fcd5417102b934df2f5406269fbf9fc1c25cf2c718fc5ec3fcf957ac4546cab0e97c54f424fc02658a3f3570b62886e7290ae3bd4c7b793c1d0182562e97e23e0f09ac9c0515f000d87c0a6fbddc0b20992a83d9e2860545c1d84ad783c005ab8140795646e847aa0ca5511ec2f643f7d0668558fe93be9389d1d00751c7ca0e0b574ed2a93f2d3f24d86d16733fb4015324df0b101c04799c73ef70fb66b3b74d5baeb77d7259a7e2ad5c082011e8cba50fd16785b65b76480625f4054816d&yb_uid=61375796", nil)
	if err != nil {
		os.Exit(1)
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 yiban_iOS/5.0.12")
	req.Header.Add("Referer", "https://yfd.imau.edu.cn/webApp/xuegong/index.html")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("Host", "yfd.imau.edu.cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("loginToken", "2dec08501bd77f0abf9edc962deb8f56")
	req.Header.Add("AppVersion", "5.0.12")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	// 这里有坑 如果不设置 cookie 那么 charles 抓到 set-cookie 了，但是 go 过滤了 cookie，导致没有
	req.Header.Add("Cookie", "JSESSIONID=9B678EA2D7B0CBCE6F5F49FEC1C634D9-n1.yfd1")

	client.Timeout = 1 * time.Minute
	_, err = client.Do(req)

	if err != nil {
		os.Exit(1)
	}

	time.Sleep(1 * time.Second)
	/*
		if len(resp.Cookies()) > 0 {
			fmt.Println(resp.Cookies()[0].String())
			return &client
		}
	*/
	// u := &url.URL{Path: "https://yfd.imau.edu.cn"}
	// fmt.Println(client.Jar.Cookies(u))
	return &client
}

func operation(c string) {
	client := http.Client{}
	str := "data=%7B%22xmqkb%22%3A%7B%22id%22%3A%22ff808081706d93e201707f2f7e86525d%22%7D%2C%22c6%22%3A%22%E5%90%A6%22%2C%22c1%22%3A%22%E5%90%A6%22%2C%22c2%22%3A%22%E5%90%A6%22%2C%22c3%22%3A%22%E5%90%A6%22%2C%22c4%22%3A%22%E5%90%A6%22%2C%22c5%22%3A%22%E5%90%A6%22%2C%22c8%22%3A%22%E5%90%A6%22%2C%22c9%22%3A%22%E5%90%A6%22%2C%22c11%22%3A%2236.6%22%2C%22c7%22%3A%22%E5%AE%8C%E6%88%90%E7%AC%AC%E4%BA%8C%E9%92%88%E6%8E%A5%E7%A7%8D%22%2C%22c20%22%3A%22%E5%90%A6%22%2C%22type%22%3A%22YQSJSB%22%2C%22location_longitude%22%3A109.93735270182292%2C%22location_latitude%22%3A35.185643446180556%2C%22location_address%22%3A%22%E9%99%95%E8%A5%BF%E7%9C%81%E6%B8%AD%E5%8D%97%E5%B8%82%E6%BE%84%E5%9F%8E%E5%8E%BF%E9%9D%92%E6%AD%A3%E8%A1%97%E9%9D%A0%E8%BF%91%E7%9B%9B%E5%9F%BA%E5%BA%9C%22%7D&msgUrl=syt%2Fzzapply%2Flist.htm%3Ftype%3DYQSJSB%26xmid%3Dff808081706d93e201707f2f7e86525d&uploadFileStr=%7B%7D&multiSelectData=%7B%7D&type=YQSJSB"
	req, err := http.NewRequest("POST", "https://yfd.imau.edu.cn/syt/zzapply/operation.htm", strings.NewReader(str))
	if err != nil {
		fmt.Println(111)
		os.Exit(1)
	}

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 yiban_iOS/5.0.12")
	req.Header.Add("Referer", "https://yfd.imau.edu.cn/webApp/xuegong/index.html")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", "https://yfd.imau.edu.cn")
	cookie := strings.TrimSpace(c)
	req.Header.Add("Cookie", cookie)

	client.Timeout = 1 * time.Minute
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	byte, _ := io.ReadAll(resp.Body)
	fmt.Println(string(byte))
}

func main() {
	getC()
	str, err := os.ReadFile("1.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(str))
	operation(string(str))
}

func getC() {
	cmd := exec.Command("./a.sh")
	cmd.Run()
}
