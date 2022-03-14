package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://developer.work.weixin.qq.com/document/path/90376#弹出地理位置选择器的事件推送

func init() {
	// 添加可解析的回调事件
	supportCallback(EventLocationSelect{})
}

type EventLocationSelect struct {
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
	EventKey struct {
		Text string `xml:",chardata"`
	} `xml:"EventKey"`
	SendLocationInfo struct {
		Text      string `xml:",chardata"`
		LocationX struct {
			Text string `xml:",chardata"`
		} `xml:"Location_X"`
		LocationY struct {
			Text string `xml:",chardata"`
		} `xml:"Location_Y"`
		Scale struct {
			Text string `xml:",chardata"`
		} `xml:"Scale"`
		Label struct {
			Text string `xml:",chardata"`
		} `xml:"Label"`
		Poiname struct {
			Text string `xml:",chardata"`
		} `xml:"Poiname"`
	} `xml:"SendLocationInfo"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
	AppType struct {
		Text string `xml:",chardata"`
	} `xml:"AppType"`
}

func (EventLocationSelect) GetMessageType() string {
	return "event"
}

func (EventLocationSelect) GetEventType() string {
	return "location_select"
}

func (EventLocationSelect) GetChangeType() string {
	return ""
}

func (m EventLocationSelect) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventLocationSelect) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp EventLocationSelect
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
