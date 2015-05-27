# mdpp
mdpp searches .mobileprovision file from "~/Library/MobileDevice/Provisioning Profiles/".

## Usage
```bash
$ mdpp search [Provisioning Profiles Name]
```

```bash
$ mdpp search FooBarApp
/Users/hoge/Library/MobileDevice/Provisioning Profiles/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.mobileprovision
```

## Install
To install, use `go get`:

```bash
$ go get -d github.com/keygx/mdpp
```
```bash
$ cd $GOPATH/src/github.com/keygx/mdpp
```
```bash
$ go install
```

## License
mdpp is released under the MIT license. See LICENSE for details.

## Author
Yukihiko Kagiyama (keygx) <https://twitter.com/keygx>
