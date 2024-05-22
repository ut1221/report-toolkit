package pkg

import (
	"fmt"
	"github.com/tidwall/gjson"
	"net/url"
	"report-toolkit/cmd/config"
	"report-toolkit/cmd/utils"
)

// Translate
//
//	@Description: 百度翻译
//	@Author PTJ 2024-05-22 23:33:22
//	@param query
//	@param from
//	@param to
//	@param conf
//	@return string
//	@return error
func Translate(query, from, to string, conf *config.BaiduTranslate) (string, error) {
	salt := "1435660288"
	data := url.Values{}
	data.Set("q", query)
	data.Set("salt", salt)
	data.Set("appid", conf.AppId)
	data.Set("from", from)
	data.Set("to", to)
	data.Set("sign", buildSign(query, salt, conf))
	res, err := utils.PostForm("https://fanyi-api.baidu.com/api/trans/vip/translate", data)
	res = gjson.Get(res, "trans_result.0.dst").String()
	res = utils.UrlDecode(res)
	return res, err
}
func buildSign(query, salt string, conf *config.BaiduTranslate) string {
	str := fmt.Sprintf("%s%s%s%s", conf.AppId, query, salt, conf.Secure)
	ret := utils.Md5(str)
	return ret
}
