pipeline {
    agent any
    tools {
        go 'Go 1.15.6'
    }
    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
                sh 'go test -v'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'go build'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}