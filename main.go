package main

import "fmt"

func main() {
	setting := ReadSetting()
	fmt.Println(setting)
	// urls := []string{
	// 	"http://blog.noitamina.moe/221528410736",
	// 	"http://blog.naver.com/cobb333/221391135993",
	// 	"http://blog.naver.com/PostList.nhn?blogId=harne_&categoryNo=260&from=postList",
	// 	"http://melody88.tistory.com/631",
	// 	"https://mihorima.blogspot.com/2018/11/05_4.html",
	// 	"http://godsungin.tistory.com/200",
	// 	"https://blog.naver.com/qtr01122/221391146050",
	// }

	// for _, url := range urls {
	// 	fmt.Println(GetBlog(url).GetSub())
	// }
}
