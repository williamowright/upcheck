CREATE TABLE checks (
  id SERIAL PRIMARY KEY,
  url TEXT,
  status_code INT,
  response_time INT,
  is_up BOOLEAN
)