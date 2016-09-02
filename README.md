# Thrift-demo

Thrift-demo is written so other engineers could understand how to use Thrift, and how to make Thrift
RPC calls. It supports PHP and Go, with Scala to be added soon.
By default, all binary protocol is used together with HTTP transport. You can easily change
protocol - just use different protocol class. Remember you have to change both server and client
code. Changing transport is slightly more difficult and it isn't covered by this demo.

You need to start only one server and then use any client to make Thrift RPC calls. All calls to be
made from root directory.

# make
Thrift generated files are already included in this repository (under gen-go and gen-php). If you
want to rebuild, you can run `make` (the output should be the same if Thrift version is 0.9.3).

# Servers

## PHP
To start a PHP server, run:

    php -S localhost:8080 php/server.php

## Go
To start Go server, symlink this repo to `$GOHOME/src/`. Then, run

    go run go/server.go

## Scala
TODO

# Clients
When running client, make sure server is up and running. Otherwise, you will get Thrift exception
connection refused.

## PHP
To run a PHP client, run following:

    php php/client.php

## Go
To run a Go client script, symlink this repo to $GOHOME/src (if already not done), and execute:

    go run go/client.go

## Scala
TODO
