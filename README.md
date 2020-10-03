## Purpose
Refresh my React knowledge, learn Docker and Golang on the fly while I'm at it.

## Acceptance Criteria
Have a list of pictures where you are able to tag the images and preserve the tagged state after refresh

## Requirements
- docker (I used version 19.03.12)
- docker-compose (I used version 1.27.2)

Free Ports:
- 8080
- 5432
- 3000

## Instructions
```
docker-compose build
docker-compose up # CMD + C
docker-compose up # Need to run twice (DB vs. Web service race condition - out of scope!)
```
Go to `http://localhost:3000`
