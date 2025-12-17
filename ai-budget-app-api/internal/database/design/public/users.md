# users

| カラム名      | 型名      | 説明                                                    |
| ------------- | --------- | ------------------------------------------------------- |
| id            | UUID (PK) | ユーザー ID                                             |
| name          | VARCHAR   | ユーザー名                                              |
| disp_name     | VARCHAR   | 表示名                                                  |
| email         | VARCHAR   | メールアドレス（ログイン用）                            |
| password_hash | TEXT      | パスワード（ハッシュ化済み）                            |
| schema_name   | TEXT      | 各ユーザー用スキーマ名（例：`user_abc123`）             |
| account_type  | INTEGER   | アカウントタイプ（0: 通常ログイン、1: Google ログイン） |
| created_at    | TIMESTAMP | アカウント作成日時（UTC）                               |
| login_at      | TIMESTAMP | 最終ログイン日時（UTC）                                 |
| deleted_flag  | BOOL      | 削除フラグ                                              |
