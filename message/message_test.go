package message

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestJSON(t *testing.T) {
	resp := QuerySendDetailsResponse{}
	resp.Code = "1223"
	resp.Message = "测试消息"
	resp.RequestID = "请求ID"
	resp.TotalCount = 12
	resp.TotalPage = 2

	dtolist := make([]*SMSSendDetailDTO, 0)
	for i := 0; i < 5; i++ {
		dto := SMSSendDetailDTO{}
		dto.Content = "内容" + strconv.Itoa(i)
		dto.PhoneNum = "18630399986"
		dto.PhoneNum = strconv.Itoa(i)
		sd := time.Now()
		dto.SendDate = GenDatetime(sd)
		dto.ReceiveDate = "2018-04-12 13:36:06"
		dtolist = append(dtolist, &dto)
	}
	dtos := SMSSendDetailDTOs{dtolist}
	resp.SMSSendDetailDTOs = &dtos
	buf, err := json.Marshal(&resp)
	if err != nil {
		t.Errorf("model json maarshal error:%v", err)
		return
	}
	jsonstr := string(buf)
	t.Logf("QuerySendDetailsResponse json: %v", jsonstr)
}
