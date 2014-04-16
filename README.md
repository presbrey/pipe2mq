# pipe2mq

## Install

Use the `go get` command eg.

    go get github.com/presbrey/pipe2mq

## Usage

`pipe2mq` accepts command-line arguments:
~~~
  -backlog=8192: incoming channel capacity
  -backoff=1s: pause between errors
  -escapeBody=false: request body will be Go-escaped
  -exchange="test": AMQP exchange name
  -key="test": AMQP routing key
  -tag="pipe2mq": AMQP consumer tag
  -uri="amqp://localhost:5672/": AMQP URI
~~~

## License

[MIT](http://joe.mit-license.org/)
