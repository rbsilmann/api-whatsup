agent{
    label 'linux'
}

environment{
    DOCKERHUB_CREDENTIALS=credentials('dockerhub')
}

stages{
    stage('gitclone'){
        steps{
            git 'https://github.com/rbsilmann/api-whatsup.git'
        }
    }

    stage('build'){
        steps{
            sh 'docker build -t rbsilmann/api-whatsup:latest'
        }
    }

    stage('login'){
        steps{
            sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
        }
    }

    stage('push'){
        steps{
            sh 'docker push rbsilmann/api-whatsup:latest'
        }
    }
}

post{
    always {
        sh 'docker logout'
    }
}