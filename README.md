# go-raft
Implementation of Raft consensus algorithm in Go

## Building
Build using [Bazel](https://bazel.build/). This builds all source files,
mocks, tests and generates Protocol Buffers/gRPC files:
```
  $ bazelisk build ...
  (or bazelisk build //pkg/raft:go_default_library)
```

Note bazelisk is used to ensure the correct version of Bazel is always used,
installed from [Bazelisk](https://github.com/bazelbuild/bazelisk).

Use [Gazelle](https://github.com/bazelbuild/bazel-gazelle) to generate
all Bazel build files:
Generate `BUILD.bazel` files when a dependency is added or file layout
modified using [Gazelle](https://github.com/bazelbuild/bazel-gazelle):
```
  $ bazelisk run //:gazelle
```

Run unit tests:
```
  $ bazelisk test ...
```
