#!/bin/bash

curl -I --location --request GET 'https://yfd.imau.edu.cn/nonlogin/yiban/authentication/ff8080816ef7cbdf016ef7e70a650000.htm?verify_request=91bc00809d1864febba0726b2ac7d2b14e4ab6af7016bd88ab3425269e50e60b629a80bca46b98bf5cbd0e9303aa7d81bdbb227fd34ef7ccf274978fe4c8126f7813c5566feb92e66f0c5a3e56743d7f33467b6b5ce08d3c83291602c04c493eb887e38bf7cbc962d9959bc62db7616c84ed14aee95249dcae46a3069411606198aa446b81c6afc416274490989ef040f28087bb9e014cc82369fd3b14265a2d4fcd5417102b934df2f5406269fbf9fc1c25cf2c718fc5ec3fcf957ac4546cab0e97c54f424fc02658a3f3570b62886e7290ae3bd4c7b793c1d0182562e97e23e0f09ac9c0515f000d87c0a6fbddc0b20992a83d9e2860545c1d84ad783c005ab8140795646e847aa0ca5511ec2f643f7d0668558fe93be9389d1d00751c7ca0e0b574ed2a93f2d3f24d86d16733fb4015324df0b101c04799c73ef70fb66b3b74d5baeb77d7259a7e2ad5c082011e8cba50fd16785b65b76480625f4054816d&yb_uid=61375796' \
--header 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
--header 'User-Agent: Mozilla/5.0 (iPad; CPU OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 yiban_iOS/5.0.12' \
--header 'Referer: https://yfd.imau.edu.cn/webApp/xuegong/index.html' \
--header 'Accept-Language: zh-cn' \
--header 'Host: yfd.imau.edu.cn' \
--header 'Connection: keep-alive' \
--header 'loginToken: 2dec08501bd77f0abf9edc962deb8f56' \
--header 'AppVersion: 5.0.12'

a=`curl -I --location --request GET 'https://yfd.imau.edu.cn/nonlogin/yiban/authentication/ff8080816ef7cbdf016ef7e70a650000.htm?verify_request=91bc00809d1864febba0726b2ac7d2b14e4ab6af7016bd88ab3425269e50e60b629a80bca46b98bf5cbd0e9303aa7d81bdbb227fd34ef7ccf274978fe4c8126f7813c5566feb92e66f0c5a3e56743d7f33467b6b5ce08d3c83291602c04c493eb887e38bf7cbc962d9959bc62db7616c84ed14aee95249dcae46a3069411606198aa446b81c6afc416274490989ef040f28087bb9e014cc82369fd3b14265a2d4fcd5417102b934df2f5406269fbf9fc1c25cf2c718fc5ec3fcf957ac4546cab0e97c54f424fc02658a3f3570b62886e7290ae3bd4c7b793c1d0182562e97e23e0f09ac9c0515f000d87c0a6fbddc0b20992a83d9e2860545c1d84ad783c005ab8140795646e847aa0ca5511ec2f643f7d0668558fe93be9389d1d00751c7ca0e0b574ed2a93f2d3f24d86d16733fb4015324df0b101c04799c73ef70fb66b3b74d5baeb77d7259a7e2ad5c082011e8cba50fd16785b65b76480625f4054816d&yb_uid=61375796' \
--header 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
--header 'User-Agent: Mozilla/5.0 (iPad; CPU OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 yiban_iOS/5.0.12' \
--header 'Referer: https://yfd.imau.edu.cn/webApp/xuegong/index.html' \
--header 'Accept-Language: zh-cn' \
--header 'Host: yfd.imau.edu.cn' \
--header 'Connection: keep-alive' \
--header 'loginToken: 2dec08501bd77f0abf9edc962deb8f56' \
--header 'AppVersion: 5.0.12' | grep Set-Cookie`

echo $a
a=${a: 11}
echo $a > 1.txt
