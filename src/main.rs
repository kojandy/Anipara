use regex::Regex;
use rss::Channel;
use scraper::Html;
use scraper::Selector;
use serde::Deserialize;
use std::fs::File;
use std::io::BufReader;

#[derive(Deserialize, Debug)]
struct Source {
    name: String,
    rss: String,
    pattern: String,
}

impl Source {
    fn get_links(&self, name: &str) -> Vec<String> {
        let chan = Channel::from_url(&self.rss).unwrap();
        let links: Vec<String> = vec![];

        for item in chan.items() {
            if let Some(title) = item.title() {
                if title.to_lowercase().contains(name) {
                    println!("{}", title);
                }
            }
        }

        links
    }
}

#[derive(Deserialize, Debug)]
struct SubtitleSource {
    anissia: i32,
    author: String,
}

#[derive(Deserialize, Debug)]
struct Subscription {
    source: String,
    subtitle: SubtitleSource,
    title: String,
}

#[derive(Deserialize, Debug)]
struct Settings {
    source: Vec<Source>,
    subscribe: Vec<Subscription>,
}

impl Settings {
    fn parse_yml(path: &str) -> Settings {
        let file = File::open(path).unwrap();
        let reader = BufReader::new(file);

        serde_yaml::from_reader(reader).unwrap()
    }
}

#[derive(Debug)]
enum Blog {
    Naver(String),
    Unknown(String),
}

impl Blog {
    fn from_url(url: &str) -> Blog {
        if url.contains("naver") {
            let body = reqwest::Client::new()
                .get(url)
                .header("User-Agent", "Android")
                .send()
                .unwrap()
                .text()
                .unwrap();
            let m_url = Regex::new(r"top\.location\.replace\('(.*)'\);")
                .unwrap()
                .captures(&body)
                .unwrap()
                .get(1)
                .unwrap()
                .as_str()
                .replace(r"\/", "/");

            if m_url.contains("PostList.nhn") {
                let body = reqwest::get(url).unwrap().text().unwrap();
                let blog_id = Regex::new(r"blogId=([a-z0-9_-]{5,20})&")
                    .unwrap()
                    .captures(&m_url)
                    .unwrap()
                    .get(1)
                    .unwrap()
                    .as_str();
                let log_no = Regex::new(&format!("\"url_{}_([0-9]{{12}})\"", blog_id))
                    .unwrap()
                    .captures(&body)
                    .unwrap()
                    .get(1)
                    .unwrap()
                    .as_str();

                Blog::Naver(format!("http://m.blog.naver.com/{}/{}", blog_id, log_no))
            } else {
                Blog::Naver(m_url)
            }
        } else {
            Blog::Unknown(String::from(url))
        }
    }

    fn get_subs(&self) -> Vec<String> {
        match self {
            Blog::Naver(url) => {
                let body = reqwest::get(url).unwrap().text().unwrap();
                let document = Html::parse_document(&body);
                let selector = Selector::parse(r#"a[href*="smi"]"#).unwrap();

                for e in document.select(&selector) {
                    let txt: Vec<&str> = e.text().collect();
                    println!("{:?}", txt);
                    println!("{}", e.value().attr("href").unwrap());
                }

                println!(
                    "{:?}",
                    document
                        .select(&selector)
                        .map(|e| e.value().attr("href").unwrap())
                        .collect::<Vec<&str>>()
                );
                vec![]
            }
            Blog::Unknown(url) => {
                let body = reqwest::get(url).unwrap().text().unwrap();
                let document = Html::parse_document(&body);
                let selector = Selector::parse(r#"a[href*="zip"]"#).unwrap();
                for e in document.select(&selector) {
                    let txt: Vec<&str> = e.text().collect();
                    println!("{:?}", txt);
                    println!("{}", e.value().attr("href").unwrap());
                }

                println!(
                    "{:?}",
                    document
                        .select(&selector)
                        .map(|e| e.value().attr("href").unwrap())
                        .collect::<Vec<&str>>()
                );
                vec![]
            }
        }
    }
}

fn main() {
    let settings = Settings::parse_yml("settings.yml");
    settings.source[0].get_links("dumbbell");
    // let urls = [
    //     "http://blog.noitamina.moe/221391147667",
    //     "http://blog.naver.com/cobb333/221391135993",
    //     "http://blog.naver.com/PostList.nhn?blogId=harne_&categoryNo=260&from=postList",
    //     "http://melody88.tistory.com/631",
    //     "https://mihorima.blogspot.com/2018/11/05_4.html",
    //     "http://godsungin.tistory.com/200",
    //     "https://blog.naver.com/qtr01122/221391146050",
    // ];

    // for url in urls.iter() {
    //     let blog = Blog::from_url(url);
    //     println!("{:?}", blog.get_subs());
    // }
}
