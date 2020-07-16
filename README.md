# wine-dal
Wine Data Access Layer

## Start the API
Run the following command in build folder:
`docker-compose up -d`

This brings up a three node Elastic cluster and a Kibana instance.

To see if all the nodes are up and running:
`curl -X GET "localhost:9200/_cat/nodes?v&pretty"`
Once they are up, Kibana can be accessed in http://localhost:5601/

## Setting up index
New index created in Kibana PUT /bottles with mapping settings in file:

data/elastic-index-mapping.json (in management > dev tools)

## Endpoints of API
API can be started at the moment with the following command under the ./src/ folder:
`go run *.go`

- GET : /wine/{wineID}

Returns info about the wine in JSON format

- POST : /wine/
  
Example body request call:

```
{
    "country": "Spain",
    "region": "Rioja",
    "lage": "Morstein",
    "sweetness": "lieblich",
    "sugarLevel": "18%",
    "wineType": "Stillwein",
    "wineColor": "white",
    "title": "Doppio passo",
    "description": "This wine is...",
    "alcoholLevel": "12%",
    "vintage": "2018",
    "validEAN": true,
    "acidity": "5%",
    "winery": "Ramon Bilbao",
    "grape": "Tempranillo",
    "appellation": "Chianti"
}
```
- POST : /wine/bulk

Example body request call:

```
{
"wines" : [
{"country":"Spain","region":"Rioja","lage":"Morstein","sweetness":"lieblich","sugarLevel":"18%","wineType":"Stillwein","wineColor":"white","title":"Doppio passo","description":"This wine is...","alcoholLevel":"12%","vintage":"2018","validEAN":true,"acidity":"5%","winery":"Ramon Bilbao","grape":"Tempranillo","appellation":"Chianti"},
{"country":"Spain","region":"Rioja","lage":"Morstein","sweetness":"lieblich","sugarLevel":"18%","wineType":"Stillwein","wineColor":"white","title":"Doppio passo","description":"This wine is...","alcoholLevel":"12%","vintage":"2025","validEAN":true,"acidity":"5%","winery":"Ramon Bilbao","grape":"Tempranillo","appellation":"Chianti"}
]
}
```

Licensed under the [MIT License](LICENSE)
