# api 仕様

## 認証

### POST /auth/signup

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

ボディ

```json
{
  "company_id": "string | undefined",
  "display_name": "string",
  "icon_url": "string | undefined"
}
```

#### レスポンス

```json
{
  "company_id": "string | undefined",
  "user_name": "string",
  "display_name": "string",
  "icon_url": "string | undefined"
}
```

#### 備考

login 用のエンドポイントは必要ない。login を「id トークンを発行する処理」とするなら、この「id トークンを発行する処理」は
Firebase 側でやってくれるから。

uid をフロントから送る必要はなさそう。uid を verify する関数の戻り値で uid を取得できるため。
https://github.com/firebase/firebase-admin-go/blob/master/auth/auth.go#L259

下の2つの処理を signup 時に行い、トークンの有効性を確認する。
ユーザ登録時はセキュリティを高めたいため。

1. Firebase Admin SDK を使用して ID トークンを確認する
   https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja#verify_id_tokens_using_the_firebase_admin_sdk

2. SDK で ID トークンの取り消しを検出する
   https://firebase.google.com/docs/auth/admin/manage-sessions?hl=ja#detect_id_token_revocation_in_the_sdk

signup 以外の任意のエンドポイントを叩く時はトークンの有効性を確認するだけ、つまり「Firebase Admin SDK を使用して ID
トークンを確認する」処理しか行わないことにする。

## ホーム

### GET /achievement-rate

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "garbage_collection_rate": "number"
}
```

#### 備考

トップページだと仮定

サーバの現在の日時を持ってきて、現在の月を算出する。10月だった場合、10月のイベント参加履歴を db
から持ってくる。参加履歴からそのユーザがその月に何時間ゴミ拾いをしたかを求められる。

参考:

- https://www.asahi.com/articles/ASLDV5WRQLDVPLBJ00D.html
- https://www3.nhk.or.jp/news/html/20230902/k10014182011000.html

1時間あたりにとれるゴミの量を算出する。（1 kg と仮定）
→ ここが出ていないから問題

1時間あたりにとれるゴミの量とその月に何時間ゴミ拾いしたかをかけることでユーザがその月にどれくらいゴミを拾ったかを算出できる。

1ヶ月あたりに最大に拾えるゴミの量を算出する。これを使って割合を算出する。

### GET /event/recommendation?latitude="number"&longitude="number"&limit="number"&offset="number"

limit は1~10 の範囲で指定できる。指定しない場合は 10 になる。

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "events": [
    {
      "id": "string",
      "title": "string",
      "host_company_name": "string",
      "address": "string",
      "participant_count": "number",
      "unit_price": "number",
      "will_start_at": "Date",
      "will_complete_at": "Date",
      "application_deadline": "Date"
    },
    ...
  ]
}
```

#### 備考

現時点で latitude、longitude クエリが使用されるのはユーザがまだイベントに1回も参加していない時のみです。
つまり、基本的には latitude、longitude クエリをつけてリクエストを送ってほしいです。フロント側ではユーザがイベントに参加したことがあるかの有無を判別できないためです。
もし、ユーザにイベントの参加経験がない、かつ latitude、longitude クエリの付与もない場合は events テーブルに入っているデータを順番に返します。

ユーザがまだイベントに1回も参加していない時、ユーザの位置情報を使ってユーザに近いイベントの情報を取得。
ユーザがイベントに1回は参加している時、そのユーザが前回参加したイベントの開催地から近いイベントの情報を取得します。

### GET /event/schedule

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "events": [
    {
      "id": "string",
      "title": "string",
      "host_company_name": "string",
      "address": "string",
      "participant_count": "number",
      "unit_price": "number",
      "will_start_at": "Date",
      "will_complete_at": "Date",
      "application_deadline": "Date"
    },
    ...
  ]
}
```

#### 備考

ユーザが参加予定のイベントを直近順に取得。

## 検索

### GET /event/search?keyword="string"&min_unit_price="number"&will_start_at="Date"&limit="number"&offset="number"

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "events": [
    {
      "id": "string",
      "title": "string",
      "host_company_name": "string",
      "address": "string",
      "participant_count": "number",
      "unit_price": "number",
      "will_start_at": "Date",
      "will_complete_at": "Date",
      "application_deadline": "Date"
    },
    ...
  ]
}
```

#### 備考

host_company_name は db 側では like 句を使って検索する。
min_unit_price で指定した単価以上のイベントを返します。
is_not_full_house を true にすることで定員に空きがあるイベントのみを返すようにします。
（定員に空きがあるかの絞り込みをフロント側で行うような場合は、そのイベントが定員に空きがあるかの状態も一緒に返してあげるようにします。）

### GET /event/order-recommendation?address="string"&start_at="Date"&complete_at="Date"&interval_minute="number"

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "events": [
    {
      "id": "string",
      "title": "string",
      "host_company_name": "string",
      "address": "string",
      "participant_count": "number",
      "unit_price": "number",
      "will_start_at": "Date",
      "will_complete_at": "Date",
      "application_deadline": "Date"
    },
    ...
  ]
}
```

#### 備考

1日に複数のイベントを回る際の、おすすめのイベント参加順の提案をする機能です。
参照: https://www.notion.so/d53b9d98864a4629aac88f65c77df930

start_at と complete_at で指定した時間の間に開催されているイベントで、interval_minute
で指定した分数の間隔でイベントを回るときのおすすめのイベント参加順を配列で返します。

## イベント詳細

### GET /event/:id

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "id": "string",
  "title": "string",
  "host_company_name": "string",
  "description": "string",
  "address": "string",
  "latitude": "number",
  "longitude": "number",
  "participant_count": "number",
  "unit_price": "number",
  "will_start_at": "Date",
  "will_complete_at": "Date",
  "application_deadline": "Date",
  "leader_name": "string | undefined",
  "started_at": "Date | undefined",
  "completed_at": "Date | undefined"
}
```

### GET /event/:id/participant

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "participants": [
    {
      "user_id": "string",
      "status": "string"
    },
    ...
  ]
}
```

##### 備考

status には次の3つの文字列のどれかが入ります。

- completed（イベントに参加し、報告も完了している）
- not_reported（報告がまだ完了していない。）
- absent（欠席）

## 参加履歴

### GET /event/participation-history?limit="number"&offset="number"

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "events": [
    {
      "id": "string",
      "title": "string",
      "host_company_name": "string",
      "address": "string",
      "participant_count": "number",
      "unit_price": "number",
      "will_start_at": "Date",
      "will_complete_at": "Date",
      "status": "string"
    },
    ...
  ]
}
```

#### 備考

status には次の3つの文字列のどれかが入ります。

- completed（イベントに参加し、報告も完了している）
- not_reported（報告がまだ完了していない。）
- absent（欠席）

## アカウント設定

### GET /user/:id

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "id": "string",
  "company_id": "string | undefined",
  "user_name": "string",
  "display_name": "string",
  "sex": "string",
  "age": "number",
  "icon_url": "string"
}
```

#### 備考

性別はフロントに一応返しておくけど表示するかは自由です。

## イベント開始〜開催中〜終了

### POST /event/:id/start

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "id": "string",
  "title": "string",
  "started_at": "Date"
}
```

#### 備考

そのイベントのリーダのみが叩けるイベントを開始するエンドポイントです。

### POST /event/:id/complete

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

ボディ

```json
{
  "proof_participants_image_url": "string",
  "proof_garbage_image_url": "string",
  "report": "string"
}
```

#### レスポンス

```json
{
  "id": "string",
  "title": "string",
  "completed_at": "Date",
  "proof_participants_image_url": "string",
  "proof_garbage_image_url": "string",
  "report": "string"
}
```

#### 備考

そのイベントのリーダのみが叩けるイベントを完了するエンドポイントです。
proof_participants_image_url は参加者の写真を撮った画像のオブジェクト URL です。proof_garbage_image_url
は拾ったゴミの写真を撮った画像のオブジェクト URL です。
report にはそのイベントの様子などを書いてもらいます。態度の悪い人や欠席者等もこの report に書いてもらいます。

### POST /event/:id/report

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

ボディ

```json
{
  "latitude": "number",
  "longitude": "number"
}
```

#### レスポンス

```json
{
  "id": "string",
  "title": "string",
  "latitude": "number",
  "longitude": "number"
}
```

#### 備考

そのイベントに参加しているユーザが叩けるイベントの報告をするエンドポイントです。
リクエストボディで送られるユーザーの現在位置がイベント場所からずれていたらエラー(400 Bad Request)を返します。

## イベント管理

### GET /user/:id/event?limit="number"&offset="number"

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "events": [
    {
      "id": "string",
      "title": "string",
      "host_company_name": "string",
      "address": "string",
      "participant_count": "number",
      "unit_price": "number",
      "will_start_at": "Date",
      "will_complete_at": "Date",
      "application_deadline": "Date",
      "leader_name": "string | undefined"
    },
    ...
  ]
}
```

#### 備考

そのユーザが主催しているイベント一覧を取得するエンドポイントです。
現時点では会社に所属しているユーザのみが主催できるようにしているため、company_id が undefined のユーザはこのエンドポイントを叩けません。

## イベント詳細

### GET /user/:id/event/:id

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

#### レスポンス

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "address": "string",
  "latitude": "number",
  "longitude": "number",
  "participant_count": "number",
  "unit_price": "number",
  "will_start_at": "Date",
  "will_complete_at": "Date",
  "application_deadline": "Date",
  "leader": "string | undefined"
}
```

## イベント作成、編集

### POST /user/:id/event

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

ボディ

```json
{
  "title": "string",
  "description": "string",
  "address": "string",
  "latitude": "number",
  "longitude": "number",
  "participant_count": "number",
  "unit_price": "number",
  "will_start_at": "Date",
  "will_complete_at": "Date",
  "application_deadline": "Date"
}
```

#### レスポンス

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "address": "string",
  "latitude": "number",
  "longitude": "number",
  "participant_count": "number",
  "unit_price": "number",
  "will_start_at": "Date",
  "will_complete_at": "Date",
  "application_deadline": "Date"
}
```

#### 備考

新たにイベントを主催するエンドポイントです。
現時点では会社に所属しているユーザのみが主催できるようにしているため、company_id が undefined のユーザはこのエンドポイントを叩けません。

### PATCH /user/:id/event/:id

#### リクエスト

ヘッダ

```
Authorization: Barer "id_token"
```

ボディ

```json
{
  "title": "string | undefined",
  "description": "string | undefined",
  "address": "string | undefined",
  "latitude": "number | undefined",
  "longitude": "number | undefined",
  "participant_count": "number | undefined",
  "unit_price": "number | undefined",
  "will_start_at": "Date | undefined",
  "will_complete_at": "Date | undefined",
  "application_deadline": "Date | undefined",
  "leader": "string | undefined"
}
```

#### レスポンス

```json
{
  "title": "string | undefined",
  "description": "string | undefined",
  "address": "string | undefined",
  "latitude": "number | undefined",
  "longitude": "number | undefined",
  "participant_count": "number | undefined",
  "unit_price": "number | undefined",
  "will_start_at": "Date | undefined",
  "will_complete_at": "Date | undefined",
  "application_deadline": "Date | undefined",
  "leader": "string | undefined"
}
```

#### 備考

新たにイベントを主催するエンドポイントです。
現時点では会社に所属しているユーザのみが主催できるようにしているため、company_id が undefined のユーザはこのエンドポイントを叩けません。
leader はそのイベントのリーダのユーザ ID です。leader にユーザ ID を指定すると、そのユーザがリーダになります。
