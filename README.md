# go-raft Leader Election
Implementation of leader election described in the Raft consensus algorithm

## Design
The design of this is built around the state pattern where the node state
acts as an FSM with incoming requests and timeouts acting as inputs. The
`pkg/raft` contains the main thread which spawns goroutines for the server
and client which pass the requests/responses back to the main goroutine
via channels to be handled by the FSM.

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

## Future
* Extend to covery log replication and see full dissertation
https://github.com/ongardie/dissertation/blob/master/stanford.pdf
