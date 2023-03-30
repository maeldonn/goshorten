# Goshorten

_(An URL shortener written in Go)_

## Running the server

To run the server you will need GNU make and Go SDK installed.

```sh
$ make run
```

## Interact with the server

To create shortened URLs, you will need an HTTP client. I recommand using curl.

```sh
$ curl localhost:8080/slug
```

## Useful commands

Create a new slug
```
$ curl http://localhost:8080 -d '{"slug": "goshorten", "url": "http://github.com/maeldonn/goshorten"}'
```

Use shortened URL
```
$ curl http://localhost:8080/goshorten
```

Delete a shortened URL
```
$ curl -X DELETE http://localhost:8080/goshorten
```
