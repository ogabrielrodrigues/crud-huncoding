CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

USE crud_huncoding;

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  age INTEGER NOT NULL,
  password TEXT NOT NULL
);