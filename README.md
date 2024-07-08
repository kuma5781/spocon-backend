# go-restapi-sample


## ディレクトリ構成

```
.
└── spocon-backend/
    ├── cmd
    ├── internal/
    │   ├── domain/
    │   │   ├── model
    │   │   ├── repository
    │   │   └── service
    │   ├── handler
    │   ├── infra
    │   ├── openapi
    │   └── usecase
    ├── openapi
    └── tools
```

### cmd
- アプリケーションの実行ファイルであるmain.goのみ管理する
- main.goがエントリーポイントとなり、アプリケーション内の処理は他ディレクトリで管理する

### internal
- 外部アプリケーションからインポートできないディレクトリ
- Spoconのバックエンドは、外部アプリケーションからインポートされない想定なので、基本この中で管理する

#### domain
- model
    - 値オブジェクト、エンティティのようなドメインモデルを管理
- repository
    - リポジトリのインターフェース
    - infraディレクトリで実装を管理して、ドメインロジックがinfraに依存しないにrepositoryディレクトリでインターフェースを置く
- service
    - ドメインサービス

#### handler
- APIエンドポイントの実装となるhandlerを管理する
- Controllerと同じような役割
- クライアントから受け取ったデータのバリデーション、クライアントにステータスなどを含めたデータを返す処理を行う

#### infra
- データベースへのアクセスを管理する

#### openapi
- openapi.yamlから生成したコードを管理

#### usecase
- handlerから呼び出され、ユースケースを実現を管理
- ユースケースの実現は適宜ドメイン層を呼び出して、実現する

### openapi
- APIドキュメントを管理
- このドキュメントからAPIインターフェースのコード生成を行う

### tools
- TODO: 要修正
- 各開発者がgo installして、異なるバージョンを入れないように、tools.goで管理して、go mod tidyしたときに同じバージョンのツールがgo.modファイルに記載されるようにするため


## コード整形
プロジェクトのルートディレクトリ配下で以下のコマンドを実行する
```
find . -name "*.go" -exec gofmt -w {} \;
```

## マイグレーションファイル作成手順

- ファイル名を指定して、以下のコマンドを実行する
    ```
    go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir database/migrations -seq {ファイル名}
    ```
- 実行後、database/migrations配下に指定した名前でファイルが作成されていることを確認する
- 〇〇.up.sqlに追加したいSQL、〇〇.down.sqlに追加した分を戻すSQLを記載する