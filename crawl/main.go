package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/imroc/req/v3"
	_ "github.com/imroc/req/v3"
	_ "github.com/levigross/grequests"
	"github.com/spf13/cast"
	"golang.org/x/net/html"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

var (
	fileTypes map[string]bool
	quit chan struct{}
	ch chan *Infos
)
func init(){
	fileTypes=map[string]bool{"[电影]":true,"[电视剧]":true}
	quit=make(chan struct{})
	ch=make(chan *Infos,10)
}

type Infos struct {
	gorm.Model
	FilmStyle string `gorm:"column:filmstyle`
	FilmName string	`gorm:"column:filmname`
	Url string	`gorm:"column:url`
	Score float32	`gorm:"column:score`
	Introduction string	`gorm:"column:introduction`
	Actors string	`gorm:"column:actors`
	Year int	`gorm:"column:year`
}

type ResJson struct {
	Items []string
	Limit int32
	More bool
	Total int32
}

func fetch(keyword string,start int) ResJson  {
	url:="https://www.douban.com/j/search"

	var resJson ResJson
	client:=req.C()

	_,err:=client.R().
		SetHeader("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36").
		SetHeader("Accept", "*/*").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9").
		SetHeader("Referer", "https://www.douban.com/search?q=%E5%BC%A0%E8%AF%91").
		SetQueryParam("q",keyword).
		SetQueryParam("start",cast.ToString(start)).
		SetQueryParam("subtype","item").
		SetResult(&resJson).
		Get(url)
	if err!=nil {
		panic(err)
	}
	return resJson
}
func extractText(node *html.Node,xpath string) string{
	subnode:=htmlquery.FindOne(node,xpath)
	if subnode!=nil{
		return htmlquery.InnerText(subnode)
	}
	return ""
}

func extract(html string) (Infos,bool) {
	selector,err:=htmlquery.Parse(strings.NewReader(html))
	if err!=nil{
		panic(err)
	}
	contentNode:=htmlquery.FindOne(selector,".//div[@class='result']/div[@class='content']")
	introduction:= extractText(contentNode,"./p/text()")
	title:=htmlquery.FindOne(contentNode,"./div[@class='title']")
	filmStyle:=extractText(title,"./h3//span")  //?
	if _,ok:=fileTypes[filmStyle];ok{
		filmName:=extractText(title,"./h3//a")
		url:=extractText(title,"./h3//a/@href")
		score:= extractText(title,"./div[@class='rating-info']/span[@class='rating_nums']")
		actorsInfos:=extractText(title,"./div[@class='rating-info']/span[@class='subject-cast']")
		actorAndOtherInfo:=strings.Split(actorsInfos,"/")
		year:=strings.TrimSpace(actorAndOtherInfo[len(actorAndOtherInfo)-1])
		actors:=actorAndOtherInfo[1:len(actorAndOtherInfo)-1]
		return Infos{FilmStyle:filmStyle,FilmName:filmName,Url:url,Score:cast.ToFloat32(score),Introduction:introduction,Actors:strings.Join(actors,","),Year:cast.ToInt(year)},ok
	}else{
		return Infos{},ok
	}

}

func getItems(keyword string) {
	 wg:= sync.WaitGroup{}
	for start:=0;start<=550;start=start+20{
			go func(i int){
				wg.Add(1)
				resJson:=fetch(keyword,i)
				//extract
				for _,item := range resJson.Items{
					if infos,ok:=extract(item);ok{
						fmt.Printf("[Query] %s %d\n",infos.FilmName,i)
						//to channel
						ch<-&infos
					}
				}
				wg.Done()
			}(start)
	}
	wg.Wait()
	close(ch)
}

func save(){
	dns:="root:root@tcp(127.0.0.1:3306)/douban?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dns),&gorm.Config{})
	if err!=nil{
		panic(err)
	}
	db.AutoMigrate(&Infos{})
	//db.Delete(&Infos{})
	if err!=nil{
		panic(err)
	}
	for{
		if infos,ok:=<-ch;ok{
			db.Create(infos)
			fmt.Printf("[Insert] %s\n",infos.FilmName)
		}else{
			fmt.Println("[Closed] Closed ch")
			break
		}
	}
	quit<-struct{}{}
}
func main() {
	startTime:=time.Now().Unix()
	var keyword string
	fmt.Scanln(&keyword)
	fmt.Println(keyword)
	go getItems(keyword)
	go save()
	<-quit
	fmt.Printf("[Cost] %d",time.Now().Unix()-startTime)
}
