package send

import (
	"HFish/error"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
type AlterOverModel struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
func SendAlterOver(findAlterConfig []map[string]interface{},args ...string)interface{}{
	infoArr:=strings.Split(findAlterConfig[0]["info"].(string),"&&")
	resp, err := http.PostForm("https://api.alertover.com/v1/alert",
		url.Values{
		"source": {infoArr[0]},
		"receiver": {infoArr[1]},
		"title":{args[0]},
		"content":{args[1]},
	})
	if err != nil {
		error.Check(err, "发送alterOver通知失败1")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		error.Check(err, "发送alterOver通知失败2")
	}
	_alterOverModel:= AlterOverModel{}
	result :=json.Unmarshal(body,&_alterOverModel)
	return result
}
