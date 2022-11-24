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
        USER_CREDENTIALS=credentials('jenkinsuser')
    }

    stages {
        stage('git') {
            steps{
                git branch: 'main', url: 'https://github.com/rbsilmann/api-whatsup'
            }
        }

        stage('build') {
            steps {
                script {
                    docker.withDockerServer([uri: 'tcp://172.17.0.1:4243', credentialsId: USER_CREDENTIALS]) {
                        DOCKER_IMAGE = docker.build REGISTRY + ":$BUILD_NUMBER"
                    }
                }
            }
        }

        stage('push') {
            steps{
                script{
                    docker.withDockerServer([uri: 'tcp://172.17.0.1:4243', credentialsId: USER_CREDENTIALS]) {
                        docker.withDockerRegistry('', REGISTRY_CREDENTIALS) {
                            DOCKER_IMAGE.push()
                        }
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