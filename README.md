#### AWS Plus

##### Introduction 

Adds functionality not found on the AWS console.

Initially the project will manage three extra pieces of information when someone adds an inbound rule to a security group.

Those pieces of information...
1. Who added the rule
2. When was the rule added
3. Description of the rule (e.g.: what is the ip address range for)

##### To Do

* Write a Lambda function to write an SQS message when a security group ingress is changed.
* Write DB init to create/populate database if not already done.
* Write code to update database when read message from SQS.

##### Design

* Use PostgreSQL or MongoDB? Go with PostgreSQL initially.
* Use RDS or local? Local.
* Run as Docker container? Any advantage?
* Define web server port etc in config.json.
* In DB init function, populate security group table if no rows.
* When someone adds an ingress security group rule lambda function will:
* (1) Add a row to the database table;
* (2) Put a message on in SQS.
* App will check SQS on startup and periodically after that.
* Write lambda code in Python and add to this project.
* Use Vue.js for front-end? Or just jQuery / jQueryUI.

##### Contents

* _README.md_: This file
* _config/config.go_: Read config settings
* _datastore/datastore.go_: Open db connection; db operations
* _logging/logging.go_: Log stuff
* _frontend_: Static files (css, js, etc)
* _routes_: Routes available to the app
* _templates_: Templates used by each route
* _main.go_: Entrypoint, starts up web server
* _runtime_: Runtime config and logs

##### Setting Up

* Install postgresql https://fedoraproject.org/wiki/PostgreSQL
* go get github.com/lib/pq


