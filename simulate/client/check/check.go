package check

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

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
func ckLogin() (string, error) {
	fetchUrl := "http://10.253.0.1/drcom/chkstatus?callback=dr1002&jsVersion=4.X&v=4197&lang=zh"

	content, _ := http.Get(fetchUrl)

	defer func() {
		err := content.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	resp, err := io.ReadAll(content.Body)
	if err != nil {
		panic(err)
	}

	html, err := decodeResponse(string(resp))
	if err != nil {
		return "", errors.New("解码失败")
	}

	return html, nil
}

// 正则提取
func regxParser(response string) (bool, *SaveType, error) {
	var saveType SaveType

	// 使用正则匹配 json 字符串部分
	re := regexp.MustCompile(`\(([\s\S]*)\)`)
	match := re.FindStringSubmatch(response)
	if len(match) != 2 {
		return false, nil, errors.New("无法匹配")
	}

	// 将匹配到的字符串解析为 json 对象
	var m map[string]interface{}
	err := json.Unmarshal([]byte(match[1]), &m)
	if err != nil {
		return false, nil, errors.New("JSON解析失败")
	}

	if uid, ok := m["uid"]; ok {
		str := fmt.Sprintf("%v", uid)
		saveType.SchoolCard = str
		log.Println("[ 账号 ]     => ", uid)
	} else {
		return false, nil, errors.New("IP不可用")
	}

	if nid, ok := m["NID"]; ok {
		str := fmt.Sprintf("%v", nid)
		saveType.UserName = str
		log.Println("[ 用户名 ]   => ", nid)
	}

	if v4ip, ok := m["v4ip"]; ok {
		str := fmt.Sprintf("%v", v4ip)
		saveType.IPAddress = str
		log.Println("[ IP地址 ]   => ", v4ip)
	}

	if loginTime, ok := m["stime"]; ok {
		str := fmt.Sprintf("%v", loginTime)
		saveType.LoginTime = str
		log.Println("[ 登录时间 ] => ", loginTime)
	}

	return true, &saveType, nil
}

// 编码解析
func decodeResponse(body string) (string, error) {
	decoder := simplifiedchinese.GB18030.NewDecoder()
	result, err := decoder.String(body)
	if err != nil {
		log.Println("decode error:", err)
		return "", err
	}
	return result, nil
}

func Writefile(filePath string, content *SaveType) {
	// 解析为JSON数据
	data, err := json.Marshal(content)
	checkError(err)

	// 将json格式的数据写入文件
	err = ioutil.WriteFile(filePath, data, 0777)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// LoginCheck 检测函数
func LoginCheck() (bool, *SaveType, error) {
	htmlData, err := ckLogin()
	if err != nil {
		return false, nil, err
	}

	flag, parserContent, err := regxParser(htmlData)
	if err != nil {
		return false, nil, err
	}

	Writefile("login-info.json", parserContent)

	return flag, parserContent, nil
}
