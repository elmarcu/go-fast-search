CREATE TABLE IF NOT EXISTS user_activity (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    action VARCHAR(255) NOT NULL,
    metadata JSON,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- Example Inserts
INSERT INTO user_activity (user_id, action, metadata) VALUES
('user_001', 'login', '{"ip": "192.168.1.10"}'),
('user_002', 'view_page', '{"page": "home"}'),
('user_003', 'purchase', '{"item_id": "item_123", "amount": 49.99}');

-- Another batch
INSERT INTO user_activity (user_id, action, metadata) VALUES
('user_001', 'logout', '{}'),
('user_004', 'signup', '{"referral": "user_002"}'),
('user_002', 'add_to_cart', '{"item_id": "item_456"}');
