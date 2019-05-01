package main

import (
	"fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "regexp"

    "github.com/antchfx/htmlquery"
)

func get_service(url string) string {
    if strings.Contains(url, "naver") {
        return "naver"
    } else {
        return "unknown"
    }
}

func get_sub(url string) {
	service := get_service(url)

	switch service {
	case "naver":
        req, _ := http.NewRequest("GET", url, nil)
        req.Header.Add("User-Agent", "Android")
        resp, _ := (&http.Client{}).Do(req)
        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)

        reg, _ := regexp.Compile(`top\.location\.replace\('(.*)'\);`)
        found := reg.FindSubmatch(body)
        m_url := strings.ReplaceAll(string(found[1]), `\/`, "/")

        if strings.Contains(m_url, "PostList.nhn") {
            resp, err := http.Get(url)
            if err != nil {
                panic(err)
            }
            defer resp.Body.Close()

            body, _ := ioutil.ReadAll(resp.Body)

            reg, _ := regexp.Compile(`blogId=([a-z0-9_-]{5,20})&`)
            found := reg.FindStringSubmatch(m_url)
            blog_id := found[1]
            reg, _ = regexp.Compile(fmt.Sprintf(`url_%s_([0-9]{12})`, blog_id))
            found = reg.FindStringSubmatch(string(body))
            log_no := found[1]

            m_url = fmt.Sprintf("http://m.blog.naver.com/%s/%s", blog_id, log_no)
        }

        resp, _ = http.Get(m_url)
        defer resp.Body.Close()

        body, _ = ioutil.ReadAll(resp.Body)

        doc, _ := htmlquery.LoadURL(m_url)
        list := htmlquery.Find(doc, `//a[contains(@href,"blogattach")]/@href`)

        for _, n := range list {
            fmt.Println(htmlquery.InnerText(n))
        }
    default:
        fmt.Println(url)
	}
}

func main() {
    tests := []string{
		"http://blog.noitamina.moe/221391147667",
		"http://blog.naver.com/cobb333/221391135993",
		"http://blog.naver.com/PostList.nhn?blogId=harne_&categoryNo=260&from=postList",
		"http://melody88.tistory.com/631",
		"https://mihorima.blogspot.com/2018/11/05_4.html",
		"http://godsungin.tistory.com/200",
		"https://blog.naver.com/qtr01122/221391146050",
	}

    for _, e := range tests {
        get_sub(e)
    }
}
