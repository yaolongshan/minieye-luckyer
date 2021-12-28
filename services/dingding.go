package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dingtalkrobot_1_0 "github.com/alibabacloud-go/dingtalk/robot_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/*
@brief: 发送钉钉消息给指定用户
@param appKey: 钉钉机器人的 appKey
@param appSecret: 钉钉机器人的 appSecret
@param userPhone: 要接收消息的用户手机号码
@param userName: 要接收消息的用户名字
@param awardLevel：奖品名字
*/
func SendDingDingMsg(
	appKey string,
	appSecret string,
	userPhone string,
	userName string,
	awardLevel string,
) error {
	accessToken, err := genAccessToken(appKey, appSecret)
	if err != nil {
		log.Printf("genAccessToken failed, err = %v", err)
		return err
	}

	userID, err := getUserIDByPhone(accessToken, userPhone)
	if err != nil {
		log.Printf("getUserIDByPhone failed, err = %v", err)
		return err
	}

	log.Printf("userID = %v\n", userID)

	err = sendMsg(
		appKey,
		userID,
		userName,
		awardLevel,
		accessToken,
	)
	if err != nil {
		log.Printf("sendMsg failed, err = %v", err)
		return err
	}

	return nil
}

func genAccessToken(
	appKey string,
	appSecret string,
) (string, error) {
	accessTokenFileName := "/tmp/minieye_robot_access_token"
	data1, err1 := ioutil.ReadFile(accessTokenFileName)
	if err1 != nil {
		data2, err2 := genAccessTokenByHTTP(appKey, appSecret)
		if err2 != nil {
			return "", err2
		} else {
			ioutil.WriteFile(accessTokenFileName, []byte(data2), 0644)
			return string(data2), nil
		}
	}
	return string(data1), nil
}

/*
@brief: 通过 appKey 和 appSecret 生成 access_token
// 通过 AppKey 和 AppSecret 获取一个 access_token
https://open.dingtalk.com/document/orgapp-server/obtain-orgapp-token
https://oapi.dingtalk.com/gettoken?appkey=xxx&appsecret=xxx
{
	"errcode": 0,
	"access_token": "xxxxxx",
	"errmsg": "ok",
	"expires_in": 7200
}
*/
func genAccessTokenByHTTP(
	appKey string,
	appSecret string,
) (string, error) {
	log.Printf("genAccessTokenByHTTP start\n")
	defer log.Printf("genAccessTokenByHTTP end\n")

	req, err := http.NewRequest("GET", "https://oapi.dingtalk.com/gettoken", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("appkey", appKey)
	q.Add("appsecret", appSecret)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client do failed, err = %v\n", err)
		return "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})
	if err := json.Unmarshal(resp_body, &m); err != nil {
		return "", err
	}

	return m["access_token"].(string), nil
}

/*
@brief: 通过电话号码获取用户 userID
// 通过 access_token 和用户手机号获取用户 userid
https://open.dingtalk.com/document/orgapp-server/query-users-by-phone-number
{
    "errcode": 0,
    "errmsg": "ok",
    "result": {
        "userid": "xxxx"
    },
    "request_id": "wkyjps1mcm35"
}
*/
func getUserIDByPhone(accessToken string, userPhone string) (string, error) {
	url := "https://oapi.dingtalk.com/topapi/v2/user/getbymobile"

	bodyMap := make(map[string]interface{})
	bodyMap["mobile"] = userPhone
	bodyData, _ := json.Marshal(bodyMap)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})
	if err := json.Unmarshal(resp_body, &m); err != nil {
		log.Printf("body = %v, err = %v", string(resp_body), err)
		return "", err
	}

	resultM := m["result"].(map[string]interface{})
	return resultM["userid"].(string), nil
}

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dingtalkrobot_1_0.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkrobot_1_0.Client{}
	_result, _err = dingtalkrobot_1_0.NewClient(config)
	return _result, _err
}

/*
https://open.dingtalk.com/document/robots/chatbots-send-one-on-one-chat-messages-in-batches
*/
func sendMsg(
	appKey string,
	userID string,
	userName string,
	awardLevel string,
	accessToken string,
) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	log.Printf("create client success\n")

	msgTitle := "【深圳佑驾】"
	msgText := fmt.Sprintf("【深圳佑驾】 恭喜%v获得2022MINIEYE新春年会%v，请尽快联系行政部工作人员领取。",
		userName, awardLevel)
	msgMap := make(map[string]string)
	msgMap["title"] = msgTitle
	msgMap["text"] = msgText
	msgData, _ := json.Marshal(msgMap)

	batchSendOTOHeaders := &dingtalkrobot_1_0.BatchSendOTOHeaders{}
	batchSendOTOHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	batchSendOTORequest := &dingtalkrobot_1_0.BatchSendOTORequest{
		RobotCode: tea.String(appKey),
		UserIds:   []*string{tea.String(userID)},
		MsgKey:    tea.String("sampleMarkdown"),
		MsgParam:  tea.String(string(msgData)),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		log.Printf("client start send msg\n")
		_, _err = client.BatchSendOTOWithOptions(
			batchSendOTORequest,
			batchSendOTOHeaders,
			&util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(err.Code)) && !tea.BoolValue(util.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
			log.Printf("err.Code = %v, err.Message = %v", err.Code, err.Message)
		}
	}

	log.Printf("_main end\n")
	return _err
}

// func main() {
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	if err := SendDingDingMsg(
// 		"dinggoy40j72loamriym",
// 		"UR66MipyIFSRD1vMZv7jB2iNKFVtdaIZ-K1qFP6qPajid5gXyXyqIg6Xs53RSfC0",
// 		"15960389469",
// 		"黄剑",
// 		"特等奖，RTX3090显卡一张",
// 	); err != nil {
// 		panic(err)
// 	}
// }
