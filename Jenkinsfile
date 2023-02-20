pipeline {
  agent any
  environment {
    BRANCH = "${env.GIT_BRANCH}"
    SLACK_TOKEN = credentials('slackcred')
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
    stage('Notify Slack') {
      when {
        expression {
          return env.BRANCH_NAME.startsWith('main') || env.BRANCH_NAME.startsWith('FIS')
        }
      }
      steps {
        script {
          def color = ""
          def status = currentBuild.currentResult
          if (status == "SUCCESS") {
            color = "#36a64f"
          } else if (status == "FAILURE") {
            color = "#FF0000"
          } else {
            color = "#FFFF00"
          }
          def message = """
            *Job:* ${env.JOB_NAME}
            *Build:* <${env.BUILD_URL}|${env.BUILD_NUMBER}>
            *Status:* ${status}
            *Branch:* ${env.BRANCH_NAME}
            *Commit:* <${env.CHANGE_URL}|${env.CHANGE_ID}>
            *Author:* ${env.CHANGE_AUTHOR}
            *Message:* ${env.CHANGE_TITLE}
            *Duration:* ${currentBuild.durationString}
          """
          slackSend (color: color, message: message, tokenCredentialId: 'slackcred', channel: '#jenkins-qa', baseUrl: 'https://slack.com/api/')
        }
      }
    }
  }
}
