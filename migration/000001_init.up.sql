CREATE TABLE IF NOT EXISTS movie(
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	url TEXT NOT NULL,
	thumbnail TEXT NOT NULL,
	description TEXT NOT NULL
);

CREATE INDEX movie_title ON movie(title);

CREATE VIRTUAL TABLE IF NOT EXISTS movie_fts USING fts5(
	title,
	description,
	tokenize = 'porter'
);

CREATE TRIGGER movie_fts_delete AFTER DELETE ON movie
BEGIN
	INSERT INTO movie_fts(movie_fts, rowid, title, description) VALUES('delete', old.id, old.title, old.description);
END;

CREATE TRIGGER movie_fts_insert AFTER INSERT ON movie
BEGIN
	INSERT INTO movie_fts(rowid, title, description) VALUES(new.id, new.title, new.description);
END;

CREATE TRIGGER movie_fts_update AFTER UPDATE ON movie
BEGIN
	INSERT INTO movie_fts(movie_fts, rowid, title, description) VALUES('delete', old.id, old.title, old.description);
	INSERT INTO movie_fts(rowid, title, description) VALUES(new.id, new.title, new.description);
END;

CREATE TABLE IF NOT EXISTS movie_tag(
	id INTEGER PRIMARY KEY,
	movie_id INTEGER NOT NULL,
	name TEXT NOT NULL,

	FOREIGN KEY(movie_id) REFERENCES movie(id)
);
CREATE INDEX tag_name ON movie_tag(name);

CREATE TABLE IF NOT EXISTS movie_url(
	id INTEGER PRIMARY KEY,
	movie_id INTEGER NOT NULL,
	server TEXT NOT NULL,
	resolution TEXT NOT NULL,
	url TEXT NOT NULL,

	FOREIGN KEY(movie_id) REFERENCES movie(id)
);

CREATE INDEX movie_url_idx ON movie_url(server, resolution);
