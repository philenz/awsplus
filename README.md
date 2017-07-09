#### AWS Plus

Adds functionality not found on the AWS console.

Initially the project will manage three extra pieces of information when someone adds an inbound rule to a security group.

Those pieces of information...
1. Who added the rule
2. When was the rule added
3. Description of the rule (e.g.: what is the ip address range for)

The extra information will be stored in a PostgreSQL database. Either locally or in RDS.

The program will run as a web server on port 9090 (as defined in config.json).

##### Contents

* _README.md_: This file

##### Setting Up

* Install postgresql https://fedoraproject.org/wiki/PostgreSQL
* go get github.com/lib/pq


