# URLShortner
This service shortens urls via http post\get and grpc with post and get rpc's.
This service works with content-type "text/plain; charset=utf-8".

# Prerequestities
Project is built with the option of usage of Postgres database.

Database Login credentials should be set in ./app/app.go with constants. 

By default service runs with in-memory storage, to run with postgre in should be assigned 1 run argument "postgres".

Dockerfile is set for default in-memoru use, for postgres it should be changed.

Default address is set in the app package.
