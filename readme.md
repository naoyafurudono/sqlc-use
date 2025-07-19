# sqlc-use

sqlc-useはsqlcプラグインで、それぞれのクエリがテーブルに対してどのような操作をするかを表現するJSONを生成します。

## 仕様

### 出力ファイル
- **ファイル名**: `query_usage.json`
- **出力先**: `sqlc.yaml`の`codegen.out`で指定したディレクトリ

### 出力フォーマット
```json
{
  "クエリ名": [
    {
      "operation": "操作タイプ",
      "table": "テーブル名"
    }
  ]
}
```

- **操作タイプ**: `select`, `insert`, `update`, `delete`
- **テーブル名**: SQL内で参照されているテーブル名
- JOINで複数テーブルを参照する場合、各テーブルが個別に記録されます

### 設定方法

`sqlc.yaml`:
```yaml
version: '2'
plugins:
  - name: sqlc-use
    process:
      cmd: sqlc-use  # パスまたはコマンド名
sql:
  - schema: schema.sql
    queries: query.sql
    engine: mysql      # 現在はmysqlのみサポート
    codegen:
      - out: gen       # 出力ディレクトリ
        plugin: sqlc-use
```

### オプション
現在、プラグイン固有のオプションはありません。

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

`gen/query_usage.json`:
```json
{
  "ListOrganizationMember": [
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
  "AddMember": [
    {
      "operation": "insert",
      "table": "member"
    }
  ],
  "RemoveMember": [
    {
      "operation": "delete",
      "table": "member"
    }
  ]
}
```

## インストール

```bash
go install github.com/naoyafurudono/sqlc-use/cmd/sqlc-use@latest
```

## 使い方

1. `sqlc.yaml`に上記の設定を追加
2. `sqlc generate`を実行
3. 指定した出力ディレクトリに`query_usage.json`が生成されます
