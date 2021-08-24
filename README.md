# Simple Price Tracker

## Track amazon prices for a particular product regularly
1. give product url in the urls.txt file
2. run `track_price.py` file

```bash
python3 track_price.py
Product: OnePlus-125-7-inches-Android-50U1S, Price: 46999
Product: Samsung-inches-Crystal-Ultra-UA50AUE70AKLXL, Price: 52999.0
```

## Get Game CD Prices from gamenation.in
1. run `go run src/main.go` to start the server
2. make a GET api call '`localhost:8080/gameprice/<keyword>`'
you'll get a list of prices for given keyword

for example `localhost:8080/gameprice/cyber`
```json
[
    {
        "name": "Wolfenstein Cyberpilot VR - PS4 (Pre-owned)",
        "price": "₹599"
    },
    {
        "name": "Microsoft Xbox One X 1TB Cyberpunk 2077 Limited Edition - Xbox One (Pre-owned)",
        "price": "₹27999"
    },
    {
        "name": "Xbox One Controller (3rd Gen) Cyberpunk 2077 - Xbox One (Pre-owned)",
        "price": "₹3799"
    },
    {
        "name": "Cyberpunk 2077 - PS4",
        "price": "₹2399"
    },
    {
        "name": "Cyberpunk 2077 - PS4 (Pre-owned)",
        "price": "₹1999"
    },
    {
        "name": "Cyberpunk 2077 - Xbox One",
        "price": "₹2799"
    },
    {
        "name": "Cyberpunk 2077 - Xbox One (Pre-owned)",
        "price": "₹2199"
    }
]
```