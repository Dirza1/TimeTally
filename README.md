# TimeTally

TimeTally is a local-first solution for small businesses that need a sophisticated yet straightforward way to track finances and project hours. Unlike cloud services, TimeTally keeps your data secure on your own machine using a fast, local PostgreSQL database—no fees, no subscriptions, just your records at your fingertips.

Built with Go, TimeTally offers a responsive command-line interface that makes recording transactions and time entries quick and efficient—no bulky software, just streamlined productivity.

## Why use TimeTally?

As a Beekeeper, I was tired of trying to manage my time and expenses using clunky time-tracking apps that required constant internet connections and charged monthly fees just to access my own data. I needed something simple, secure, and professional-looking for client reports and tax auditors. That's why I built TimeTally - a tool that puts you back in control of your time tracking.

Here's what makes TimeTally different:
- All your data is locally hosted on your PC. No cloud data leaks, no subscription fees for accessing your own data its all on your local computer.
- Professional look. If/When your customers ask for a justification of your time spent, showing a print out from TimeTally looks a lot more professional than an Excel sheet. Whether you need a overview by date ranges, categories or description TimeTally has you covered
- You don’t have to start up and shut down a program. TimeTally is a CLI tool that automatically logs you out when you are inactive.
- TimeTally will automatically backup your data as a database and a CSV file to ensure that your data is safe.
- Don’t compromise on security. TimeTally has a Login function with three access levels to ensure personnel only have the access they need.

## Quick Start

This part of the guild will explain how to get started.

### Install and initialize PostgreSQL
First we need to install a PostgreSQL database.
This is a widely used locally run database that is open source.
Documentation on this database can be found here: [Postgresql.org](https://www.postgresql.org/)
The easiest installation can be found here: [Install postgreSQL](https://www.w3schools.com/postgresql/postgresql_install.php)
Note the following: You need to ensure you have the following data saved:
-Your password
-The port you will use
-The username you select

Complete the instructions in the "install" and "Get Started"sections of the website.
You should continue until you have the screen which says: postgres=#
Then you have PostgreSQL installed.

To finish the installation in the SQL Shell type this:
"CREATE DATABASE your_database_name;" replace your_database_name with what you want your database called, and don’t forget the semicolon at the end.
Mine is called "Accounting".

At the end of this installation you should have noted down:
-Your password
-The port you will use
-The username you select
-Your database name

We need this information to set up the .env file in the next section.

### Set up your .env file
Now that we have PostgreSQL installed we need a way for our program to connect to it.
That is what the .env file is for.

WARNING:
The .env file will contain two important things:
1. Your database connection string that you can use to connect to the database, but other people can use it as well on your PC.
2. Your secret password used for resetting a database.
Ensure that unauthorized people cannot access your .env file!

The program should have been supplied with a .env.example file. You can remove the .example from the name.
Otherwise you need to make a new file ending in .env with the following content:
DB_URL="postgres://user:password@localhost:5432/mydb?sslmode=disable"
MANUAL_CONNECTION=" psql postgres://user:password@localhost:5432/mydb?sslmode=disable"
reset_password="YourSecretPassWord"

Change the following:
user -> your username
password -> your password
5432 -> the port you chose for the database
mydb -> your database name
YourSecretPassword -> A password you want to use to reset your databases

Save the .env file in the same folder you have the program itself.

### Initialize your database
Now that we have your database running, we can get it ready to work.
We need to initialize (also called migrate).
The program can do this for you.

#### Windows
Press windows + r, type CMD and press enter.
You are now in your command line. 

#### General
Navigate to the folder your program is saved in.
Then type: **TimeTally migrate**
Your should then see a message stating "migration completed"
This shows that the database is ready to go.
Otherwise the error message should give you information on how too solve it.
Double check the data in your .env file. Likely the issue is there

### Create your first administrator account
To get started you need your first account.
The program allows you to make a administrator account if none are available.
In your command line when you are in the filter the program is located type:
**TimeTally FirstAdmin -u username -p password**

The system then makes a username with the username and password.
You are now ready to start using the program!

## Using the program   
Using a CLI tool is verry easy.
We use the terminal or Command Line Interface the same we did for generating the first admin and database migration.
We use the terminal and navigate to the folder that contains the program.
(note to make it easy you can put the program in the folder that you start your terminal in)

When you are there you always work with the program like this:

**TimeTally [command] -[flag] [argument]**

So you always start with TimeTally, then the command you want to use.
You then put the required flags in place and the argument that is supposed to go with that flag.
You can find the flags with a command by typing this:

**TimeTally [command] -h**

That will give you a help section on the selected command.

## Commands
The folowing commands are avalible:
* AddAdmin
* AddUser
* deleteEntry
* DeleteUser
* FirstAdmin
* Login
* Logout
* migrate
* overview
* overviewByCategory
* overviewDates
* registerTime
* register Transaction
* reset
* updateTime
* updateTransaction
* UpdateUser
* UserOverview

You can also generate an overview of the available commands by typing:
**TimeTally -h**
It will show you a list of all available commands.

## Closure words
I hope you will find this program usefull.
Should you find any bugs, have any sugestions please let me know via: jasper.olthof@xs4all.nl

## Contributing
If you would like to Contribute to the program please create a fork andd make a pull request to the main brannch.
New functions should contain tests or have a clear explanation on why tests are not used.

### Clone the repo

```bash
git clone https://github.com/Dirza1/TimeTally
cd TimeTally
```

### Build the project

```bash
go build
```

Run the project and make the improvements.
Generate Tests and then make you pull request.

note: Your improvemennts may not be merged base on the creators ideas.
If you want to make sure your imporvements will be added, conntact the creator prior to starting.