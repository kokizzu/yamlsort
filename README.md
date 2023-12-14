
# yamlsort

a simple tool to sort yaml

similar to [yamlflatten](//github.com/kokizzu/yamlflatten) this tool make it easy to diff 2 different yaml by sorting it (but not flatten it)

```
# example usage
./yamlsort input.yaml > output.yaml
```

for example:
```
# input.yaml
b:
 y: "abc"
 x:
  - 2
  - 1
 z: foo
c: bar
a: baz
```

became:
```
a: baz
b:
  x:
  - 2
  - 1
  "y": abc
  z: foo
c: bar
```

