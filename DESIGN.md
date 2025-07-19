# sqlc-use プラグイン設計書

## 概要
sqlc-useは、SQLクエリが使用するテーブルと操作を解析し、JSONフォーマットで出力するsqlcプラグインです。

## 技術スタック
- **言語**: Go
- **プラグインタイプ**: プロセスプラグイン
- **主要依存関係**:
  - github.com/sqlc-dev/plugin-sdk-go
  - github.com/pingcap/tidb/parser (MySQL構文解析用)

## アーキテクチャ

### プロジェクト構造
```
sqlc-use/
├── cmd/sqlc-use/          # エントリーポイント
├── internal/
│   ├── plugin/            # プラグインインターフェース
│   ├── analyzer/          # SQL解析ロジック
│   └── formatter/         # 出力フォーマット
└── examples/              # 使用例
```

### コアモジュール

#### 1. Plugin Interface
```go
type Plugin interface {
    Generate(req *plugin.GenerateRequest) (*plugin.GenerateResponse, error)
}
```

#### 2. Query Analyzer
```go
type QueryAnalyzer interface {
    Analyze(query Query) ([]TableOperation, error)
}

type TableOperation struct {
    Operation string // "select", "insert", "update", "delete"
    Table     string
}
```

#### 3. Output Formatter
```go
type OutputFormatter interface {
    Format(operations map[string][]TableOperation) ([]byte, error)
}
```

## 責務分割

### sqlcの責務
- SQLファイルのパースと構文解析
- 型推論と検証
- プラグインシステムの管理
- スキーマ情報の提供

### sqlc-useの責務
- クエリの使用テーブル解析
- 操作タイプ（SELECT/INSERT等）の分類
- JSON形式での出力生成

### データフロー
```
sqlc → (GenerateRequest) → sqlc-use → (GenerateResponse) → sqlc
```

## 設定例
```yaml
version: '2'
plugins:
  - name: sqlc-use
    process:
      cmd: sqlc-use
sql:
  - schema: schema.sql
    queries: query.sql
    engine: mysql
    codegen:
      - out: gen
        plugin: sqlc-use
        options:
          format: json
```

## 実装フェーズ

### Phase 1: 基本機能
- プラグインインターフェース
- 基本的なSQL解析（SELECT/INSERT/UPDATE/DELETE）
- JSON出力

### Phase 2: 高度な機能
- JOIN解析
- サブクエリ対応
- PostgreSQL/SQLite対応追加

### Phase 3: 最適化
- パフォーマンス改善
- エラーハンドリング強化

## 設計原則
- **Deep Modules**: シンプルなインターフェースで複雑な実装を隠蔽
- **単一責任**: 各モジュールは明確な一つの責務を持つ
- **テスタビリティ**: 独立したテストが可能な設計
