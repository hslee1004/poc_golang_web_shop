package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/couchbaselabs/gocb.v1"
)

const (
	CASHTYPE_ALL          = ""
	CASHTYPE_CREDIT_ONLY  = "credit_only"
	CASHTYPE_PREPAID_ONLY = "prepaid_only"
)

func (u *Invoice) GetRuleId() string {
	switch u.OptionCashType {
	case CASHTYPE_ALL:
		return "WSR1"
	case CASHTYPE_CREDIT_ONLY:
		return "WSR4"
	case CASHTYPE_PREPAID_ONLY:
		return "WSR2"
	}
	return "WSR1"
}

func NewInvoice(pid string, uno string) *Invoice {
	cb := &Invoice{ProductId: pid, UserNo: uno}
	return cb
}

type Invoice struct {
	ProductId            string          `json:"product_id"`
	InvoiceId            string          `json:"invoice_id"` // new
	UserNo               string          `json:"user_no"`
	UserId               string          `json:"user_id"`
	UserIP               string          `json:"user_ip"`
	Date                 string          `json:"date"`
	Items                []Item          `json:"items"`
	TotalPrice           float32         `json:"total_price"`
	Ticket               string          `json:"ticket"`
	Token                string          `json:"access_token"`
	RuleId               string          `json:"rule_id"`
	RedirectUri          string          `json:"redirect_uri"`
	AllowedPaymentMethod []PaymentMethod `json:"allowed_payment_method"` // new definition : nx_prepaid, nx_credit
	OptionCashType       string          `json:"option_cash_type"`       // depricated, empty is all, credit_only, prepaid_only
	OptionUseReceiptFlow bool            `json:"option_use_receipt_flow"`
	TransactionId        string          `json:"transaction_id"`     // of cashbroker
	Balances             []WalletBalance `json:"balances,omitempty"` // balances
}

type InvoiceResponse struct {
	Success *InvoiceResponseDetail `json:"success,omitempty"`
	Error   *NXAPIErrorDetail      `json:"error,omitempty"`
}

type InvoiceResponseDetail struct {
	Code    int     `json:"code"`
	Invoice Invoice `json:"data,omitempty"`
}

func (u *Invoice) RegisterInvoice(bucket *gocb.Bucket) error {
	if bucket == nil {
		return errors.New("system error - saving invoice.")
	} else {
		exp, _ := beego.AppConfig.Int64("couchbase_expiry")
		_, err := bucket.Insert(u.Ticket, u, uint32(exp)) // sec : 10 mins
		return err
	}
}

func (u *Invoice) GetInvoice(bucket *gocb.Bucket) bool {
	fmt.Println("GetInvoice: key:", u.Ticket)
	if u.Ticket == "" {
		return false
	}
	if cas, err := bucket.Get(u.Ticket, &u); err == nil {
		fmt.Printf("%v, case:%v\n", u, cas)
		return true
	} else {
		fmt.Printf("%v, case:%v\n", u, cas)
	}
	return false
}

func (u *Invoice) Save(bucket *gocb.Bucket) error {
	if bucket == nil {
		return errors.New("system error - saving invoice.")
	} else {
		exp, _ := beego.AppConfig.Int64("couchbase_expiry")
		_, err := bucket.Upsert(u.Ticket, u, uint32(exp)) // sec
		return err
	}
}

func (u *Invoice) Verify() bool {
	if u.ProductId == "" {
		return false
	}
	return true
}
