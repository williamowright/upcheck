CREATE TABLE checks (
  id SERIAL,
  url TEXT,
  status_code INT,
  response_time INT,
  checked_at TIMESTAMP NOT NULL, 
  is_up BOOLEAN,
  PRIMARY KEY (id, checked_at)
)