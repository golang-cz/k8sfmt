- [ ] Move & export `formatYAML()` function in top-level pkg, once we finalize the func signatures
- [ ] Support stdin/stdout if no file is given
- [ ] Format files concurrently
- [ ] Support folders?
- [ ] Support recursive folder via ./...?
- [ ] Key sorting: Perhaps we could prioritize `apiVersion`, `kind` and `metadata` fields over `spec` and `data`?
- [ ] Support [YAML streams](https://yaml.org/spec/1.2.2/#yaml-stream) delimited by `---`?
- [ ] Support quoted scalars per spec https://yaml.org/spec/1.2.2/#yaml-stream:
    ```
    unicode: "Sosa did fine.\u263A"
    control: "\b1998\t1999\t2000\n"
    hex esc: "\x0d\x0a is \r\n"

    single: '"Howdy!" he cried.'
    quoted: ' # Not a ''comment''.'
    tie-fighter: '|\-*-/|'
    ```