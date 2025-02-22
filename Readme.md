## LSD

lsd is a command-line utility designed for developers and system administrators to visualize directory structures in a tree-like format.

### Installation

```sh
go install github.com/sharan-cj/lsd@latest
```

### Usage

```sh
lsd path/to/dir -M --verbose
```

```cmd
Flags:
  -d, --depth         Depth of the tree. Default is 2.
  -h, --help          help for lsd
  -M, --max           Print the max depth of the tree. Default is false.
      --verbose       Verbose mode. Default is false.
  -v, --version       version for lsd
```
