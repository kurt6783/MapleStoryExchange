CREATE TABLE item (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INT NOT NULL,
    owner_id INT NOT NULL,
    status BOOLEAN DEFAULT TRUE,
    price INT NOT NULL CHECK (price >= 0),
    memo TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE RESTRICT,
    CONSTRAINT fk_owner FOREIGN KEY (owner_id) REFERENCES user(id) ON DELETE RESTRICT
);
CREATE INDEX idx_item_product_id ON item(product_id);
CREATE INDEX idx_item_owner_id ON item(owner_id);

INSERT INTO item (product_id, owner_id, status, price, memo) VALUES (1, 1, true, 9000000, '帥又大碗');