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

## テスト実施

```Shell
docker exec example-of-ec-site_api_1 go test ./app/test/usecase_test/
```

## mock_repository 作成

```Shell
docker exec example-of-ec-site_api_1 mockgen -source=app/domain/repository/item_repository.go -destination app/test/mock_repository/item_repository_mock.go
```

## ローカル環境のリンク

### Adminer

<http://localhost:8080>

### Swagger UI

<http://localhost:8081>

## GitHub Actions

* fmt-and-test
  * PR時に「go fmt」と「go test」を実施する
* update-table-defenition
  * mainブランチにプッシュされると schema/tables.sql の内容で docs/schema/ 内のテーブル定義を自動更新する (差分があるときだけ)

## go.mod の説明

* github.com/golang/mock
  * 「domain/repository」にあるrepositoryインターフェースに対するmockを作成するために使用
