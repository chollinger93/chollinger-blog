import requests
import random

def req(symbol):
    response = requests.get('http://localhost:3000/stocks?symbol={}'.format(symbol))
    print(response.json())

if __name__ == '__main__':
    symbols = ["PEAR", "SQRL", "NULL", "ZVZZT"]
    for i in range(1000):
        symbol = random.choice(symbols) 
        print('Testing API with {}'.format(symbol))
        req(symbol)