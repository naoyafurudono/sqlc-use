# sqlc-use プラグイン設計提案書

## 概要
sqlc-useは、SQLクエリが使用するテーブルと操作を解析し、JSONフォーマットで出力するsqlcプラグインです。本設計は「A Philosophy of Software Design」の原則に従います。

## 設計原則

### 1. Deep Modules（深いモジュール）
シンプルなインターフェースと複雑な実装の隠蔽

### 2. 複雑性の最小化
各コンポーネントの責任を明確に分離し、相互依存を最小限に

### 3. 戦略的プログラミング
長期的な保守性を考慮した設計

## アーキテクチャ

### 実装方針
- **言語**: Go
- **プラグインタイプ**: プロセスプラグイン（標準入出力経由の通信）
- **依存関係**: sqlc plugin SDK for Go

### コアモジュール

#### 1. Plugin Interface（プラグインインターフェース）
```go
// シンプルな公開インターフェース
type Plugin interface {
    Generate(req *plugin.GenerateRequest) (*plugin.GenerateResponse, error)
}
```
- **責任**: sqlcとの通信プロトコル処理
- **隠蔽**: protobufのシリアライゼーション、エラーハンドリング
- **実装**: sqlc plugin SDK for Goを使用

#### 2. SQL Parser（SQLパーサー）
```go
// 深いモジュール: シンプルなインターフェースで複雑なSQL解析を隠蔽
type QueryAnalyzer interface {
    Analyze(query Query) ([]TableOperation, error)
}

type TableOperation struct {
    Operation string // "select", "insert", "update", "delete"
    Table     string
}
```
- **責任**: SQL文の解析とテーブル操作の抽出
- **隠蔽**: SQL構文解析の複雑性、異なるSQL方言への対応

#### 3. Output Formatter（出力フォーマッター）
```go
// 拡張可能な出力フォーマット
type OutputFormatter interface {
    Format(operations map[string][]TableOperation) ([]byte, error)
}
```
- **責任**: 解析結果のJSON形式への変換
- **隠蔽**: フォーマット詳細、インデント処理

### 情報隠蔽の原則

1. **SQL解析の複雑性を隠蔽**
   - ユーザーはSQL方言の違いを意識する必要がない
   - JOIN、サブクエリ、CTEの処理は内部で完結

2. **設定の簡潔性**
   ```yaml
   version: '2'
   plugins:
     - name: sqlc-use
       process:
         cmd: sqlc-use
   sql:
     - schema: schema.sql
       queries: query.sql
       engine: postgresql
       codegen:
         - out: gen
           plugin: sqlc-use
           options:
             format: json  # デフォルト設定で動作
   ```

3. **エラーの抽象化**
   - 内部エラーを意味のあるユーザーメッセージに変換
   - デバッグ情報は必要時のみ公開

## 実装計画

### Phase 1: 基本機能
1. プラグインインターフェースの実装
2. 基本的なSQL解析（SELECT, INSERT, UPDATE, DELETE）
3. JSON出力の実装

### Phase 2: 高度な機能
1. JOIN解析の強化
2. サブクエリサポート
3. CTE（Common Table Expression）対応

### Phase 3: 拡張性
1. カスタム出力フォーマット
2. フィルタリング機能
3. パフォーマンス最適化

## 複雑性の管理

### 認知的負荷の軽減
- 各モジュールは単一の責任を持つ
- インターフェースは最小限の知識で使用可能
- ドキュメントは「なぜ」を説明（「何を」だけでなく）

### テスタビリティ
- 各モジュールは独立してテスト可能
- モックを使った統合テスト
- 実際のSQLクエリを使用したE2Eテスト

## セキュリティ考慮事項
- プロセスプラグインのため、信頼できる環境でのみ使用
- 外部リソースへのアクセスは最小限に制限
- 入力検証の徹底
- SQLインジェクション対策（解析のみで実行はしない）

## まとめ
本設計は、シンプルなインターフェースで複雑な機能を提供し、長期的な保守性と拡張性を確保します。「A Philosophy of Software Design」の原則に従い、ユーザーが最小限の知識で最大限の価値を得られるプラグインを目指します。