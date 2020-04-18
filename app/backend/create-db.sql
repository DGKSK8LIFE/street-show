CREATE DATABASE Street-Show;
USE Street-Show;

CREATE TABLE User (
	username text,
	name text,
	email text,
	id int NOT NULL AUTO_INCREMENT,
	PRIMARY KEY (id)
);

CREATE TABLE Busker (
	username text,
	name text,
	email text,
	id int NOT NULL AUTO_INCREMENT,
	PRIMARY KEY (id)
);

CREATE TABLE Performance (
	username text,
	coordinate POINT NOT NULL,
	SPATIAL INDEX `SPATIAL` (coordinate)
) ENGINE=InnoDB;


