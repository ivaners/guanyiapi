package main

import (
	"database/sql"
	"encoding/xml"
	"flag"
	"fmt"
	_ "github.com/mattn/go-adodb"
	"guanyiapi/common"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var Input_page = flag.String("page", "1", "input ur page")

var Input_page_size = flag.String("page_size", "1000", "input ur page_size")

func main() {
	flag.Parse()

	beginTime := time.Now().Unix() - 864000
	beginFormat := time.Unix(beginTime, 0).Format("2006-01-02")

	v := url.Values{}
	v.Add("method", "ecerp.trade.get")
	v.Add("appkey", "9F4C2B07E0334C9294Bee3321924AE266C")
	v.Add("fields", "lydh,rq,shopcode,zje,wlgsmc,ckmc,paytime,sellernote,tb_bz,modify_time,zf")
	v.Add("page_no", *Input_page)
	v.Add("page_size", *Input_page_size)
	v.Add("orderby", "paytime")
	v.Add("orderbytype", "ASC")
	v.Add("condition", fmt.Sprintf(" paytime >='%s' and paytime <'%s' ", beginFormat, time.Now().Format("2006-01-02")))
	response, _ := http.Get("http://121.199.166.123:30002/data.dpk?" + v.Encode())
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	data := common.FetchXml{}
	err := xml.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	db, err := sql.Open("adodb", "Provider=SQLOLEDB;Initial Catalog=fengmingdw;Data Source=192.168.6.29;user id=ods_taobao;password=FMbi123")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStr := "INSERT INTO ods_taobao_order (order_id, order_time, shop_name, pay_amonnt, pay_time, shipping_type, warehouse_type, remark, tb_remark, update_time, zf) VALUES "
	for _, item := range data.Trades.Trade {
		query := sqlStr + fmt.Sprintf("('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", item.OrderId, strings.Replace(item.OrderTime, "t", " ", 1), item.ShopName,
			item.PayAmonnt, strings.Replace(item.PayTime, "t", " ", 1), item.ShippingType, item.WarehouseType, item.Remark, item.TbRemark, strings.Replace(item.UpdateTime, "t", " ", 1), item.Zf)
		db.Exec(query)
	}
}
