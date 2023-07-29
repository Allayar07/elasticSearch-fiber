# Clone the repository
* Run this command if you have ssh key:
```
git clone git@github.com:Allayar07/elasticSearch-fiber.git
```
* For http:
```
git clone https://github.com/Allayar07/elasticSearch-fiber.git
```
# Run  Project
* Use this command:
```
docker compose up
```
# DB Migrations
* Command:
```
migrate -path ./schema -database 'postgres://postgresql:password0701@localhost:5475/practice?sslmode=disable' up
```
# Create elastic analyzer which latin to russian:
## URL params:
* URL ```http://localhost:9200/books```
* Method ```PUT```
* Response:
```
{
  "settings": {
    "analysis": {
      "filter": {
        "transliterator": {
          "type": "icu_transform",
          "id": "Any-Latin; NFD; [:Nonspacing Mark:] Remove; NFC"
        }
      },
      "analyzer": {
        "latin_to_russian": {
          "type": "custom",
          "tokenizer": "standard",
          "filter": ["lowercase", "transliterator"]
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "description": {
        "type": "text",
        "analyzer": "latin_to_russian"
      },
			"name": {
        "type": "text",
        "analyzer": "latin_to_russian"
      }
    }
  }
}
```
## Success response:
* Status: ```200```
* Body:
```
{
	"acknowledged": true,
	"shards_acknowledged": true,
	"index": "books"
}
```
## Ror request:
## Endpoint for create
* Method ```POST```

* URL :``` http://localhost:7575/create```

* Body for create book:
```
{
	"name": "string",
	"pageCount": int,
	"author": "unkcnown",
	"description": ["string"],
	"authorEmail": "string"
}
```
* Success Response:

```
 {
	"id": 0,
}
```

## Endpoint for search:
* Method ```GET```
* Params ```query param```

* URL : ```http://localhost:7575/search?find=```

* Success Response:

```
    [
        {
            "id": 0,
            "name": "string",
            "pageCount": int,
            "author": "unkcnown",
            "description": ["string"],
            "authorEmail": "string"
        },
    ]

```

## Endpoint for update:
* Method ```PUT```

* URL : ```http://localhost:7575/update```
* Body for update book:
```
    {   
        "id": 0,
        "name": "string",
        "pageCount": int,
        "author": "unkcnown",
        "description": ["string"],
        "authorEmail": "string"
    }
```

* Success Response:

```
   "message":"OK"

```

## Endpoint for deleting book

* Method ```DELETE```

* URL : ```http://localhost:7575/delete```
* Body for update book:
```
    {   
        "ids": [0],
        
    }
```

* Success Response:

```
   "message":"OK"

```



