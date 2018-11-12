#!/usr/bin/env python
import yaml
import feedparser
import requests
import os
from lxml import etree
import re

with open('settings.yml') as f:
    settings = yaml.load(f)

def get_path(path):
    return os.path.realpath(os.path.expanduser(path))

def get_torrent_files():
    rss_cache = {}
    targets = []
    for target in settings['target']:
        if target['source'] not in rss_cache:
            for source in settings['source']:
                if source['name'] == target['source']:
                    rss_cache[target['source']] = feedparser.parse(source['rss'])
                    break
            else:
                print('No rss information: {}'.format(target['rss']))
        rss = rss_cache[target['source']]
        for entry in rss['entries']:
            if target['title'].lower() in entry['title'].lower():
                targets.append((entry['title'], entry['link']))

    for name, url in targets:
        with open(os.path.join(get_path(settings['path']['torrent']), name), 'wb') as f:
            f.write(requests.get(url).content)

def get_sub_naver(url):
    id_reg = '[a-z0-9_-]{5,20}'
    log_reg = '[0-9]{12}'
    if 'blog.naver.com' in url:
        matches = re.match('https?://blog\.naver\.com/({})/({})'.format(id_reg, log_reg), url)
    elif 'blog.me' in url:
        matches = re.match('https?://({})\.blog\.me/({})'.format(id_reg, log_reg), url)
    if not matches:
        raise
    blogId = matches.group(1)
    logNo = matches.group(2)
    url = 'https://m.blog.naver.com/{}/{}'.format(blogId, logNo)
    res = requests.get(url).text
    html = etree.HTML(res)
    urls = set(html.xpath('//a[contains(@href,"blogattach.naver.net")]/@href'))
    return list(urls)

def get_sub(url):
    res = requests.get(url, headers={'User-Agent': 'Mozilla/5.0'}).text
    html = etree.HTML(res)

    id_reg = '[a-z0-9_-]{5,20}'
    log_reg = '[0-9]{12}'
    info = None
    urls = html.xpath('//frame[contains(@src,"blog.naver.com")]/@src')
    if urls:
        info = re.search('blog\.naver\.com/({})/({})'.format(id_reg, log_reg), urls[0])
    urls = html.xpath('//frame[contains(@src,"PostView.nhn")]/@src')
    if urls:
        info = re.search('blogId=({})&logNo=({})'.format(id_reg, log_reg), urls[0])
    if info:
        blogId = info.group(1)
        logNo = info.group(2)
        url = 'https://m.blog.naver.com/{}/{}'.format(blogId, logNo)
        res = requests.get(url).text
        html = etree.HTML(res)
        urls = set(html.xpath('//a[contains(@href,"blogattach.naver.net")]'))
        return list(urls)

    links = ['tistory.com/attachment','egloos.com/pds', 'drive.google.com/uc', 'blogattach.naver.net']
    xpath = '//a[contains(@href,"{}")]'
    for link in links:
        urls = set(html.xpath(xpath.format(link)))
        if urls:
            break
    return list(urls)

def fetch_sub():
    anissia_api = 'https://www.anissia.net/anitime/cap?i={}'
    for target in settings['target']:
        api = anissia_api.format(target['subtitle']['anissia'])
        for sub in requests.get(api).json():
            if target['subtitle']['author'] == sub['n']:
                print(sub['a'])

def download(url):
    pass

test = ['https://blog.naver.com/ych622/221223304033',
        # 'https://cndska15.blog.me/221223202798',
        # 'https://gkgk6265.blog.me/221222515618',
        # 'https://blog.naver.com/neverduck1/221217466964',
        # 'https://blog.naver.com/qtr01122/221221478283',
        # 'https://blog.naver.com/ych622/221221662230',
        # 'https://blog.naver.com/bluewater91/221221171621',
        # 'https://blog.naver.com/qtr01122/221221223891',
        # 'https://blog.naver.com/qtr01122/221222045958',
        # 'https://blog.naver.com/bofgirl/221195230686',
        # 'https://mihorima.blogspot.kr/2018/03/09_52.html',
        # 'http://kwangwaul.egloos.com/6305572',
        # 'http://prisis.tistory.com/1278',
        'http://ogura-yui.moe/221221377432',
        'http://bullyangblog.tistory.com/1324',
        'https://mihorima.blogspot.kr/2018/03/09_5.html',
        'http://lalin.tistory.com/90',
        'http://lalin.tistory.com/106',
        'https://mihorima.blogspot.kr/2018/03/06.html',
        'http://blog.naver.com/cobb333/221220597539',
        'https://blog.naver.com/qtr01122/221220598858',
        'https://mihorima.blogspot.kr/2018/02/beatless-06.html',
        'http://prisis.tistory.com/1276',
        'https://blog.naver.com/bluewater91/221220663560',
        'https://blog.naver.com/sungwook0208/221222309060',
        'http://kannasub.tistory.com/entry/%EC%82%B0%EB%A6%AC%EC%98%A4-%EB%82%A8%EC%9E%90-9%ED%99%94-%EC%9E%90%EB%A7%89',
        'http://fuko.tistory.com/3030',
        'https://blog.naver.com/baby1255/221122421926',
        'https://mihorima.blogspot.kr/2018/03/citrus-09.html',
        'https://blog.naver.com/qtr01122/221179478524',
        'http://prisis.tistory.com/1277',
        'https://blog.naver.com/qtr01122/221221162614']

# for a in test:
#     subs = get_sub(a)
#     for sub in subs:
#         print('----------------------')
#         print(sub)
#         print(sub.get('href'))
#         print(sub.xpath('..//*[contains(text(), "smi")]/text()'))

# fetch_sub()
get_torrent_files()
