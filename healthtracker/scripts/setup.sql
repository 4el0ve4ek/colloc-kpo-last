CREATE TABLE "activity"(
    id SERIAL PRIMARY KEY,
    activity VARCHAR(50) NOT NULL,
    duration BIGINT NOT NULL,
    calories INT NOT NULL
);


CREATE TABLE "nutrition"(
    id SERIAL PRIMARY KEY,
    dish VARCHAR(50) NOT NULL,
    size INT NOT NULL,
    calories INT NOT NULL
);


CREATE TABLE "sleep"(
   id SERIAL PRIMARY KEY,
   duration BIGINT NOT NULL
);


