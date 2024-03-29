CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(50),
    Email VARCHAR(100) NOT NULL,
    Password VARCHAR(100) NOT NULL,
    Enable BOOLEAN NOT NULL,
    UNIQUE(Email)
);

CREATE TABLE IF NOT EXISTS products (
    ID SERIAL PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    Price INTEGER NOT NULL,
    Description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    ID SERIAL PRIMARY KEY,
    User_ID INTEGER NOT NULL,
    FOREIGN KEY(User_ID) REFERENCES users(ID),
);

CREATE TABLE IF NOT EXISTS order_items (
    Order_ID INTEGER NOT NULL,
    Product_ID INTEGER NOT NULL,
    Quantity INTEGER NOT NULL CHECK (Quantity > 0),
    PRIMARY KEY (Order_ID, Product_ID),
    FOREIGN KEY (Order_ID) REFERENCES orders(ID),
    FOREIGN KEY (Product_ID) REFERENCES products(ID)
);
