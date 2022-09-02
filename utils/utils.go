package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func Feishu(str string) {
	http.Post("https://open.feishu.cn/open-apis/bot/v2/hook/a9690171-ec55-46fb-98d8-bf375554dc7f",
		"Content-Type: application/json",
		strings.NewReader(fmt.Sprintf("{\"msg_type\":\"text\",\"content\":{\"text\":\"%s\"}}", str)))
}
