/*
* @Author: Mauhoi WU
* @Date:   2019-10-23 16:50:31
* @Last Modified by:   Mauhoi WU
* @Last Modified time: 2019-11-18 20:33:10
*/

CREATE SCHEMA ENIGM;

DROP TABLE ENIGM.TOKEN;
DROP TABLE ENIGM.USER;

CREATE TABLE IF NOT EXISTS ENIGM.USER (
	id SERIAL,
	username VARCHAR(20) UNIQUE NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(60) NOT NULL,
	country VARCHAR(3),
	full_name VARCHAR(50),
	telephone VARCHAR(15),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_at TIMESTAMP,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ENIGM.TOKEN (
	id SERIAL,
	user_id integer REFERENCES ENIGM.USER(id) ON DELETE RESTRICT,
	token VARCHAR(50) NOT NULL,
	expired_at TIMESTAMP,
	PRIMARY KEY (id, user_id)
);
