package couchbase

import (
	"fmt"

	"github.com/astaxie/beego"

	"gopkg.in/couchbaselabs/gocb.v1"
)

var (
	cbManager *CouchbaseManager
)

func init() {
}

func Service() *CouchbaseManager {
	if cbManager == nil {
		cbManager = NewCouchbaseManager("")
		cbManager.Init()
	}
	fmt.Printf("GetCBManager. \n")
	return cbManager
}

func NewCouchbaseManager(conn string) *CouchbaseManager {
	fmt.Printf("NewCouchbaseManager. \n")
	cb := &CouchbaseManager{}
	var err error
	// couchbase_server
	svr := fmt.Sprintf("couchbase://%s", beego.AppConfig.String("couchbase_server"))
	fmt.Println("couchbase: ", svr)
	cb.Cluster, err = gocb.Connect(fmt.Sprintf("couchbase://%s", beego.AppConfig.String("couchbase_server")))
	if err != nil {
		fmt.Println("couchbase connection error..")
		cb.Cluster = nil
	}
	return cb
}

type CouchbaseManager struct {
	Cluster       *gocb.Cluster
	Bucket        *gocb.Bucket
	BucketReceipt *gocb.Bucket
}

func (c *CouchbaseManager) Init() {
	if c.Cluster != nil {
		c.Bucket, _ = c.Cluster.OpenBucket("jupiterapi_shop", "jupiterapi")
		c.BucketReceipt, _ = c.Cluster.OpenBucket("jupiterapi_shop_receipt", "jupiterapi_shop_receipt")
	}
}

func (c *CouchbaseManager) GetValue(key string) {
	var value interface{}
	cas, _ := c.Bucket.Get(key, &value)
	fmt.Printf("Got value `%+v` with CAS `%08x`\n", value, cas)
}
