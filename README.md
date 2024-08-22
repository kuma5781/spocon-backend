# spocon-backend

## ローカル環境の起動準備

- [Docker Compose](https://docs.docker.com/compose/install/)をインストールします
- [task](https://taskfile.dev/installation/)をインストールします

## ローカル環境の起動方法

起動には下記コマンドを実行します。spocon-backendとmysqlがコンテナとしてビルドされ起動します。
```
task up
```
その他、開発で使用するコマンドは`task -a`で確認します。

## 開発環境の準備

開発を行うためには、起動する手順に加えて下記の手順が必要です。

- [Go 言語](https://go.dev/doc/install)をインストールします
- 下記コマンドを実行して、Go製ツールの実行ファイルを`./bin`にダウンロードします
    ```
    task setup
    ```

## ディレクトリ構成

```
spocon-backend/
    ├── cmd
    ├── internal/
    │   ├── app/
    │   ├── domain/
    │   │   ├── model
    │   │   ├── repository
    │   │   └── service
    │   ├── handler
    │   ├── infra
    │   ├── openapi
    │   └── usecase
    ├── openapi
    ├── operation
    ├── testutil
    └── tools
```

### cmd
- アプリケーションのエントリーポイントとなるmain.goを管理する

### internal
- 外部モジュールに公開されないパッケージ
- アプリケーションで利用するロジックは基本的にinternal配下で管理する
- Spoconのアプリケーションは他のモジュールに公開しない想定
- internal配下にアプリケーションのロジックを置くことで、他のモジュールから参照されないことを前提に開発を進めることができ、後方互換を気にすることなく変更できる作りとする

#### internal/app
- アプリケーションの基盤となる設定を管理する

#### internal/domain

- model
    - 値オブジェクトやエンティティなどのドメインモデルを管理する
- repository
    - リポジトリのインターフェースを管理する
- service
    - ドメイン特有のロジックを管理する

#### internal/handler
- Controllerと同様の役割
- クライアントから受け取ったデータのバリデーション、クライアントにステータスなどを含めたデータを返す処理を実装する

#### internal/infra
- データベースへのアクセスを管理する

#### internal/openapi
- openapi.yamlから生成したコードを管理する

#### internal/usecase
- handlerから呼び出される
- ドメイン層の処理を呼び出し、ユースケースの実現を行う

### openapi
- APIドキュメントを管理する
- このドキュメントからAPIインターフェースのコード生成を行う

### operation
- プロダクトコードでは利用しない運用ツールを管理する

### testutil
- プロジェクト配下の全てのテストコードから利用可能な汎用的な処理を管理する

### tools
- Go製ツールのインポート定義を管理する
- 運用上で利用するツールを開発者間で同じバージョンで利用するために、本パーッケージに定義しておく