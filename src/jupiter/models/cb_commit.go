package models

import (
	"encoding/xml"
	"fmt"
	"jupiter/libs"
)

const XML_COMMITPAYMENT_ATTR = "Nexon.CashBroker.Universal.Entity.CommitPaymentRequest, Nexon.CashBroker.Universal, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null"

type CommitPaymentSection struct {
	XMLName  xml.Name               `xml:"xmlCashBrokerSection"`
	PropType string                 `xml:"type,attr"`
	Request  CommitPaymentRequest   `xml:"CommitPaymentRequest,omitempty"`
	Response *CommitPaymentResponse `xml:"CommitPaymentResponse,omitempty"`
}

type CommitPaymentRequest struct {
	MethodName    string `xml:"MethodName,attr"`
	Code          string `xml:"Code"`
	UserId        string `xml:"UserId"`
	RuleId        string `xml:"RuleId"`
	TransactionId string `xml:"TransactionId"`
	RemoteIp      string `xml:"RemoteIp"`
	RequestDate   string `xml:"RequestDate"`
}

type CommitPaymentResponse struct {
	Result        string `xml:"Result,omitempty"`
	DetailMessage string `xml:"DetailMessage,omitempty"`
	PaymentNo     string `xml:"PaymentNo,omitempty"`
}

func NewCommitPayReq(trxId string, v *Invoice) *CommitPaymentSection {
	//	var items []Item      //fixing
	//	libs.JsonUnmarshal([]byte(v.Items), &items)
	//	fmt.Printf("%v", items)
	s := &CommitPaymentSection{
		PropType: XML_COMMITPAYMENT_ATTR,
		Request: CommitPaymentRequest{
			MethodName:    "CommitPayment",
			Code:          "30203",
			UserId:        v.UserId,
			RuleId:        v.GetRuleId(), // WSR2:prepaid, WSR4
			RemoteIp:      v.UserIP,
			RequestDate:   v.Date,
			TransactionId: trxId,
		},
		Response: nil,
	}
	return s
}

func (t *CommitPaymentSection) XMLDecode(bytes []byte, out interface{}) error {
	return xml.Unmarshal([]byte(libs.ReplaceXMLHeader(string(bytes))), &out)
}

func (t *CommitPaymentSection) XMLEncode(value interface{}) ([]byte, error) {
	s, err := xml.MarshalIndent(t, "  ", "    ")
	bytes := []byte(xml.Header + string(s))
	fmt.Printf("XMLEncode request:%s\n", bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
