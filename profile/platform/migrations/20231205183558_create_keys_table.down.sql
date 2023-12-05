IF EXISTS (SELECT 1 FROM sys.tables WHERE name = 'keys')
BEGIN
DROP TABLE keys;
END

