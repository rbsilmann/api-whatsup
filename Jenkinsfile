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
    stage('Initialize'){
      steps {
        def dockerHome = tool 'myDocker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
      }
    }
    stage('Build') {
      when {
        // branch "fix-*"
        branch "main"
      }
      steps {
        script {
          app = docker.build("rbsilmann/api-whatsup:${BRANCH}")
        }
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