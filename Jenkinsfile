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
                        def response = sh(script: 'curl -I -s -o /dev/null -w "%{http_code}" http://localhost:9098', returnStdout: true).trim()
                        if (response == '200') {
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
  }
}