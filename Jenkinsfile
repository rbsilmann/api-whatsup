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
    stage('Test') {
      when {
        expression {
          return env.BRANCH_NAME.startsWith('main') || env.BRANCH_NAME.startsWith('FIS')
        }
      }
      steps {
        script {
          def containerId = sh(script: 'docker run -d -p 9098:9098 rbsilmann/api-whatsup:$BRANCH', returnStdout: true).trim()
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
          sh 'docker push rbsilmann/api-whatsup:$BRANCH'
        }
      }
    }
    post {
      always {
        script {
          def slackChannel = '#alertmanager'
          def slackToken = 'xoxe.xoxp-1-Mi0yLTQzOTQ5MTIzNTAzNzAtNDM5NDgwODM4NjExNS00ODMwMDA5OTUxMzYzLTQ4MjcxODk1NTc4MTMtMzQyMTYzMGJiZjA2MDg3MWFkOTFkMWNhNGRiZGVmMjc1MDY4YmYwZTIxYWUyMzhkNzEwNGU2MTI1ZWNkZmE0NA'
          def slackUrl = "https://slack.com/api/chat.postMessage?token=${slackToken}&channel=${slackChannel}&text="
          def emailTo = 'operacoes@vrsoft.com.br'
          def emailSubject = 'Resultado do pipeline'
          def emailBody = 'O pipeline ${env.JOB_NAME} para a branch ${env.BRANCH_NAME} foi concluído com o status ${currentBuild.result}'

          if (env.SLACK_NOTIFY == 'true') {
            sh "curl -X POST '${slackUrl}Pipeline completed for *${env.JOB_NAME}* branch *${env.BRANCH_NAME}* with status *${currentBuild.result}*'"
          }

          if (env.EMAIL_NOTIFY == 'true') {
            emailext (
              to: emailTo,
              subject: emailSubject,
              body: emailBody
            )
          }
        }
      }
    }
  }
}