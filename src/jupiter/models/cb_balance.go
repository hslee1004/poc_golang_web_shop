package models

import (
	"encoding/xml"
	"fmt"
	"jupiter/libs"

	"github.com/astaxie/beego"
)

const check_balance_type = "Nexon.CashBroker.Universal.Entity.CheckBalanceRequest, Nexon.CashBroker.Universal, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null"

type XmlCashBrokerSection struct {
	PropType             string               `xml:"type,attr"`
	CheckBalanceRequest  CheckBalanceRequest  `xml:"CheckBalanceRequest,omitempty"`
	CheckBalanceResponse CheckBalanceResponse `xml:"CheckBalanceResponse,omitempty"`
}

type CheckBalanceRequest struct {
	MethodName string `xml:"MethodName,attr"`
	Code       string `xml:"Code"`
	UserId     string `xml:"UserId"`
	RuleId     string `xml:"RuleId"`
}

type CheckBalanceResponse struct {
	Result        string  `xml:"Result"`
	DetailMessage string  `xml:"DetailMessage"`
	Balance       float32 `xml:"Balance"`
	BalanceByRule float32 `xml:"BalanceByRule"`
}

func NewBalanceReq(user_id string) *XmlCashBrokerSection {
	cs := &XmlCashBrokerSection{
		PropType: check_balance_type,
		CheckBalanceRequest: CheckBalanceRequest{
			MethodName: "CheckBalance",
			Code:       "30201", // check balance
			UserId:     user_id,
			RuleId:     beego.AppConfig.String("ruleid_nx_prepaid"),
		},
	}
	return cs
}

func (t *XmlCashBrokerSection) XMXDecode(bytes []byte, out interface{}) error {
	return xml.Unmarshal([]byte(libs.ReplaceXMLHeader(string(bytes))), &out)
}

func (t *XmlCashBrokerSection) XMLEncode(value interface{}) ([]byte, error) {
	s, err := xml.MarshalIndent(t, "  ", "    ")
	bytes := []byte(xml.Header + string(s))
	fmt.Printf("request:%s", bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
