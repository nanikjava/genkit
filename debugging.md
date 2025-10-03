smart search enginer finding info about programming language...send brief response


# basic-gemini example

To run so that it reload

```
find /workspaces/genkit/go -name '*.go' | entr -r air
```

make sure `entr` and `air` is installed 

`go.mod` must be


```
module basic_gemini

go 1.24.5

require (
	github.com/firebase/genkit/go v1.0.5
)

require (
...
)

replace github.com/firebase/genkit/go => /workspaces/genkit/go
```

The file `.air.toml` looks like the following

```
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
cmd = "go build -o basic_gemini"
bin = ""
full_bin = ""
include_dir = ["/workspaces/genkit/go/genkit"]
include_ext = ["go", "tpl", "tmpl", "html"]
include_file = ["go.mod", "go.sum"]
....
```

# Typescript

Make sure `pnpm` is installed

From `/workspaces/genkit/genkit-tools/cli` directory run this.

```
nodemon --watch src/bin/* --exec "GEMINI_API_KEY=<key> npx ts-node src/bin/genkit.ts start -- /workspaces/genkit/go/samples/basic-gemini/basic_gemini"
```

When changes are made into file inside `/src/bin/*` it will reload.

This also is useful if the Go example code (or libraries) are changes to reload/