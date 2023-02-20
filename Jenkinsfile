pipeline {
  agent any
  environment {
    BRANCH = "${env.GIT_BRANCH}"
  }
  options {
    buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '5', daysToKeepStr: '', numToKeepStr: '5')
  }
  stages {
        stage('Build Docker Image') {
            when {
                expression {
                    return env.BRANCH_NAME.startsWith('main') || env.BRANCH_NAME.startsWith('FIS')
                }
            }
            steps {
                sh 'docker build -t rbsilmann/api-whatsup:$BRANCH .'
            }
        }
        stage('Push Docker Image') {
            when {
                expression {
                    return env.BRANCH_NAME.startsWith('main') || env.BRANCH_NAME.startsWith('FIS')
                }
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'regcred', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh 'docker login -u $USERNAME -p $PASSWORD rbsilmann'
                    sh 'docker push rbsilmann/api-whatsup:$BRANCH'
                }
            }
        }
  }
}