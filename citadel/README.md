## Demo Code for self learning

This repo contains code of a demo website API where you can store your book details and fetch it.

This webapp API has following things written into it: -

* MySql integration and CRUD operation on MySql.
* Exposing API metrics using Prometheus on `/metrics` endpoint.
* A health check endpoint that display the app status initializing / ready /database connection status on `/healthz` endpoint.

To run Mysql use the following command: -

`docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:latest`