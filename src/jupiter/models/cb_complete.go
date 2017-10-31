package models

import (
	"encoding/xml"
	"fmt"
	"jupiter/libs"
)

const XML_COMPLETE_PAYMENT_ATTR = "Nexon.CashBroker.Universal.Entity.CompletePaymentRequest, Nexon.CashBroker.Universal, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null"

type CompletePaymentSection struct {
	XMLName  xml.Name                 `xml:"xmlCashBrokerSection"`
	PropType string                   `xml:"type,attr"`
	Request  CompletePaymentRequest   `xml:"CommitPaymentRequest,omitempty"`
	Response *CompletePaymentResponse `xml:"ommitPaymentResponse,omitempty"`
}

type CompletePaymentRequest struct {
	MethodName    string `xml:"MethodName,attr"`
	Code          string `xml:"Code"`
	UserId        string `xml:"UserId"`
	RuleId        string `xml:"RuleId"`
	TransactionId string `xml:"TransactionId"`
	RemoteIp      string `xml:"RemoteIp"`
	RequestDate   string `xml:"RequestDate"`
}

type CompletePaymentResponse struct {
	Result    string `xml:"Result,omitempty"`
	PaymentNo string `xml:"PaymentNo,omitempty"`
}

func NewCompletePayReq(trxId string, v *Invoice) *CompletePaymentSection {
	//	var items []Item	//fixing
	//	libs.JsonUnmarshal([]byte(v.Items), &items)
	//	fmt.Printf("%v", items)
	s := &CompletePaymentSection{
		PropType: XML_COMMITPAYMENT_ATTR,
		Request: CompletePaymentRequest{
			MethodName:    "CommitPayment",
			Code:          "30205",
			UserId:        v.UserId,
			RuleId:        v.GetRuleId(), // WSR2:prepaid, WSR4
			RemoteIp:      v.UserIP,
			RequestDate:   v.Date,
			TransactionId: trxId,
		},
	}
	return s
}

func (t *CompletePaymentSection) XMLDecode(bytes []byte, out interface{}) error {
	return xml.Unmarshal([]byte(libs.ReplaceXMLHeader(string(bytes))), &out)
}

func (t *CompletePaymentSection) XMLEncode(value interface{}) ([]byte, error) {
	s, err := xml.MarshalIndent(t, "  ", "    ")
	bytes := []byte(xml.Header + string(s))
	fmt.Printf("XMLEncode request:%s", bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
