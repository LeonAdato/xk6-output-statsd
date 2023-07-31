# xk6-output-template
Is template for k6 output [extensions](https://k6.io/docs/extensions/guides/what-are-k6-extensions/)

You should make a repo from this template and go through the code and replace everywhere where it says `template` in order to use it.
There are more instructions and comments inline.

> :warning: the API of k6 outputs [will likely change in the future](https://github.com/grafana/k6/issues/2430), so repos using it (like this repo) are not guaranteed to be working with any future version of k6.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git
- [xk6](https://github.com/grafana/xk6)

1. Build with `xk6`:

```bash
xk6 build --with github.com/grafana/xk6-output-template
```

This will result in a `k6` binary in the current directory.

2. Run with the just build `k6:

```bash
./k6 run -o xk6-template <script.js>
```
