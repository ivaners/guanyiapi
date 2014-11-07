package main

import (
	"encoding/xml"
	"fmt"
	"guanyiapi/common"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	// "os"
	"os/exec"
	// "path/filepath"
	"strconv"
	"time"
)

func main() {
	beginTime := time.Now().Unix() - 864000
	beginFormat := time.Unix(beginTime, 0).Format("2006-01-02")

	v := url.Values{}
	v.Add("method", "ecerp.trade.get")
	v.Add("appkey", "9F4C2B07E0334C9294B2F122333266C")
	v.Add("fields", "lydh,rq,shopcode,zje,wlgsmc,ckmc,paytime,sellernote,tb_bz,modify_time,zf")
	v.Add("page_no", "1")
	v.Add("page_size", "1")
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
		return
	}

	var pageSize float64 = 1000
	pageCount := math.Ceil(data.Total / pageSize)

	var i float64 = 1

	for ; i < pageCount; i++ {
		cmd := exec.Command("sub", "-page="+strconv.FormatFloat(i, 'f', 0, 64), "-page_size="+strconv.FormatFloat(pageSize, 'f', 0, 64))
		cmd.Output()
		fmt.Println(i)
	}
}
