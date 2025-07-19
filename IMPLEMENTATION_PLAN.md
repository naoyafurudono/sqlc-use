# sqlc-use 実装計画

## 技術スタック
- **言語**: Go
- **プラグインタイプ**: プロセスプラグイン
- **依存関係**:
  - github.com/sqlc-dev/plugin-sdk-go
  - github.com/pganalyze/pg_query_go/v5 (PostgreSQL用)
  - その他必要に応じて追加

## プロジェクト構造
```
sqlc-use/
├── cmd/
│   └── sqlc-use/
│       └── main.go          # エントリーポイント
├── internal/
│   ├── plugin/
│   │   └── plugin.go        # プラグインインターフェース実装
│   ├── analyzer/
│   │   ├── analyzer.go      # SQL解析インターフェース
│   │   ├── postgres.go      # PostgreSQL専用解析
│   │   ├── mysql.go         # MySQL専用解析
│   │   └── sqlite.go        # SQLite専用解析
│   ├── formatter/
│   │   └── json.go          # JSON出力フォーマッター
│   └── models/
│       └── types.go         # 共通データ型定義
├── go.mod
├── go.sum
├── Makefile
└── examples/
    ├── sqlc.yaml            # 設定例
    ├── schema.sql           # スキーマ例
    └── query.sql            # クエリ例
```

## 実装ステップ

### Phase 1: 基本実装（1週目）
1. **プロジェクトセットアップ**
   - go.mod初期化
   - 基本的なディレクトリ構造作成
   - Makefile作成

2. **プラグインインターフェース実装**
   - sqlc plugin SDKを使った基本実装
   - stdin/stdout通信の確立
   - エラーハンドリング

3. **基本的なSQL解析**
   - SELECT文の解析
   - テーブル名抽出
   - 簡単なJSONフォーマット出力

### Phase 2: 機能拡張（2週目）
1. **SQL操作の完全サポート**
   - INSERT, UPDATE, DELETE対応
   - JOIN解析
   - FROM句のサブクエリ対応

2. **複数データベース対応**
   - PostgreSQL専用パーサー
   - MySQL対応
   - SQLite対応

3. **テストスイート構築**
   - ユニットテスト
   - 統合テスト
   - E2Eテスト

### Phase 3: 品質向上（3週目）
1. **エラーハンドリング強化**
   - 詳細なエラーメッセージ
   - デバッグモード

2. **パフォーマンス最適化**
   - 大規模クエリの処理
   - メモリ使用量の最適化

3. **ドキュメント整備**
   - README.md
   - API仕様書
   - 使用例

## 開発フロー
1. **ブランチ戦略**
   - main: 安定版
   - develop: 開発版
   - feature/*: 機能開発

2. **テスト駆動開発**
   - テストファースト
   - CI/CD設定（GitHub Actions）

3. **リリース計画**
   - v0.1.0: 基本機能（Phase 1完了時）
   - v0.2.0: 機能完成（Phase 2完了時）
   - v1.0.0: 本番対応（Phase 3完了時）

## 次のアクション
1. go.mod初期化とプロジェクト構造作成
2. sqlc plugin SDKの調査とサンプル実装
3. 基本的なプラグイン動作確認