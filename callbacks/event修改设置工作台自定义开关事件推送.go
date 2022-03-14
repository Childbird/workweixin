package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/94620#修改设置工作台自定义开关事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventSwitchWorkbenchMode{})
}

type EventSwitchWorkbenchMode struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName struct {
		Text string `xml:",chardata"`
	} `xml:"ToUserName"`
	FromUserName struct {
		Text string `xml:",chardata"`
	} `xml:"FromUserName"`
	CreateTime struct {
		Text string `xml:",chardata"`
	} `xml:"CreateTime"`
	MsgType struct {
		Text string `xml:",chardata"`
	} `xml:"MsgType"`
	Event struct {
		Text string `xml:",chardata"`
	} `xml:"Event"`
	Mode struct {
		Text string `xml:",chardata"`
	} `xml:"Mode"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (EventSwitchWorkbenchMode) GetMessageType() string {
	return "event"
}

func (EventSwitchWorkbenchMode) GetEventType() string {
	return "switch_workbench_mode"
}

func (EventSwitchWorkbenchMode) GetChangeType() string {
	return ""
}

func (m EventSwitchWorkbenchMode) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventSwitchWorkbenchMode) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventSwitchWorkbenchMode
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
