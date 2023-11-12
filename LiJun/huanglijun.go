package LiJun

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type HuangLijun struct {
}

func (HuangLijun *HuangLijun) getAccessToken(refreshToken string) (string, string, error) {
	url := "https://auth.aliyundrive.com/v2/account/token"
	var dataMap = make(map[string]string)
	dataMap["grant_type"] = "refresh_token"
	dataMap["refresh_token"] = refreshToken
	dataByte, _ := json.Marshal(dataMap)
	req, err := http.NewRequest("POST", url, bytes.NewReader(dataByte))
	if err != nil {
		return "", "", err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	var resMap map[string]interface{}
	json.Unmarshal(body, &resMap)
	if accessToken, ok := resMap["access_token"].(string); ok {
		if nick_name, ok := resMap["nick_name"].(string); ok {
			return accessToken, nick_name, nil
		}
		return accessToken, "", nil
	}
	return "", "", errors.New("请稍后再试")
}

func (HuangLijun *HuangLijun) signIn(accessToken string) (string, error) {
	url := "https://member.aliyundrive.com/v1/activity/sign_in_list"
	data := []byte(`{
		"_rx-s":"mobile"
	}`)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", accessToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var resMap map[string]interface{}
	json.Unmarshal(body, &resMap)
	res := resMap["result"].(map[string]interface{})["signInCount"].(float64)
	signInCount := strconv.FormatFloat(res, 'f', 0, 64)
	return signInCount, nil
}

func (HuangLijun *HuangLijun) QYWX(messages string){
	data := []byte(buildMsg(messages, true))
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=e28f7b52-155f-4ac7-a824-77e285ddd086"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(messages))
}

func (HuangLijun *HuangLijun) getReward(accessToken string, signInCount string) (string, error) {
	url := "https://member.aliyundrive.com/v1/activity/sign_in_reward?_rx-s=mobile"
	var dataMap = make(map[string]string)
	dataMap["signInDay"] = signInCount
	dataByte, _ := json.Marshal(dataMap)
	req, err := http.NewRequest("POST", url, bytes.NewReader(dataByte))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", accessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var resMap map[string]interface{}
	json.Unmarshal(body, &resMap)
	if reward, ok := resMap["result"].(map[string]interface{})["notice"].(string); ok {
		return reward, nil
	}
	return "", errors.New("获取侍寝奖励失败")
}

func (HuangLijun *HuangLijun) qianDao(refreshToken string) (string, string, string, error) {
	accessToken, nick_name, err := HuangLijun.getAccessToken(refreshToken)
	if err != nil {
		return "", "", "", err
	}
	signInCount, err := HuangLijun.signIn(accessToken)
	if err != nil {
		return "", "", "", err
	}
	reward, err := HuangLijun.getReward(accessToken, signInCount)
	if err != nil {
		return "", "", "", err
	}
	return signInCount, reward, nick_name, nil
}

func (HuangLijun *HuangLijun) Run(refreshToken string) {
	var signInCount string
	var reward string
	var nick_name string
	var err error
	var title = "黄丽君 侍寝 \n"
	var sendMessages string
	signInCount, reward, nick_name, err = HuangLijun.qianDao(refreshToken)
	if err != nil {
		if err.Error() == "refreshToken过期,请更改后重试" {
			fmt.Println("请稍后再试")
		} else {
			for i := 0; i < 100; i++ {
				signInCount, reward, nick_name, err = HuangLijun.qianDao(refreshToken)
				if err == nil {
					content := "黄丽君：" + nick_name + " =>> 正在侍寝, 将奖励==>" + reward + ", 本月侍寝" + signInCount + "次 "
					fmt.Println(content)
					sendMessages += content + "\n"
					break
				}
			}
		}
	} else {
		content := "黄丽君：" + nick_name + " =>> 今日已侍寝, 已奖励==>" + reward + ", 本月侍寝" + signInCount + "次 "
		fmt.Println(content)
		sendMessages += content + "\n"
	}
	HuangLijun.QYWX(title+"\n"+sendMessages)
}
