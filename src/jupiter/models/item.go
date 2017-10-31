package models

//type Item struct {
//	Id    string  `json:"id", xml:"ItemId"`
//	Name  string  `json:"name", xml:"ItemName"`
//	Price float32 `json:"price, xml:"Price"`
//	Qty   int     `json:"qty", xml:"Quantity`
//}

type Item struct {
	Id    string  `json:"id" xml:"ItemId"`
	Name  string  `json:"name" xml:"ItemName"`
	Price float32 `json:"price" xml:"Price"`
	Qty   int     `json:"qty" xml:"Quantity"`
}

type ItemList struct {
	ItemList []Item `xml:"Item"`
}

func (i *Item) Sum(total *float32) float32 {
	*total = *total + i.Price
	return *total
}

// for test
func GetTestItems() []Item {
	items := [...]Item{
		{Id: "10001", Name: "test_jupiter_item_1", Price: 1, Qty: 1},
		{Id: "10002", Name: "test_jupiter_item_2", Price: 2, Qty: 1},
		{Id: "10003", Name: "test_jupiter_item_3", Price: 3, Qty: 1},
		{Id: "10004", Name: "test_jupiter_item_4", Price: 4, Qty: 1},
		{Id: "10005", Name: "test_jupiter_item_5", Price: 5, Qty: 1},
	}
	return items[:]
}

func GetSelectedTestItems(items []string) []Item {
	rs := []Item{}
	for _, item := range GetTestItems() {
		if stringInSlice(item.Id, items) {
			rs = append(rs, item)
		}
	}
	return rs[:]
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
