CREATE TABLE user (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  mobile VARCHAR(20),
  latitude DECIMAL(10, 6),
  longitude DECIMAL(10, 6),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);



for creating user

INSERT INTO user ( name, mobile, latitude, longitude) VALUES
( 'John Doe', '1234567890', 37.7749, -122.4194),
('Jane Smith', '9876543210', 40.7128, -74.0060),
( 'Mike Johnson', '5551234567', 51.5074, -0.1278);