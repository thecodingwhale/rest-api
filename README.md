# REST API
a simple rest api created with go with simple test.

to run quickly the app
```
go build
./rest-api
```

to run test
```
go test
```

# Setup
create `.env` file will the config
```
APP_DB_USERNAME=root
APP_DB_PASSWORD=
APP_DB_NAME=rest_api

TEST_DB_USERNAME=root
TEST_DB_PASSWORD=
TEST_DB_NAME=rest_api_test
```

# Database Migration:
Currently the app uses goose for handling database migrations:

on first run:
```
goose mysql "root:password@/dbName" up
```

# TODO
  - add more test
