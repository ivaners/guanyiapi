package common

type FetchXml struct {
	Total  float64 `xml:"total_results"`
	Trades Trades  `xml:"trades"`
}

type Trades struct {
	Trade []Trade `xml:"trade"`
}

type Trade struct {
	OrderId       string `xml:"lydh"`
	PayAmonnt     string `xml:"zje"`
	OrderTime     string `xml:"rq"`
	ShopName      string `xml:"shopcode"`
	ShippingType  string `xml:"wlgsmc"`
	Zf            string `xml:"zf"`
	WarehouseType string `xml:"ckmc"`
	Remark        string `xml:"sellernote"`
	TbRemark      string `xml:"tb_bz"`
	UpdateTime    string `xml:"modify_time"`
	PayTime       string `xml:"paytime"`
}
