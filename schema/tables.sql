CREATE TABLE IF NOT EXISTS item (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '商品ID',
    name VARCHAR(255) NOT NULL COMMENT '商品名',
    price INT NOT NULL COMMENT '価格(税込み)',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時'
)ENGINE=InnoDB DEFAULT CHARACTER SET=utf8 COMMENT '商品テーブル';
