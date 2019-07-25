package send

import (
	"HFish/error"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*发送pushBullet推送*/
func SendPushBullet(findAlterConfig []map[string]interface{},args ...string)interface{}{
	infoArr:=strings.Split(findAlterConfig[0]["info"].(string),"&&")
	fmt.Println(infoArr[0],args[0],args[1])
	var req *http.Request
	body:= map[string]string{
		"title":args[0],
		"body":args[1],
		"type":"note",  //默认消息类型
	}
	headers:=map[string]string{
		"Access-Token":infoArr[0],
		"Content-Type":"application/json",
	}
	url:="https://api.pushbullet.com/v2/pushes"
	bodyJson, _ := json.Marshal(body)
	req,err:= http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err!=nil{
		error.Check(err,"发送pushBullet创建req报错")
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	result,err :=client.Do(req)
	if err!=nil{
		error.Check(err,"发送pushBullet返回resp报错")
	}
	return result
}
