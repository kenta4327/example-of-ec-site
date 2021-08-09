# example-of-ec-site

## ローカル環境構築

```Shell
docker-compose up -d
```

## DBマイグレーション

```Shell
# 確認
make db-migrate-dry-run

# 実施
make db-migrate
```

## リンク

### Adminer

<http://localhost:8080>

### Swagger UI

<http://localhost:8081>
