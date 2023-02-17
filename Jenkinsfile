pipeline {
  agent any
  environment {
    BRANCH = "${env.GIT_BRANCH}"
  }
  options {
    buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '5', daysToKeepStr: '', numToKeepStr: '5')
  }
  stages {
    stage('Hello') {
      steps {
        sh '''
          echo $BRANCH
        '''
      }
    }
    stage('Build image') {
      when {
        branch "main"
        // branch "fix-*"
      }
      steps {
        script {
          app = docker.build("rbsilmann/api-whatsup:${BRANCH}")
        }
      }
    }
    stage('Push image') {
      steps {
        docker.withDockerRegistry([ credentialsId: "regcred", url: "" ]) {
          app.push()
        }
      }
    }
  }
}