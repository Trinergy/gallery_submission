CREATE TABLE images (
	id serial PRIMARY KEY,
	url VARCHAR ( 255 ) NOT NULL,
  flagged Boolean NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ
);

INSERT INTO images(url, flagged, created_at, updated_at) 
VALUES 
('https://picsum.photos/id/237/200/300', false, current_timestamp, current_timestamp),
('https://picsum.photos/id/1000/5626/3635', false, current_timestamp, current_timestamp),
('https://picsum.photos/id/1025/4951/3301', false, current_timestamp, current_timestamp),
('https://picsum.photos/id/102/4320/3240', false, current_timestamp, current_timestamp);