CREATE TABLE IF NOT EXISTS orders (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  buyer INT NOT NULL,
  seller INT NOT NULL,
  delivery_source_address VARCHAR(255) NOT NULL,
  delivery_destination_address VARCHAR(255) NOT NULL,
  items VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  price VARCHAR(255) NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  status INT NOT NULL DEFAULT 0,
  FOREIGN KEY (buyer) REFERENCES buyer(id),
  FOREIGN KEY (seller) REFERENCES seller(id)
);
