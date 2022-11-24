#!groovy
pipeline {
    agent any
    options {
        buildDiscarder(logRotator(numToKeepStr: '5', daysToKeepStr: '5'))
        timestamps()
    }

    environment {
        REGISTRY = "rbsilmann/api-whatsup"
        REGISTRY_CREDENTIALS=credentials('dockerhub')
    }

    stages {
        stage('git') {
            steps{
                git 'https://github.com/rbsilmann/api-whatsup.git'
            }
        }

        stage('build') {
            steps {
                script {
                    DOCKER_IMAGE = docker.build REGISTRY + ":$BUILD_NUMBER"
                }
            }
        }

        stage('push') {
            steps{
                script{
                    docker.withDockerRegistry('', REGISTRY_CREDENTIALS) {
                        DOCKER_IMAGE.push()
                    }
                }
            }
        }
    }
}

post {
    always {
        cleanWs()
    }
}