-- Usuário conservador, transação normal
INSERT INTO users (id, name, profile, email, phone, created_at, status)
VALUES ('user-conservative', 'Ana Silva', 'conservative', 'ana.silva@email.com', '+55-11-98765-4321', '2023-01-15T10:30:00Z', 'active')
ON CONFLICT (id) DO NOTHING;

INSERT INTO transactions (user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp)
VALUES
('user-conservative', 100.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-01T10:00:00Z');

-- Cenário: Viagem impossível
INSERT INTO transactions (user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp)
VALUES
('user-conservative', 200.0, 'pix', 'credito', 'New York', 'US', 40.7128, -74.0060, 'ana-iphone-12', '2024-01-01T10:30:00Z');

-- Cenário: Valor anômalo
INSERT INTO transactions (user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp)
VALUES
('user-conservative', 1500.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-02T10:00:00Z');

-- Cenário: Dispositivo desconhecido
INSERT INTO transactions (user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp)
VALUES
('user-conservative', 100.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'unknown-device', '2024-01-03T10:00:00Z');

-- Cenário: Velocidade de transações
INSERT INTO transactions (user_id, amount, type, direction, city, country, latitude, longitude, device_id, timestamp)
VALUES
('user-conservative', 50.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-04T10:00:00Z'),
('user-conservative', 60.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-04T10:00:10Z'),
('user-conservative', 70.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-04T10:00:20Z'),
('user-conservative', 80.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-04T10:00:30Z'),
('user-conservative', 90.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-04T10:00:40Z'),
('user-conservative', 100.0, 'pix', 'credito', 'São Paulo', 'BR', -23.5505, -46.6333, 'ana-iphone-12', '2024-01-04T10:00:50Z');
