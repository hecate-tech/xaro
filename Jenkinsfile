pipeline {
    agent any

    stages {
        stage('Checkout from Github') {
            steps {
                checkout scm
            }
        }

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
                
                
                script {
                    if (isUnix()) {
                        sh 'dep ensure'
                        sh 'go build -o Xaro .'
                        // sh 'mage'
                    } else {
                        bat 'dep ensure'
                        bat 'go build -o Xaro.exe .'
                        // bat 'mage'
                    }
                }
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