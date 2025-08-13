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

### Instal and initialise PostgreSQL
First we need to instal a PostgreSQL database.
This is a widly used localy run database that is open source.
Documentation on this databse can be found here: [link](https://www.postgresql.org/)

### Set up your .env file

### Create your first administrator account

