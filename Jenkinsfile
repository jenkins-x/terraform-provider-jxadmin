pipeline {
    agent any
    stages {
        stage('CI Build and push snapshot') {
            when {
                branch 'PR-*'
            }
            steps {
                dir ('/home/jenkins/go/src/github.com/jenkins-x/terraform-provider-jx') {
                    checkout scm

                    sh "make clean fmt build"

                    echo "Now running tests..."
                    sh "make testacc"
                }
            }
        }

        stage('Build and Push Release') {
            when {
                branch 'master'
            }
            steps {
                dir ('/home/jenkins/go/src/github.com/jenkins-x/terraform-provider-jx') {
                    git "https://github.com/jenkins-x/terraform-provider-jx"

                    sh "make"
                }
            }
        }
    }
}