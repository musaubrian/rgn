# RepoGen(rgn)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

RepoGen is a lightweight command-line interface (CLI) tool written in Go,
designed to simplify the process of creating and initializing repositories on GitHub.

You can also see what issues are assigned to you.

<details>
<summary>Repository Generation Options</summary>

- [x] Empty repositories (no README or .gitignore)
- [x] Repositories with only a .gitignore file
- [x] Repositories with only a README file
- [x] Repositories with both a README file and a .gitignore file
- [x] LICENSE generation?

</details>



## Installation

> **Note**
>
> You will need to generate new personal token with repo creation permissions

To use RepoGen, you'll need to have Go installed on your system.
Follow these steps to install and set up RepoGen:


### a. From Source

1. clone the repository
```bash
git clone https://github.com/musaubrian/rgn

cd rgn
```
2. Build it.
```bash
go build .
# or
make

# MAKE IT GLOBALLY ACCESSIBLE
# Manually move the binary(rgn) to the GOPATH usually HOME_DIR/go/bin
# or 

make install
```
3. Run it
```bash
# If you did not install it globally
./bin/rgn

# If you did install it globally
rgn
```

### b. Using `go install`
```bash
go install github.com/musaubrian/rgn@latest

# If the version installed doesn't match the current release version

go install github.com/musaubrian/rgn@current_version
```

## Uninstalling
```bash
# If you made it globally vailable
make uninstall

# If not
make clean
```
