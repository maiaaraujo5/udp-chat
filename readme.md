# *UDP CHAT*
[![Tests](https://github.com/maiaaraujo5/udp-chat/actions/workflows/test.yaml/badge.svg)](https://github.com/maiaaraujo5/udp-chat/actions/workflows/test.yaml)
[![Lint](https://github.com/maiaaraujo5/udp-chat/actions/workflows/lint.yaml/badge.svg)](https://github.com/maiaaraujo5/udp-chat/actions/workflows/lint.yaml)

# *Dependencies*
* Golang >= 1.16
* Redis

# *Architecture*
![img.png](img.png)

# *Main Rules*
* The client can only delete a message sent by him
* When one message is deleted it disappears for all client and is deleted in the repository. 
* The length of the history is configurable, now is set to **20**, but you can change in the **default.yaml**
* The user is defined by the client ip and port connected in server
* When all clients disconnect from the chat the DB is flushed

# *Technical Decisions*
* Uses the `container/list` to bring performance to the operations like delete a message from history in server
* Redis is not a strong dependency, if something bad happens to redis the chat would continue working.
* The datatype store in redis is a List
* The value store in redis is composed by {id}-{user_id}-{message}. Eg: 123-127.0.0.1:5315-Hello!

# *How to Run The Application*
 ## Server:
  To run the server we have some ways:

  **Note:** All the commands describes below looks for configuration in **default.yaml** and **development.yaml** if you want to alter some configs alter in this yamls or alter the command to search the config in another yaml

 ### make run-server-with-redis:
 Use this command if you don't have one instance of redis. This command will run one docker-compose to up one instance of redis and run the application

 ### make run-server:
 Use this command if you have one instance of redis. To use your redis instance change the configuration in **config/server/development.yaml**

 ### make docker-run-server-with-redis:
 Use this command if you want to run the server with docker, and you don't have one instance of redis. This command will run one docker-compose to up one instance of redis, build the server application, create and run the docker image of server

 ### make docker-run-server:
 Use this command if you have one instance of redis and want to run the server with docker. This command will create and run one docker image of server. To use your redis instance don't forget to change the config in **config/server/development.yaml**

 **Note:** Alter the redis configuration before run this command.

## Client:
To run the client use this command **make run-client**

# *How to use the client*
After the client is running, he will accept the following commands:

* **/msg**
Use this command to send new messages for the server. Eg:
  ``/msg Hello``
  
* **/del** Use this command to delete messages previously sent by you. This command receives the id of the message generate when the message was sent. Eg:
``/del 123``
 
* **/quit** Use this command to leave the room

# *Makefile commands*

## make unit-test-server:
* This command will run the unit tests of server

## make unit-test-client:
* This command will run the unit tests of client

## make end-to-end-tests:
* This command will run the end-to-end tests

## make lint:
* This command will run the linters configured. You can see the config in **./build/golangci-lint/config.yml**

## make docker-compose-run-redis:
* This command will run one instance of redis in your machine using docker-compose.

## make build-server:
* This command will create a server executable extension in **./dist/server**

## make build-client:
* This command will create a client executable extension in **./dist/client**

## make run-server-with-redis:
* This command will run one instance of redis in your machine using docker-compose and run the server application using go

## make run-server:
* This command will run the server application using go

## make run-client:
* This command will run the client application using go

## make docker-run-server:
* This command will create and run a docker image for the server application.

## make docker-run-server-with-redis:
* This command will create and run a docker image for the server application, and run one instance of redis in your machine using docker-compose


# *Libraries*
* [go-redis](https://github.com/go-redis/redis)
* [gostart](https://github.com/maiaaraujo5/gostart)
* [fx](https://github.com/uber-go/fx)
* [miniredis](https://github.com/alicebob/miniredis)
* [testify](https://github.com/stretchr/testify)
