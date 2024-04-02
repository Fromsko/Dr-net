package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Fromsko/gouitls/knet"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// SaveType 存储结构
type SaveType struct {
	SchoolCard string `json:"school_card"`
	UserName   string `json:"user_name"`
	IPAddress  string `json:"ip_address"`
	LoginTime  string `json:"login_time"`
}

// 检测登陆
func ckLogin() (html string, err error) {
	spider := knet.SendRequest{
		FetchURL: "http://10.253.0.1/drcom/chkstatus?callback=dr1002&jsVersion=4.X&v=4197&lang=zh",
	}
	spider.Send(func(resp []byte, cookies []*http.Cookie, _ error) {
		html, err = decodeResponse(string(resp))
	})
	return html, err
}

// 正则提取
func regxParser(response string) (*SaveType, error) {
	var saveType SaveType

	// 使用正则匹配 json 字符串部分
	re := regexp.MustCompile(`\(([\s\S]*)\)`)
	match := re.FindStringSubmatch(response)
	if len(match) != 2 {
		return nil, errors.New("无法匹配")
	}

	// 将匹配到的字符串解析为 json 对象
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(match[1]), &m); err != nil {
		return nil, errors.New("JSON解析失败")
	}

	if uid, ok := m["uid"]; ok {
		str := fmt.Sprintf("%v", uid)
		saveType.SchoolCard = str
	}

	if nid, ok := m["NID"]; ok {
		str := fmt.Sprintf("%v", nid)
		saveType.UserName = str
	}

	if v4ip, ok := m["v4ip"]; ok {
		str := fmt.Sprintf("%v", v4ip)
		saveType.IPAddress = str
	}

	if loginTime, ok := m["stime"]; ok {
		str := fmt.Sprintf("%v", loginTime)
		saveType.LoginTime = str
	}
	return &saveType, nil
}

// 编码解析
func decodeResponse(body string) (string, error) {
	decoder := simplifiedchinese.GB18030.NewDecoder()
	result, err := decoder.String(body)
	if err != nil {
		return "", err
	}
	return result, nil
}

// LoginCheck 检测函数
func LoginCheck() (parserContent *SaveType, err error) {
	htmlData, err := ckLogin()
	if err != nil {
		return nil, err
	}

	parserContent, err = regxParser(htmlData)
	if err != nil {
		return nil, err
	}
	return parserContent, nil
}
