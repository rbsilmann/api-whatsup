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
          java -version
          echo $BRANCH
        '''
      }
    }
    stage('cat README') {
      when {
        branch "fix-*"
      }
      steps {
        sh '''
          cat README.md
        '''
      }
    }
    stage('Build image') {
      dockerImage = docker.build("rbsilmann/api-whatsup:${BRANCH}")
    }
    
    stage('Push image') {
      withDockerRegistry([ credentialsId: "regcred", url: "" ]) {
        dockerImage.push()
      }
    }
  }
}