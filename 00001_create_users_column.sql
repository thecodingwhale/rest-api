-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) DEFAULT '',
  age int(11) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
