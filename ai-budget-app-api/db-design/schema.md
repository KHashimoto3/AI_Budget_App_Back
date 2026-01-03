# テーブル定義

## public（共通スキーマ）

### users（ユーザー）

| カラム名      | 型名      | 説明                                                    |
| ------------- | --------- | ------------------------------------------------------- |
| id            | UUID (PK) | ユーザー ID                                             |
| firebase_uid  | TEXT      | Firebase UID                                            |
| name          | VARCHAR   | ユーザー名                                              |
| disp_name     | VARCHAR   | 表示名                                                  |
| email         | VARCHAR   | メールアドレス（ログイン用）                            |
| password_hash | TEXT      | パスワード（ハッシュ化済み）                            |
| account_type  | INTEGER   | アカウントタイプ（0: 通常ログイン、1: Google ログイン） |
| created_at    | TIMESTAMP | アカウント作成日時（UTC）                               |
| login_at      | TIMESTAMP | 最終ログイン日時（UTC）                                 |
| deleted_flag  | BOOL      | 削除フラグ                                              |

### expenses（支出記録）

| カラム名     | 型名       | 説明                                 |
| ------------ | ---------- | ------------------------------------ |
| id           | UUID (PK)  | 支出 ID                              |
| user_id      | UUID (FK)  | ユーザー ID                          |
| expense_date | DATE       | 支出日（UTC）                        |
| amount       | INTEGER    | 金額（円）                           |
| genres_id    | UUID (FK)  | `genres.id` への外部キー             |
| shop_name    | VARCHAR    | お店の名前                           |
| memo         | TEXT       | メモ（NULL 可）                      |
| input_type   | VARCHAR    | 入力形式（`manual` or `ocr`）        |
| image_id     | UUID（FK） | images.id への外部キー（OCR 時のみ） |
| created_at   | TIMESTAMP  | 登録日時（UTC）                      |
| updated_at   | TIMESTAMP  | 更新日時（UTC）                      |

### images（画像）

| カラム名   | 型名      | 説明                         |
| ---------- | --------- | ---------------------------- |
| id         | UUID (PK) | 画像 ID                      |
| user_id    | UUID (FK) | ユーザー ID                  |
| filename   | VARCHAR   | ファイル名                   |
| file_path  | VARCHAR   | ファイルへのパス（保存先）   |
| parsed_at  | TIMESTAMP | 解析実行日時（OCR 済の場合） |
| created_at | TIMESTAMP | アップロード日時             |

### genres（ジャンル）

| カラム名 | 型名      | 説明        |
| -------- | --------- | ----------- |
| id       | UUID (PK) | ジャンル ID |
| name     | VARCHAR   | ジャンル名  |

※ジャンル名は v2.0 以降に追加。今は作成しない。
