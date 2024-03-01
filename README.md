// README.md
### Introduction
Story: There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

Goal: To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

### Required JSON structure
* [["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
* [["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
* [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

### Specifications
* Your miscroservice must listen on port 8080 and expose the flight path tracker under the /calculate endpoint.

### How to run and build
* go run . (go run main.go)
* go build .

### Usage
* Connect to the API using Postman on port 8080.
* http://localhost:8080 

### API Endpoints
| HTTP Verbs | Endpoints | Action | 
| --- | --- | --- |
| GET | /calculate | Get the final result of the flights  |

| Query | Required | Type | Example | Result |
| --- | --- | --- | --- | --- |
| path | required | string | /calculate?path=[["ATL", "EWR"],["SFO", "ATL"]] | ["SFO", "EWR"] |