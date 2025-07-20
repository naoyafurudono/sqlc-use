# sqlc-use

sqlc-useはsqlcプラグインで、それぞれのクエリがテーブルに対してどのような操作をするかを表現するJSONを生成します。

## インストール

### リリースバイナリから

最新のリリースは[GitHub Releases](https://github.com/naoyafurudono/sqlc-use/releases)からダウンロードできます。

```bash
# Linux (AMD64)
curl -L https://github.com/naoyafurudono/sqlc-use/releases/latest/download/sqlc-use-linux-amd64 -o sqlc-use
chmod +x sqlc-use

# macOS (Apple Silicon)
curl -L https://github.com/naoyafurudono/sqlc-use/releases/latest/download/sqlc-use-darwin-arm64 -o sqlc-use
chmod +x sqlc-use
```

### ソースからビルド

```bash
go install github.com/naoyafurudono/sqlc-use/cmd/sqlc-use@latest
```

## 仕様

### 出力ファイル

- **ファイル名**: `query-table-operations.json`
- **出力先**: `sqlc.yaml`の`codegen.out`で指定したディレクトリ

### 出力フォーマット

```json
{
  "[package.]クエリ名": [
    {
      "operation": "操作タイプ",
      "table": "テーブル名"
    }
  ]
}
```

- **操作タイプ**: `select`, `insert`, `update`, `delete`
- **テーブル名**: SQL内で操作されているテーブル名
- JOINで複数テーブルを操作する場合、各テーブルが個別に記録されます
- **JSON Schema**: 出力形式の詳細な定義は[schema/query-table-operations.schema.json](schema/query-table-operations.schema.json)を参照

### 設定方法

`sqlc.yaml`:

```yaml
version: "2"
plugins:
  - name: sqlc-use
    process:
      cmd: sqlc-use # パスまたはコマンド名
sql:
  - schema: schema.sql
    queries: query.sql
    engine: mysql # 現在はmysqlのみサポート
    codegen:
      - out: gen # 出力ディレクトリ
        plugin: sqlc-use
```

### オプション

プラグインオプションは `sqlc.yaml` の `options` フィールドで指定できます：

```yaml
codegen:
  - out: gen
    plugin: sqlc-use
    options:
      package: db  # クエリ名のプレフィックス（オプション）
      format: json # 出力フォーマット（デフォルト: json）
```

- **package**: 指定した場合、出力されるクエリ名が `package.QueryName` の形式になります
- **format**: 現在は `json` のみサポート

## 例

### 入力

```sql
-- name: ListOrganizationMember :many
select user.* from user
inner join member on user.id = member.user_id
inner join organization on organization.id = member.organization_id
where organization.name = ?;

-- name: AddMember :exec
insert into member (user_id, organization_id) values (?, ?);

-- name: RemoveMember :exec
delete from member where user_id = ? and organization_id = ?;
```

### 出力

`gen/query-table-operations.json`:

```json
{
  "db.ListOrganizationMember": [
    {
      "operation": "select",
      "table": "user"
    },
    {
      "operation": "select",
      "table": "member"
    },
    {
      "operation": "select",
      "table": "organization"
    }
  ],
  "db.AddMember": [
    {
      "operation": "insert",
      "table": "member"
    }
  ],
  "db.RemoveMember": [
    {
      "operation": "delete",
      "table": "member"
    }
  ]
}
```

## 使い方

1. `sqlc.yaml`に上記の設定を追加
2. `sqlc generate`を実行
3. 指定した出力ディレクトリに`query-table-operations.json`が生成されます

## バージョニング

このプロジェクトは[セマンティックバージョニング](https://semver.org/lang/ja/)に従います。

- **メジャーバージョン**: 後方互換性のない変更
- **マイナーバージョン**: 後方互換性のある機能追加
- **パッチバージョン**: 後方互換性のあるバグ修正

バージョンタグをプッシュすると、GitHub Actionsが自動的にリリースビルドを作成します。

## ライセンス

MITライセンス
