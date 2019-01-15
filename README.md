# slack-clockout
IoT Enterprise Button から Slack に退勤時間を投稿する
* シングルクリックで退勤時間を投稿する(稼働報告書作成用のメモ)
* ダブルクリックで最終退勤者報告を投稿する

## コンパイル
```bash
$ GOOS=linux go build -o myclockout
$ zip myclockout.zip myclockout
```

## 環境変数
|キー|値|備考|
|:--|:--|:--|
|SlackURL|https://hooks.slack.com/services/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|Slack Webhook URL|
|PrivateChannel|@taaazyyy|シングルクリックで退勤時間を投稿するチャンネル|
|PublicChannel|#random|ダブルクリックで最終退勤者を投稿するチャンネル|
|Name|田島 太郎|投稿者名|
|LastName|田島|最終退勤者報告|

## テストイベント
* [SINGLE](test/testevent_single.json)
* [DOUBLE](test/testevent_double.json)
* [LONG](test/testevent_long.json)