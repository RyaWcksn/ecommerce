CREATE TABLE IF NOT EXISTS seller (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  alamat_pickup VARCHAR(255) NOT NULL
);


INSERT INTO seller (email, name, password, alamat_pickup)
VALUES ('user@mail.com', 'John Doe', '$2a$10$xCniCdjZENytpQ7/NTNQduSJaZ6pl3bFMbd7bfF4OLwwbCyhGH8rC', '123 Main St');
