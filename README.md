# APIを作りながら進むGo中級者への道 の blogAPIの写経

- [APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)


## 起動
```sh
$ cp .env.dist .env
$ vim .env
$ docker compose up -d
環境変数を設定して起動
$ DB_USER=docker DB_PASSWORD=docker DB_NAME=sampledb go run main.go
```

## テスト
```sh
$ docker compose up -d
$ go test ./...
```

