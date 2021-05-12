pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'sudo apt-get install build-essential'
                sh 'make cli-integration-test'
            }
        }
    }
}