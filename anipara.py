#!/usr/bin/env python
import requests
import re
from lxml import etree


def get_service(url):
    return 'naver' if 'naver' in url else None


def get_sub(url):
    service = get_service(url)

    if service == 'naver':
        # change to mobile url (works even domain != naver.com)
        res = requests.get(url, headers={'User-Agent': 'Android'}).text
        matches = re.search("top\.location\.replace\('(.*)'\);", res)
        m_url = matches.group(1).replace('\/', '/')

        # get the first post if a given url is a postlist
        if 'PostList.nhn' in m_url:
            res = requests.get(url).text

            matches = re.search('blogId=([a-z0-9_-]{5,20})&', m_url)
            blog_id = matches.group(1)
            matches = re.search('"url_%s_([0-9]{12})"' % blog_id, res)
            log_no = matches.group(1)

            m_url = 'http://m.blog.naver.com/%s/%s' % (blog_id, log_no)

        # get download urls from mobile site
        res = requests.get(m_url).text
        html = etree.HTML(res)
        return html.xpath('//a[contains(@href,"blogattach")]/@href')
    else:
        return []


if __name__ == '__main__':
    tests = [
        'http://blog.noitamina.moe/221391147667',
        'http://blog.naver.com/cobb333/221391135993',
        'http://blog.naver.com/PostList.nhn?blogId=harne_&categoryNo=260&from=postList',
        'http://melody88.tistory.com/631',
        'https://mihorima.blogspot.com/2018/11/05_4.html',
        'http://godsungin.tistory.com/200',
        'https://blog.naver.com/qtr01122/221391146050',
    ]

    for test in tests:
        print(test)
        print(get_sub(test))
        print()
