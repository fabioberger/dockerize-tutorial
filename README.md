Example Negroni Golang + Postgres App For Dockerize Tutorial
-----------------------------------------------

This repo contains the accompanying demo app for the [How To Dockerize Golang Webapp With Postgres DB](http://fabioberger.com/blog/2014/12/19/how-to-dockerize-golang-webapp-with-postgres-db/) tutorial at [fabioberger.com](fabioberger.com).

### Set Up Postgres Database:

Run these commands from ```psql``` to set up the app's postgres user and db:

```
CREATE USER app;
CREATE DATABASE testapp;
GRANT ALL PRIVILEGES ON DATABASE testapp TO app;
```

### Install and run locally:

Install dependencies:

``` bash
go get github.com/codegangsta/negroni
go get github.com/coopernurse/gorp
go get github.com/lib/pq
```

Start Server:

``` bash
go run main.go
```
