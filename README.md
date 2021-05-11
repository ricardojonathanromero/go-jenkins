# go-jenkins
This project will be compiled and run using jenkins pipeline.

## Testing
In this section you will show how run the test suite. We have two types of test, unit and integration test that are developed by the developers, and does not have relation with the tests suites that the QA applies.

### Integration Tests
For execute the integration testing events, you need to execute via IntelliJ because the plugin only exist in this IDE.
```
make init-integration-test
```
After the integration test finish, stop the server manually.

However, if you do not have the IDE, you can run the integrations tests executing the next command.
```
make cli-integration-test
```

### Unit Tests
For run the unit tests, execute the command below.
```
make unit-tests
```

It will start the server, and then you will need to execute manually the file `integration/tests.http` 

## Documentation
Here, you will look the swagger documentation related to all services exposed.
The route is `documentation/swagger.yaml`

## Build project
For build the container, you need to execute the next command: `make build-ci`

## Deploy to Kubernetes
Finally, If you would like to deploy the service to kubernetes, you can make executing:
```
make deploy-ci
```

## CI-CD (Jenkins)
As alternative, the repo is configured to be build via jenkins using the `jenkinsfile`

## Notes
Additionally, the repo is compiled using the Dockerfile, and you need to have access to the repository for build the container, or replicate the configuration in your repo.
