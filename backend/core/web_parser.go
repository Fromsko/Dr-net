package core

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
)

// WebJsonLoad 导入json数据
func WebJsonLoad(jsonFInfo *JsonFile, netPir *NetworkPair) {
	jsonFile := filepath.Join(GetResDir(""), jsonFInfo.Local)

	file, _ := os.ReadFile(jsonFile)

	_ = json.Unmarshal(file, &netPir)
}

// GetRandData 随机获取网段数据
func GetRandData(areaAllData []*NetworkInfo, limit int) (limitInfo []*NetworkInfo) {
	limitInfo = make([]*NetworkInfo, 0)
	for i := 0; i < limit; i++ {
		randIndex := rand.Intn(len(areaAllData))
		randResult := areaAllData[randIndex]
		limitInfo = append(limitInfo, randResult)
	}
	return limitInfo
}
