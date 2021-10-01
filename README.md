# Gouuid  [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/knesklab/util/blob/master/LICENSE)

A golang utility to work with `uuid` types.

## Installation
```
go get github.com/ubgo/gouuid
```

### DefaultIfEmpty(val *uuid.UUID) uuid.UUID
FuncName will return the current function's name.
Check if pointer of uuid is empty then return the default 00000000-0000-0000-0000-000000000000
```go
gouuid.DefaultIfEmpty(nil) // output: 00000000-0000-0000-0000-000000000000

uid := uuid.New()
gouuid.DefaultIfEmpty(&uid) // output: some uid b7729c88-47e9-42a7-92d3-3e6bcc585f73
```

### ParseToDefault(val string) uuid.UUID
Parse string as uuid and return default if error
```go
ParseToDefault("not_uid") // output: 00000000-0000-0000-0000-000000000000

uid := uuid.New().String()
result = ParseToDefault(uid) // output: some uid b7729c88-47e9-42a7-92d3-3e6bcc585f73
```


### IndexOf(slice []uuid.UUID, val uuid.UUID) (int, bool) 
Check if value exist on a given uuid slice then return the index
```go
ids := []uuid.UUID{ParseToDefault("b7729c88-47e9-42a7-92d3-3e6bcc585f73"), ParseToDefault("83adb35a-847a-4962-8e09-8311a45dc2a2"), ParseToDefault("d9d65dfc-4643-44ab-920f-c564259fd96c")}
result, _ := IndexOf(ids, ParseToDefault("83adb35a-847a-4962-8e09-8311a45dc2a2"))
// output: 1, true
```


## Contribute

If you would like to contribute to the project, please fork it and send us a pull request.  Please add tests
for any new features or bug fixes.

## Stay in touch

* Author - [Aman Khanakia](https://twitter.com/mrkhanakia)
* Website - [https://khanakia.com](https://khanakia.com/)

## License

gouuid is [MIT licensed](LICENSE).
