## 目次

1. [リポジトリについて](#リポジトリについて)
2. [環境](#環境)
3. [ディレクトリ構成](#ディレクトリ構成)
4. [コマンド一覧](#コマンド一覧)


## リポジトリについて
Go言語の学習を行うプロジェクト

[公式チュートリアル](https://go.dev/doc/tutorial/)にそって作成

## 環境

| 言語・フレームワーク            | バージョン   |
|-----------------------|---------|
| Go                    | 1.22    |

## ディレクトリ構成

<!-- Treコマンドを使ってディレクトリ構成を記載 -->

<pre>
.
├── README.md
└── main
    ├── greetings
    │   ├── go.mod
    │   ├── greetings
    │   │   ├── greetings.go
    │   │   └── greetings_test.go
    │   └── hello
    │       ├── hello
    │       └── hello.go
    ├── hello
    │   ├── go.mod
    │   ├── go.sum
    │   └── hello.go
    └── workspace
        ├── example
        ├── go.work
        └── hello
            ├── go.mod
            ├── go.sum
            └── hello.go
</pre>

## コマンド一覧

| コマンド       | 使用箇所・用途               |
|------------|-----------------------|
| go run .   | go.modファイルの存在するディレクトリ/<br/>mainのコードを実行 |
| go test    | テストファイルのあるディレクトリ/<br/>テストを実行 |
| go test -v | テストファイルのあるディレクトリ/<br/>テストの経過が確認できる |
| go build   | 実行可能なバイナリファイル作成       |
