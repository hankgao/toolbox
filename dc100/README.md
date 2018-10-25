# Overview

df100 distribute new coins to 100 addresses

## API used 

Create transaction

``` bash
URI: /api/v1/wallet/transaction
Method: POST
Content-Type: application/json
Args: JSON body, see examples
```

{
    "ignore_unconfirmed": false,
    "hours_selection": {
        "type": "manual"
    },
    "wallet": {
        "id": "foo.wlt"
    },
    "change_address": "nu7eSpT6hr5P21uzw7bnbxm83B6ywSjHdq",
    "to": [{
        "address": "fznGedkc87a8SsW94dBowEv6J7zLGAjT17",
        "coins": "1.032",
        "hours": 7
    }, {
        "address": "7cpQ7t3PZZXvjTst8G7Uvs7XH4LeM8fBPD",
        "coins": "99.2",
        "hours": 0
    }]
}