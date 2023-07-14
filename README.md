# Minimal Svelte + TS + Vite + GO

This `minimal` template should help get you started developing with `Svelte` and `TypeScript` in `Vite` & `GO` as the backing server.

## Development

The Go backend in this repository uses the [`embed`](https://pkg.go.dev/embed)
package to embed the Svelte app inside the Go binary. Running `go build` in the
root will capture whatever is present in the `ui/build` subdirectory.

To ensure you have an up to date copy of the web app in your binary, you should:

- `cd ui`
- `pnpm install`
- `pnpm run build`
- `cd ..`
- `make`


## ui.go
`ui/ui.go` grabs the output of `pnpm run build` in the `build` folder and creates an embedded file-system out of it.

## server.go
`server.go` uses the embedded file-system created by `ui/ui.go` and serves it with [gorilla toolkit's mux](https://gorilla.github.io/)
I'm a big fan of the gorilla toolkit and after being archived it was recently announced that new life will be blown into the project.

If you want use the standard http package to serve the app you can easily do that by removing

```go
    r := mux.NewRouter()
    r.PathPrefix("/").Handler(fs)

    // Remove this line from the return object
    Handler: r,
```
And replace with

```go
    http.Handle("/", fs)
```

Now if you respin your application it will host your svelte app without using any 3rd party libraries.
