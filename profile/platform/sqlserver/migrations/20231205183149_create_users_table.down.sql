IF EXISTS (SELECT 1 FROM sys.tables WHERE name = 'users')
BEGIN
DROP TABLE users;
END