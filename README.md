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
- アプリケーションのエントリーポイントとなるmain.goを管理する

### internal
- 外部モジュールに公開されないパッケージ
- アプリケーションで利用するロジックは基本的にinternal配下で管理する
- Spoconのアプリケーションは他のモジュールに公開しない想定
- internal配下にアプリケーションのロジックを置くことで、他のモジュールから参照されないことを前提に開発を進めることができ、後方互換を気にすることなく変更できる作りとする

### internal/domain

#### model
- 値オブジェクトやエンティティなどのドメインモデルを管理する

#### repository
- リポジトリのインターフェースを管理する

#### service
- ドメイン特有のロジックを管理する

### internal/handler
- Controllerと同様の役割
- クライアントから受け取ったデータのバリデーション、クライアントにステータスなどを含めたデータを返す処理を実装する

#### internal/infra
- データベースへのアクセスを管理する

### internal/openapi
- openapi.yamlから生成したコードを管理

### internal/usecase
- handlerから呼び出される
- ドメイン層の処理を呼び出し、ユースケースの実現を行う

### openapi
- APIドキュメントを管理
- このドキュメントからAPIインターフェースのコード生成を行う

### tools
- Go製ツールのインポート定義を管理する
- 運用上で利用するツールを開発者間で同じバージョンで利用するために、本パーッケージに定義しておく

## マイグレーションファイル作成手順

- ファイル名を指定して、以下のコマンドを実行する
    ```
    go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir database/migrations -seq {ファイル名}
    ```
- 実行後、database/migrations配下に指定した名前でファイルが作成されていることを確認する
- 〇〇.up.sqlに追加したいSQL、〇〇.down.sqlに追加した分を戻すSQLを記載する