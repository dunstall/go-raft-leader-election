# go-raft
Implementation of Raft consensus algorithm in Go

## Building
This library is build using [Bazel](https://bazel.build/). The main benifit
of this is generating all Go files from Protocol Buffers and gRPC, and
producing the library output in a single command:
```
  $ bazelisk build ...
  (or bazelisk build //pkg/raft:go_default_library)
```

Note bazelisk is used to ensure the correct version of Bazel is always used,
installed from [Bazelisk](https://github.com/bazelbuild/bazelisk).

We also [Gazelle](https://github.com/bazelbuild/bazel-gazelle) to generate
all Bazel build files:
```
  $ bazelisk run //:gazelle
```
