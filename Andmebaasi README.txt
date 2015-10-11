Andmebaasiga connectimine:
heroku pg:psql --app piletimyyk

Väljumine:
ctrl + c
----------------------
Tabelid:
	Andmed:
	CREATE TABLE andmed(eesnimi varchar(50), perekonnanimi varchar(50), email varchar(50), telefon integer);
	Lisamine andmetesse: 
	INSERT INTO andmed VALUES ('Ranal', 'Saron', 'ranalsaron@gmail.com', 5120628);
	INSERT INTO andmed VALUES ('Art-Norman', 'Reins', 'artreins@gmail.com', 58163989);
	