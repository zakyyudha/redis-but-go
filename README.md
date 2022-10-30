# Redis But Go
A very simple ~~redis~~ key value store server writen with go

## Running the server
```cmd
$ sh ./run.sh
```

## Connect to the server
you can connect to this server using `redis-cli` with 
following command
```cmd
$ redis-cli -p 6379
```

## Available command
For now (maybe forever) the command that can be use is 
`ping`, `set [key] [value] [ex|px] [interval]`, `get [key]`

## Example
```cmd
on server
➜ redis-but-go git:(master) ✗ sh ./run.sh
Server started on 0.0.0.0:6379

on client
➜  ~ redis-cli -p 6379

set key value without expiry
127.0.0.1:6379> set foo bar
OK
127.0.0.1:6379> get foo
"bar"

set key value with 5 sec expiry
127.0.0.1:6379> set foo bar ex 5
OK
after 5 second
127.0.0.1:6379> get foo
(nil)
```