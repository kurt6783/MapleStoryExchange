CREATE TABLE product (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    category TEXT DEFAULT ''
);

INSERT INTO product (name, category) VALUES ('火槍攻擊卷軸10%', '');