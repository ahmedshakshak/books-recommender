# books-recommender

## How to run tests:

1. run the following command `docker-compose up -d test`
2. then this command to setup the env `./setup_local_env.sh`

now the app is read for local development, you can do the following to test it
1. run `docker-compose exec test sh`
2. then `go test ./...` that should run all the tests



## How to run server:

1. run the following command `docker-compose up -d database`
2. then this command to setup the env `./setup_local_env.sh` it might take time till MySQL is ready for serving requests
3. then run the following command `docker-compose up -d web` to start the server on the localhost (you can control server port number through APP_PORT env var in .env file, and even all env params like MySQl and the notifier)

now the app is ready to serve network network requests


## File Hierarchy
entities/
It’s responsible for creating application entities (shared classes that all models can deal with) and shared functions.

usecases/
It’s responsible for handling application usecases and acceptance criteria without paying attention to any technical flow, and it operates against interfaces that handle all technical flow

adapters/
It’s responsible for handling technical flow for ‘usecases/’ and delegates the technical details to other classes

models/
It’s responsible for dealing with MySQL

controllers/
It’s responsible for handling user inputs (API parameters) to fulfill application usecases and decorating usecases reply

register/
It’s responsible for solving all classes’ dependencies and providing a controller ready to handle the traffic

xxxtest/
It’s responsible for handling all service tests
