/*
receivers
*/

CREATE TABLE IF NOT EXISTS item (
    id bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name varchar(20) NOT NULL,
    labels JSON NULL,
    annotations JSON NULL
);
