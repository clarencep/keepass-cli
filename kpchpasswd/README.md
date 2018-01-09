Usage
=====

```
Usage of kpchpasswd:
  -input string
        specify the input .kdbx file
  -output string
        specify the output .kdbx file
```

Example
=======

```
> .\kpchpasswd -input test.kdbx -output test3.kdbx
```

How to build
============

1. install depends: `go get "github.com/tobischo/gokeepasslib" "golang.org/x/crypto/ssh/terminal"`
2. `go build`


Known Issues
============

1. `terminal.ReadPassword` not work in Windows.
