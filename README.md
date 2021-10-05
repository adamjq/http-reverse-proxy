# http-reverse-proxy

This repo contains a simple [reverse-proxy server](https://en.wikipedia.org/wiki/Reverse_proxy) written in Golang.

## Dependencies:

- Golang
- Docker
- [wire](https://github.com/google/wire)

## Development

Generate mocks and dependencies to inject with [wire](https://github.com/google/wire):
```bash
make generate
```

Run unit tests:

```bash
make test
```

## Usage:

```bash
docker-compose up
```

Run [the requests](./requests.http) and inspect logs to see the service working.

## References

- [What is a reverse proxy server?](https://www.nginx.com/resources/glossary/reverse-proxy-server)
- [Dependency Injection in Go](https://blog.drewolson.org/go-dependency-injection-with-wire)
