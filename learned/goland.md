## golang memo

### go module 導入
リポジトリに含まれる複数のパッケージをまとめてモジュール化しバージョニングしてくれる
- go.mod
  依存関係の管理（package.jsonのようなもの）
- go.sum
  依存モジュールのチェックサム(データに誤りがないかどうかをチェックす)の管理
  
$GO111MODULEでGo moduleを使うか制御している
デフォルトではautoになっている

- on にされていれば Go modules を使う (module-aware mode)
- off にされていれば従来どおり $GOPATH を使う (GOPATH mode)
- auto の場合 $GOPATH/src の外に対象のリポジトリがあり、 go.mod が存在する場合は module-aware mode、そうでない場合は GOPATH mode

- 使い方
```
export GO111MODULE=on でgomoduleを使う設定をする
go mod init で、初期化する
go buildなどのビルドコマンドで、依存モジュールを自動インストールする
go list -m all で、現在の依存モジュールを表示する
go get で、依存モジュールの追加やバージョンアップを行う
go mod tidy で、使われていない依存モジュールを削除する
```

### golang 記法

戻り値の破棄
```go
func multi() (int, int) {
  return 2, 3
}

a, _ = multi()
a// 2
二個目の戻り値破棄
```

例外機構がないgolangならではのエラーハンドリング処理
```go
str := "123"
w, err := str.Atoi(str)
if err != nil {
    // エラーハンドリング
}
```

関数戻り値が複数
```go
func multi() (int, int) {
  return 2, 3
}
a, b := multi()
a // 2
b // 3
```

関数を返り値にしたり、変数に関数を代入する方法
```go
func someThing() func() {
  return func() {
    fmt.println("Hello")
  }
}

hello := someTHing()
hello() // Hello
```

無名関数
```go
f := func(x, y int) int { return x + y }
f(2, 3) // 5
```
