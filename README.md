# example-of-ec-site

## ローカル環境構築

```Shell
docker-compose up -d
```

## DBマイグレーション

dockerを立ち上げただけではテーブルが無いのでここを実施する必要がある

```Shell
# 確認
make db-migrate-dry-run

# 実施
make db-migrate
```

## ローカル環境のリンク

### Adminer

<http://localhost:8080>

### Swagger UI

<http://localhost:8081>

## GitHub Actions

* mainブランチにプッシュされると schema/tables.sql の内容で doc/schema/ 内のテーブル定義を自動更新する (差分があるときだけ)