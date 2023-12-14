
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
 y:
 x:
  - 2
  - 1
 z:
c:
a:
```

became:
```
a: null
b:
  x:
  - 2
  - 1
  "y": null
  z: null
c: null
```

