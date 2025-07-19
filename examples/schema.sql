CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE organization (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE member (
    user_id INT NOT NULL,
    organization_id INT NOT NULL,
    role VARCHAR(50),
    PRIMARY KEY (user_id, organization_id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (organization_id) REFERENCES organization(id)
);