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
          docker build -t rbsilmann/api-whatsup:$BRANCH
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
    // stage('Build image') {
    //   steps {
    //     dockerImage = docker.build("rbsilmann/api-whatsup:${BRANCH}")
    //   }
    // }
    
    // stage('Push image') {
    //   steps {
    //     withDockerRegistry([ credentialsId: "regcred", url: "" ]) {
    //       dockerImage.push()
    //     }
    //   }
    // }
  }
}