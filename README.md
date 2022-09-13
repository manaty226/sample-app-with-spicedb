# SpiceDBを利用したAPIの認可制御のサンプル

SpiceDBで認可制御するWeb APIのサンプル．
本コードは[ブログ記事](https://zenn.dev/manaty226/articles/71bee4c1a02761)に関連しています．

## エンドポイントとメソッド
本APIはブログの作成．取得，編集が行える．それぞれのエンドポイントとHTTPメソッドは次のとおり．
| エンドポイント | メソッド | 概要 |
| -----------  | -------|-----|
| /blogs  | POST | ブログの作成 |
| /blogs/:id | GET | ブログの取得 |
| /blogs/:id | PUT | ブログの編集 |

## ユーザ認証
ユーザはあらかじめ`.config/user.csv`に記載しておく必要がある．各APIはBASIC認証で保護されている．

