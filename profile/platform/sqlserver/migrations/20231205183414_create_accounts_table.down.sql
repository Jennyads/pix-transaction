IF EXISTS (SELECT 1 FROM sys.tables WHERE name = 'accounts')
BEGIN
DROP TABLE accounts;
END
