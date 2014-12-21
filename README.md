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

Install:

``` bash
go get github.com/fabioberger/dockerize-tutorial
```

Start Server:

``` bash
go run main.go
```

### Install and run in a docker container:

Setup a Postgres Docker Container as described in the tutorial. Then run these two commands:

docker pull fabioberger/dockerize-tutorial

docker run -d -p 8080:4000 --name tutapp --link dbtest:postgres fabioberger/dockerize-tutorial