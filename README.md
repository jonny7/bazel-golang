# Overview

Bazel with Go for Monorepos.

# Instructions
1. Install Bazel
```shell
brew install bazel
```
2. Create a WORKSPACE file in the route of the project
3. Create a BUILD file in the route
4. Run Gazelle
```shell
bazel run //:gazelle  
```
5. Update project dependencies (if needed)
```shell
bazel run //:gazelle --update-repos -from_file=go.mod
```
6. Might need to `chmod` the shell file
```shell
sudo chmod 755 status.sh
```
7. Run with stamp
```shell
bazel build --stamp --workspace_status_command=$(pwd)/status.sh //... 
```
8. Run the binaries
```shell
bazel-bin/cmd/a/a_/a 
bazel-bin/cmd/b/b_/b
bazel-bin/cmd/c/c_/c
```
9. Update a single binary and recheck the stamps
