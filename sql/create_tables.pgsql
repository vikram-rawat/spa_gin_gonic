
-- Database: hospital_tourism

-- DROP DATABASE hospital_tourism;

CREATE DATABASE hospital_tourism
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_India.1252'
    LC_CTYPE = 'English_India.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- add extensions

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION pgcrypto;

