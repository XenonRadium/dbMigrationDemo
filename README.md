# dbMigrationDemo

This project serves as a demonstration of a prototype local database migration in Golang in combination with the golang-migrate library and the Cobra library
to create new command line commands.

An executable file is created to be ran in the command line accompanied with the commands to be run.
## ./migrationCli.exe migrate <arguments>

Current version of the project supports all up migrations and down migrations via the
## migrate up / migrate down
