# expenses（支出記録）

| カラム名     | 型名       | 説明                                 |
| ------------ | ---------- | ------------------------------------ |
| id           | UUID (PK)  | 支出 ID                              |
| expense_date | DATE       | 支出日（UTC）                        |
| amount       | INTEGER    | 金額（円）                           |
| genres_id    | UUID (FK)  | `genres.id` への外部キー             |
| shop_name    | VARCHAR    | お店の名前                           |
| memo         | TEXT       | メモ（NULL 可）                      |
| input_type   | VARCHAR    | 入力形式（`manual` or `ocr`）        |
| image_id     | UUID（FK） | images.id への外部キー（OCR 時のみ） |
| created_at   | TIMESTAMP  | 登録日時（UTC）                      |
| updated_at   | TIMESTAMP  | 更新日時（UTC）                      |
