Skill_Manager_App is an application to manage their employee skills. The application has following modules:

Set skill : This module allocates skill to an employee.
Get employee/s : This module fetches the employee/s based on skill.
Add employee/s : This module adds employee/s to existing skill.
Delete employee : This module deletes employee from a skill.
Delete skill : This module deletes skill which is no longer required by the organization.

utility	package contains utilities for db connection, reading properties file, custom error structs and logging	(separation of concerns)

model	package contains entity and data transfer struct instances	

validator package	contains functionality for validating model struct(s)

repository apckage contains functionalities for interacting with db	

service	package contains business logic of the application	

controller package contains Gin Handler functions for the REST Endpoints	

middleware package contains custom middleware for handling errors for REST Endpoints	

main.go file contains all required routes defined.
