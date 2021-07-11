import requests
from selectorlib import Extractor, Formatter
from pprint import pprint
import re

# Define a formatter for Price
class Price(Formatter):
    def format(self, text):
        price = re.findall(r'\d+\.\d+',text)
        if price:
            return price[0]
        return None
formatters = Formatter.get_all()
extractor = Extractor.from_yaml_file('./selectors.yml',formatters=formatters)

#Download the HTML and use Extractor
r = requests.get('https://www.amazon.in/Samsung-inches-Crystal-Ultra-UA50AUE70AKLXL/dp/B091GY19M3/')
data = extractor.extract(r.text)
pprint(data)