package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
)

type Blog struct {
	service string
	Url     string
}

func GetBlog(url string) Blog {
	if strings.Contains(url, "naver") {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("User-Agent", "Android")
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		reg, _ := regexp.Compile(`top\.location\.replace\('(.*)'\);`)
		found := reg.FindSubmatch(body)
		mUrl := strings.ReplaceAll(string(found[1]), `\/`, "/")

		if strings.Contains(mUrl, "PostList.nhn") {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			reg, _ := regexp.Compile(`blogId=([a-z0-9_-]{5,20})&`)
			found := reg.FindStringSubmatch(mUrl)
			blogId := found[1]
			reg, _ = regexp.Compile(fmt.Sprintf(`url_%s_([0-9]{12})`, blogId))
			found = reg.FindStringSubmatch(string(body))
			logNo := found[1]

			mUrl = fmt.Sprintf("http://m.blog.naver.com/%s/%s", blogId, logNo)
		}

		return Blog{"naver", mUrl}
	} else {
		return Blog{"unknown", url}
	}
}

func (b Blog) GetSub() []string {
    downlodable := []string{}

	switch b.service {
	case "naver":
        doc, _ := htmlquery.LoadURL(b.Url)
		found := htmlquery.Find(doc, `//a[contains(@href,"blogattach")]/@href`)

        for _, n := range found {
            downlodable = append(downlodable, htmlquery.InnerText(n))
        }
	default:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header.Add("User-Agent", "Android")
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

        doc, _ := htmlquery.Parse(resp.Body)

        links := []string{"tistory.com/attachment", "egloos.com/pds", "drive.google.com/uc"}

        for _, link := range links {
            found := htmlquery.Find(doc, fmt.Sprintf(`//a[contains(@href,"%s")]/@href`, link))
            for _, n := range found {
                downlodable = append(downlodable, htmlquery.InnerText(n))
            }
        }
	}

    return downlodable
}

func main() {
	urls := []string{
		"http://blog.noitamina.moe/221528410736",
		"http://blog.naver.com/cobb333/221391135993",
		"http://blog.naver.com/PostList.nhn?blogId=harne_&categoryNo=260&from=postList",
		"http://melody88.tistory.com/631",
		"https://mihorima.blogspot.com/2018/11/05_4.html",
		"http://godsungin.tistory.com/200",
		"https://blog.naver.com/qtr01122/221391146050",
	}

	for _, url := range urls {
		fmt.Println(GetBlog(url).GetSub())
	}
}
