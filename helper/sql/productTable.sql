CREATE TABLE products (    
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255),    
    product_description TEXT,   
    product_price DECIMAL(10, 2),  
    compressed_product_images TEXT,   
    product_images TEXT,
    created_at TIMESTAMP,    
    updated_at TIMESTAMP 
);


INSERT INTO products (product_name, product_description,  product_price, compressed_product_images) VALUES
( 'iPhone 12', 'The iPhone 12 is a powerful smartphone with an A14 Bionic chip and a Super Retina XDR display.', 999.99, 'compressed_image1.jpg,compressed_image2.jpg,compressed_image3.jpg'),
( 'Samsung Galaxy S21', 'The Samsung Galaxy S21 features a stunning Dynamic AMOLED display and a versatile camera system.',899.99, 'compressed_image4.jpg,compressed_image5.jpg,compressed_image6.jpg'),
( 'Google Pixel 6', 'The Google Pixel 6 offers a pure Android experience and exceptional camera capabilities.', 799.99, 'compressed_image7.jpg,compressed_image8.jpg,compressed_image9.jpg');
