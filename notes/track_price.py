import requests
import json
import re

from bs4 import BeautifulSoup

find_tags_dict = {
    '{"key":"vas-common-vm"}': "buyboxPrice",
    '{"key":"buyBackPageState"}': "productPrice",
}


def get_product_name(url: str):
    """
    Get Product Name from url given `urls.txt` file
    This re, replace solution is not optimal, works for now
    need to refactor later
    """
    product_name = re.match(r"https\:\/\/www\.amazon\.in\/[A-Za-z0-9-]*\/", url)
    product_name = product_name.group(0)
    product_name = product_name.replace("https://www.amazon.in/", "")
    product_name = product_name.replace("/", "")
    return product_name


def find_script_tags(key: str, broth):
    """
    return all <script> tags containing data-a-state of keys
    """
    return str(broth.find_all("script", {"data-a-state": key})[0].contents[0])


def get_price_from_tags(soup):
    """
    check find_tags_dict and return price value
    """
    for key, price_tag_value in find_tags_dict.items():
        script_tags = find_script_tags(key, soup)
        if script_tags:
            output = json.loads(script_tags)
            if price_tag_value in output:
                return output[price_tag_value]


def get_url_price(url):
    headers = {
        "authority": "www.amazon.in",
        "pragma": "no-cache",
        "cache-control": "no-cache",
        "dnt": "1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
        "sec-fetch-site": "none",
        "sec-fetch-mode": "navigate",
        "sec-fetch-dest": "document",
        "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
    }
    # Download the page using requests
    # print("Downloading {}".format(url))
    r = requests.get(url, headers=headers)
    # Simple check to check if page was blocked (Usually 503)
    if r.status_code > 500:
        if "To discuss automated access to Amazon data please contact" in r.text:
            print("Page %s was blocked by Amazon. Please try using better proxies\n" % url)
        else:
            print(
                "Page %s must have been blocked by Amazon as the status code was %d"
                % (url, r.status_code)
            )
        return None
    # Pass the HTML of the page and create
    soup = BeautifulSoup(r.text, "html.parser")
    return get_price_from_tags(soup)


if __name__ == "__main__":
    with open("urls.txt") as u:
        for url in u:
            print("Product: {}, Price: {}".format(get_product_name(url), get_url_price(url)))
