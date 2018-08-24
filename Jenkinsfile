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
                
                mkdir -p $GOPATH/src/github.com/damienfamed75/engo-xaro
                ln -s $WORKSPACE $GOPATH/src/github.com/damienfamed75/engo-xaro
            }
        }
        stage('Build') {
            steps {    
                echo 'Building...'
                
                
                script {
                    if (isUnix()) {
                        sh 'dep init'
                        sh 'go build -o Xaro .'
                        // sh 'mage'
                    } else {
                        bat 'dep init'
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