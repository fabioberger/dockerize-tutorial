Example Golang + Postgres App For Dockerize Tutorial
-----------------------------------------------

This repo contains the accompanying demo app for the [How To Dockerize Golang Webapp With Postgres DB](http://www.fabioberger.com/) tutorial at [fabioberger.com](fabioberger.com).

DB instructions:

```
CREATE USER app;
CREATE DATABASE testapp;
GRANT ALL PRIVILEGES ON DATABASE testapp TO app;
```