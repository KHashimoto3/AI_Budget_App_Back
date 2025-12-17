# images（画像）

| カラム名   | 型名      | 説明                         |
| ---------- | --------- | ---------------------------- |
| id         | UUID (PK) | 画像 ID                      |
| filename   | VARCHAR   | ファイル名                   |
| file_path  | VARCHAR   | ファイルへのパス（保存先）   |
| parsed_at  | TIMESTAMP | 解析実行日時（OCR 済の場合） |
| created_at | TIMESTAMP | アップロード日時             |
