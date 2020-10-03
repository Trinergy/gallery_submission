## Purpose
Refresh my React knowledge, learn Docker and Golang on the fly while I'm at it.

## Acceptance Criteria
- Gallery/List view of images
- Images have tag/flag button
- Tagged images are still rendered as tagged

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
docker-compose up # Need to run twice (Web service fails on first DB init - known docker problem, will not address, out of scope.)
```
Go to `http://localhost:3000`
