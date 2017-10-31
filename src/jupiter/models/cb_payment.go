package models

import (
	"encoding/xml"
	"fmt"
	"jupiter/libs"
)

const XML_PAYMENT_ATTR = "Nexon.CashBroker.Universal.Entity.RequestPaymentRequest, Nexon.CashBroker.Universal, Version=1.0.0.0, Culture=neutral, PublicKeyToken=null"

//type RequestPaymentCashBrokerSection struct {
type RequestPaymentSection struct {
	XMLName  xml.Name                `xml:"xmlCashBrokerSection"`
	PropType string                  `xml:"type,attr"`
	Request  RequestPaymentRequest   `xml:"RequestPaymentRequest,omitempty"`
	Response *RequestPaymentResponse `xml:"RequestPaymentResponse,omitempty"`
}

type RequestPaymentRequest struct {
	MethodName  string   `xml:"MethodName,attr"`
	Code        string   `xml:"Code"`
	UserId      string   `xml:"UserId"`
	OrderID     string   `xml:"OrderID"`
	RuleId      string   `xml:"RuleId"`
	ServiceCode string   `xml:"ServiceCode"`
	CharacterId string   `xml:"CharacterId"`
	RemoteIp    string   `xml:"RemoteIp"`
	RequestDate string   `xml:"RequestDate"`
	ItemList    ItemList `xml:"ItemList"`
}

type RequestPaymentResponse struct {
	Result        string `xml:"Result,omitempty"`
	TransactionId string `xml:"TransactionId,omitempty"`
}

func NewPayReq(serviceCode string, v *Invoice) *RequestPaymentSection {
	//	var items []Item	//fixing
	//	libs.JsonUnmarshal([]byte(v.Items), &items)
	//	fmt.Printf("%v", items)
	s := &RequestPaymentSection{
		PropType: XML_PAYMENT_ATTR,
		Request: RequestPaymentRequest{
			MethodName: "RequestPayment",
			Code:       "30202",
			UserId:     v.UserId,
			OrderID:    v.ProductId,
			//RuleId:      v.RuleId,
			RuleId:      v.GetRuleId(), // WSR2:prepaid, WSR4
			ServiceCode: serviceCode,
			CharacterId: v.UserNo,
			RemoteIp:    v.UserIP,
			RequestDate: v.Date,
			ItemList:    ItemList{ItemList: v.Items},
		},
		Response: nil,
	}
	//ItemList:    items,  //fixing
	return s
}

//type ItemList struct {
//	ItemList []Item `xml:"ItemList"`
//}

func (t *RequestPaymentSection) XMLDecode(bytes []byte, out interface{}) error {
	return xml.Unmarshal([]byte(libs.ReplaceXMLHeader(string(bytes))), &out)
}

func (t *RequestPaymentSection) XMLEncode(value interface{}) ([]byte, error) {
	s, err := xml.MarshalIndent(t, "  ", "    ")
	bytes := []byte(xml.Header + string(s))
	fmt.Printf("XMLEncode request:%s\n", bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

/*
 <RequestPaymentRequest MethodName="RequestPayment">
    <Code>30202</Code>
    <UserId>Nxtest61</UserId>
    <OrderID>TEST</OrderID>
    <RuleId>WSR1</RuleId>
    <ServiceCode>SVG035</ServiceCode>
    <CharacterId>koreaysm</CharacterId>
    <RemoteIp>192.168.1.234</RemoteIp>
    <RequestDate>2008-01-02 1:33:23</RequestDate>
    <ItemList>
      <Item>
        <ItemId>12345</ItemId>
        <ItemName>test</ItemName>
        <Price>100</Price>
        <Quantity>1</Quantity>
      </Item>
      <Item>
        <ItemId>TEST2</ItemId>
        <ItemName>test2</ItemName>
        <Price>1000</Price>
        <Quantity>1</Quantity>
      </Item>
    </ItemList>
  </RequestPaymentRequest>
</xmlCashBrokerSection>

<RequestPaymentResponse>
	<Result>50000</Result>
	<TransactionId>43952FD2FEC5450AA9AE0124C9D538BE</TransactionId>
</RequestPaymentResponse>
*/
