pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'echo | sudo -S apt-get install build-essential'
                sh 'make cli-integration-test'
            }
        }
    }
}