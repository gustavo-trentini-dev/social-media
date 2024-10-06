CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT current_timestamp
)

CREATE TABLE followers (
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,


follower_id INT NOT NULL,
FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE,

PRIMARY KEY(user_id, follower_id) ) 

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(300) NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE cascade,
    likes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT current_timestamp
)