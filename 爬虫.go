package main
func main(){
	res,err := http.Get("https://blog.lenconda.top/")
	if err != nil {
		panic(err)
	}
	defer res.body.Close()

	d,err := goquery.NewDocumentFromReader(res.body)
	if err != nil {
		panic(err)
	}
	ch1_title := make(chan,string,600)
	ch2_time := make(chan,string,600)
	ch3_tag := make(chan,string,600)
	ch4_outline := make(chan,string,600)
	d.Find("div.container article div.post-box>h2").Each(func(i int,s *goquery.selection){
		fmt.println("标题："+s.Text())
		ch1_title <- s.Text()
	})
	d.Find("div.container article div.post-box>span:nth-child(1)").Each(func(i int,s *goquery.selection){
		fmt.println("发表时间："+s.Text())
		ch2_time <- s.Text()
	})
	d.Find("div.container article div.post-box>span:nth-child(2)").Each(func(i int,s *goquery.selection){
		fmt.println("标签："+s.Text())
		ch3_tag <- s.Text()
	})
	d.Find("div.container article div.post-box>p").Each(func(i int,s *goquery.selection){
		fmt.println("概要："+s.Text())
		ch4_outline <- s.Text()
	})
	Close(ch1_title)
	Close(ch2_time)
	Close(ch3_tag)
	Close(ch4_outline)
	f,err := os.Create("end.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i <= 600;i++{
		f.WriteString(<-ch1_title +"\t" + <-ch2_time + "\t\t" + <-ch3_tag + "\t" + "\r\n" + <-ch4_outline + "\t\t\t\t\t")
	}
}