# Development Log

## 2025-07-19

### sqlcプラグインアーキテクチャの調査完了
- sqlcはVersion 2の設定ファイルでプラグインをサポート
- プラグインタイプ：WASM（推奨、サンドボックス環境）とProcess（信頼できるコードのみ）
- プロトコル：Protocol Bufferメッセージによるstdin/stdout通信
- gRPCサービスインターフェースでメソッドサポート

### 次のステップ
- A Philosophy of Software Designの原則に基づいた設計提案書の作成