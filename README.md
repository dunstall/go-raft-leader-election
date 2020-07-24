# go-raft
Implementation of Raft consensus algorithm in Go

## Building
This library is build using [Bazel](https://bazel.build/) which builds
all source files, mocks, tests and Protocol Buffers/gRPC:
```
  $ bazelisk build ...
  (or bazelisk build //pkg/raft:go_default_library)
```

Note bazelisk is used to ensure the correct version of Bazel is always used,
installed from [Bazelisk](https://github.com/bazelbuild/bazelisk).

Use [Gazelle](https://github.com/bazelbuild/bazel-gazelle) to generate
all Bazel build files:
```
  $ bazelisk run //:gazelle
```

To run unittests:
```
  $ bazelisk test ...
```
