pipeline {
    agent {
        label "jenkins-go"
    }
    stages {
        stage('CI Build and push snapshot') {
            when {
                branch 'PR-*'
            }
            steps {
                container('jx-base') {
                    sh "make testacc"
                }
            }
        }

        stage('Build and Push Release') {
            when {
                branch 'master'
            }
            steps {
                container('jx-base') {
                    sh "make"
                }
            }
        }
    }
}