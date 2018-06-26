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
                dir ('/home/jenkins/go/src/github.com/jenkins-x/terraform-provider-jx') {
                    checkout scm
                    container('go') {
                        sh "make fmt testacc"
                    }
                }
            }
        }

        stage('Build and Push Release') {
            when {
                branch 'master'
            }
            steps {
                dir ('/home/jenkins/go/src/github.com/jenkins-x/terraform-provider-jx') {
                    checkout scm
                    container('go') {
                        sh "make"
                    }
                }
            }
        }
    }
}