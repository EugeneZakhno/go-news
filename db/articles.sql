CREATE TABLE articles (
                          id SERIAL PRIMARY KEY,
                          title VARCHAR(255),
                          anons VARCHAR(255),
                          full_text text
);