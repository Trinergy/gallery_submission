CREATE TABLE images (
	id serial PRIMARY KEY,
	url VARCHAR ( 255 ) NOT NULL,
  flagged Boolean NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ
);

INSERT INTO images(url, flagged, created_at, updated_at) 
VALUES 
('https://picsum.photos/200/100', false, current_timestamp, current_timestamp),
('https://source.unsplash.com/random', false, current_timestamp, current_timestamp),
('https://loremflickr.com/320/240', false, current_timestamp, current_timestamp),
('https://picsum.photos/200/300', false, current_timestamp, current_timestamp);