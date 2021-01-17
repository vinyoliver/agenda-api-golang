01. An endpoint for pushing new contacts - OK
02. An endpoint for editing contact information - OK
03. An endpoint for deleting a contact - OK
04. An endpoint for searching a contact by it's id - OK
05. An endpoint for searching contacts by a part of their name - OK
06. An endpoint that lists all contacts - OK

07. The http service should be configurable through flags on startup (timeouts and port to use)
08. Log messages should be written for each endpoint hit, with the response status code and the time it took to fulfill the request
09. If an error occurs, a log message should be written with the response status code and a helpful error message, to help an engineer troubleshoot the issue
10. Service and host metrics should be collected. I suggest using Prometheus (https://prometheus.io/docs/guides/go-application/)
11. The application should have a reasonable test coverage, preferably above 70%
12. The application should have end-to-end tests (this is a good way to try out the http client)
13. The application should contain a buildable Dockerfile (https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e) -- care on this, as using plainly the scratch image might hinder you from making https requests. Not that this will impact our example, but something to always take care into the future
14. It would be nice for the application to have some type of storage to persist the data. I'll leave this open, feel free to pick any type of storage you want - OK