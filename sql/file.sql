-- name: create-table
CREATE TABLE IF NOT EXISTS file (
    id INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
	name NVARCHAR(250)  NULL,
    extension NVARCHAR(250)  NULL,
	path NVARCHAR(250)  NULL
);

-- name: create-many
INSERT INTO file (path, name, extension)
VALUES (?, ?, ?);

-- name: find-by-name-extension
SELECT id FROM file WHERE name = ? AND extension = ?;

-- name: find-by-name-extension-path
SELECT id FROM file WHERE name = ? AND extension = ? AND path = ?;

-- name: read-all-path
select f.path
from file as f
group by f.path
order by f.path ASC;

-- name: drop-table
DROP TABLE IF EXISTS file;
