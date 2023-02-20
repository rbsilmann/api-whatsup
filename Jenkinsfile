pipeline {
  agent any
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
        sh 'docker build -t rbsilmann/api-whatsup:${env.BRANCH_NAME} .'
      }
    }
    stage('Test') {
      when {
        expression {
          return env.BRANCH_NAME.startsWith('main') || env.BRANCH_NAME.startsWith('FIS')
        }
      }
      steps {
        script {
          def containerId = sh(script: 'docker run -d -p 9098:9098 rbsilmann/api-whatsup:${env.BRANCH_NAME}', returnStdout: true).trim()
          try {
            def status = sh(script: "docker inspect -f '{{.State.Status}}' ${containerId}", returnStatus: true)
            if (status == 0) {
              echo 'Test passed!'
            } else {
              error 'Test failed!'
            }
          } finally {
            sh "docker stop ${containerId}"
          }
        }
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
          sh 'echo $PASSWORD | docker login -u $USERNAME --password-stdin'
          sh 'docker push rbsilmann/api-whatsup:${env.BRANCH_NAME}'
        }
      }
    }
  }
}