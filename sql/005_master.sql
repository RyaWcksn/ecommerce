CREATE TABLE IF NOT EXISTS buyer (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  alamat_pengiriman VARCHAR(255) NOT NULL
);


INSERT INTO buyer (email, name, password, alamat_pengiriman)
VALUES ('user@mail.com', 'John Doe', '$2a$10$xCniCdjZENytpQ7/NTNQduSJaZ6pl3bFMbd7bfF4OLwwbCyhGH8rC', '123 Main St');


CREATE TABLE IF NOT EXISTS seller (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  alamat_pickup VARCHAR(255) NOT NULL
);


INSERT INTO seller (email, name, password, alamat_pickup)
VALUES ('seller@mail.com', 'John Doe', '$2a$10$xCniCdjZENytpQ7/NTNQduSJaZ6pl3bFMbd7bfF4OLwwbCyhGH8rC', '123 Main St');

CREATE TABLE IF NOT EXISTS product (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  seller INT NOT NULL,
  FOREIGN KEY (seller) REFERENCES seller(id)
);

CREATE TABLE IF NOT EXISTS orders (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  buyer INT NOT NULL,
  seller INT NOT NULL,
  delivery_source_address VARCHAR(255) NOT NULL,
  delivery_destination_address VARCHAR(255) NOT NULL,
  items VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  status INT NOT NULL DEFAULT 0,
  FOREIGN KEY (buyer) REFERENCES buyer(id),
  FOREIGN KEY (seller) REFERENCES seller(id)
);
