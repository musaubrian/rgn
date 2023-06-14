# RepoGen(rgn)

RepoGen is a lightweight command-line interface (CLI) tool written in Go,
designed to simplify the process of creating and initializing repositories on GitHub.

It offers a **simpler* approach to repository creation by providing options to generate repositories:
 - Empty repostitores(**no README or .gitignore**)
 - repositories with only a .gitignore file,
 - repositories with only a README file,
 - repositories with both a README file and a .gitignore file.

## Installation

To use RepoGen, you'll need to have Go installed on your system.
Follow these steps to install and set up RepoGen:

### From Source
1. Clone the RepoGen repository from GitHub:
```bash
git clone https://github.com/musaubrian/rgn.git

#OR

git clone git@github.com:musaubrian/rgn.git
```
3. Build the RepoGen binary:

```bash
cd rgn

go build
```
4.
   a) Move the binary to a directory in your system's PATH (e.g., /usr/local/bin)
or add the current directory to your PATH environment variable.
 OR just run the generated executable
 ```bash
 ./rgn

 # Windows

 ./rgn.exe
 ```
   b) Install it system wide using `make`
 ```bash
 make install

 rgn
 ```

6. Once you've completed the installation, you can start using RepoGen right away.

### Using `go install`
1. Install the binary
```bash
go install github.com/musaubrian/rgn@latest
```
2. That's it, just run the binary
```bash
rgn
```
