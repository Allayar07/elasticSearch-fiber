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

# For see all data in elasticsearch:
```
http://localhost:9200/books/_search?size=100&q=*:*
```

