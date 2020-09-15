# mrpoc

Steps to reproduce:

- `rice embed-go -v`
- `go build`
- run `./mrpoc` and it consumes memory until OOM or manually killed, the
  migration never finishes.
