# GayaON-Server

GayaON! という Web サービスで使用するサーバプログラムです。

フロントエンドのプログラムは[こちら](https://example.com)。

クライアントプログラムは[こちら](https://github.com/isso-719/gaya-on-client)。

## 概要

GayaON! は、プレゼンターの画面上に聴衆のコメントや絵文字をリアルタイムに表示することで、
プレゼンターが聴衆のリアクションを感じ取りやすくなる Web サービスです。

## 使い方

Docker-compose でサーバプログラムを動かします。Docker-compose をインストールしてない場合は[ここ](https://qiita.com/isso_719/items/8b4dfc6f441cf52a88b2)を参照してください。

- 起動
```bash
make run
```

- SwaggerUI で API 仕様と動作を確認することができます。
```bash
make open
```

- 停止
```bash
make stop
```

## .env について

環境変数は Docker-compose で動くことを前提にプログラム上でデフォルト値が設定されていますが、
`.env` ファイルを作成することで環境変数を上書きすることができます。

その場合は `sample.env` を参考にして、`.env` ファイルを作成してください。

なお、デフォルトでは以下の環境変数が設定されています。

```bash
DB_TYPE=mysql
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_ADDRESS=db:3306
MYSQL_DATABASE=g_gayaon
```