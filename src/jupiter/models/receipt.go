package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/couchbaselabs/gocb.v1"
)

type Receipt struct {
	ReceiptId string   `json:"receipt_id"`
	Invoice   *Invoice `json:"invoice"`
}

func (r *Receipt) Get(bucket *gocb.Bucket) bool {
	fmt.Println("Get receipt: key:", r.ReceiptId)
	if r.ReceiptId != "" {
		cas, err := bucket.Get(r.ReceiptId, &r.Invoice)
		fmt.Printf("invoice:%v, case:%v\n", r.Invoice, cas)
		if err == nil {
			return true
		}
	}
	return false
}

func (r *Receipt) Save(bucket *gocb.Bucket) error {
	if bucket == nil {
		return errors.New("system error - saving receipt.")
	} else {
		exp, _ := beego.AppConfig.Int64("couchbase_expiry")
		_, err := bucket.Upsert(r.ReceiptId, r.Invoice, uint32(exp))
		return err
	}
}

func (r *Receipt) Verify() bool {
	if r.Invoice.TransactionId == "" {
		return false
	}
	return true
}

func (r *Receipt) GetReceiptAPIURL() string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("api_host"), beego.AppConfig.String("api_endpoint_receipt"))
}
