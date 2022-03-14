package callbacks

import "encoding/xml"

// 自动生成的回调结构,按需修改
// 文档地址
// https://developer.work.weixin.qq.com/document/path/92277#企业客户标签创建事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeExternalTagCreate{})
}

type ThirdChangeExternalTagCreate struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	ID struct {
		Text string `xml:",chardata"`
	} `xml:"Id"`
	TagType struct {
		Text string `xml:",chardata"`
	} `xml:"TagType"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
}

func (ThirdChangeExternalTagCreate) GetMessageType() string {
	return "third"
}

func (ThirdChangeExternalTagCreate) GetEventType() string {
	return "change_external_tag"
}

func (ThirdChangeExternalTagCreate) GetChangeType() string {
	return "create"
}

func (m ThirdChangeExternalTagCreate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeExternalTagCreate) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeExternalTagCreate
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
