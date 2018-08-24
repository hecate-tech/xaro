pipeline {
    agent any

    stages {
        stage('Preparation') {
            steps {
                echo 'Preparing...'
                // Get repository
                // git 'https://github.com/Damienfamed75/engo-xaro.git'
                
                // script {
                //     if (isUnix()) {
                //         sh 'go version'
                //         sh 'go get github.com/magefile/file'
                //     } else {
                //         bat 'go version'
                //         bat 'go get github.com/magefile/file'
                //     }
                // }
            }
        }
        stage('Build') {
            steps {    
                echo 'Building...'
                
                // script {
                //     if (isUnix()) {
                //         // sh """cd $GOPATH/src/github.com/engo-xaro/"""
                //         sh 'mage'
                //     } else {
                //         // bat """cd $GOPATH/src/github.com/engo-xaro/"""
                //         bat 'mage'
                //     }
                // }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
            }
        }
    }
}