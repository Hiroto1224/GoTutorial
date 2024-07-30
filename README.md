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

.<br>
├── .idea<br>
│   ├── .gitignore<br>
│   ├── modules.xml<br>
│   ├── tutorial.iml<br>
│   └── vcs.xml<br>
├── README.md<br>
└── main<br>
├── greetings<br>
│   ├── go.mod<br>
│   ├── greetings<br>
│   │   ├── greetings.go<br>
│   │   └── greetings_test.go<br>
│   └── hello<br>
│       ├── hello<br>
│       └── hello.go<br>
└── hello<br>
├── go.mod<br>
├── go.sum<br>
└── hello.go<br>

### コマンド一覧

| コマンド       | 使用箇所・用途               |
|------------|-----------------------|
| go run .   | go.modファイルの存在するディレクトリ |
| go test    | テストファイルのあるディレクトリ      |
| go test -v | テストの経過が確認できる          |
| go build   | 実行可能なバイナリファイル作成       |
