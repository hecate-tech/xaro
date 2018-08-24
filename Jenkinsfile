node {
    def root = tool name: 'Go1.10.1', type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/damienfamed75/engo-xaro") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"
            
            stage 'Checkout'
        
            git url: 'https://github.com/Damienfamed75/engo-xaro.git'
        
            stage 'preTest'
            sh 'go version'
            sh 'go get -u github.com/golang/dep/...'
            sh 'dep init'
            
            stage 'Test'
            sh 'go vet'
            sh 'go test -cover'
            
            stage 'Build'
            sh 'go build .'
            
            stage 'Deploy'
            // Do nothing.
        }
    }
}

// pipeline {
//     agent any

//     String applicationName = "engo-xaro"
//     // a basic build number so that when we build and push to Artifactory we will not overwrite our previous builds
//     String buildNumber = "0.1.${env.BUILD_NUMBER}"
//     // Path we will mount the project to for the Docker container
//     String goPath = "/go/src/github.com/damienfamed75/${applicationName}"

//     stages {
//         stage('Checkout from Github') {
//             steps {
//                 checkout scm
//             }
//         }

//         stage('Preparation') {
//             steps {
//                 echo 'Preparing...'
//                 // Get repository
//                 // git 'https://github.com/Damienfamed75/engo-xaro.git'
                
//                 // script {
//                 //     if (isUnix()) {
//                 //         sh 'go version'
//                 //         sh 'go get github.com/magefile/file'
//                 //     } else {
//                 //         bat 'go version'
//                 //         bat 'go get github.com/magefile/file'
//                 //     }
//                 // }
//             }
//         }
//         stage('Build') {
//             steps {    
//                 echo 'Building...'
                
                
//                 script {
//                     if (isUnix()) {
//                         sh 'dep ensure'
//                         sh 'go build -o Xaro .'
//                         // sh 'mage'
//                     } else {
//                         bat 'dep ensure'
//                         bat 'go build -o Xaro.exe .'
//                         // bat 'mage'
//                     }
//                 }
//             }
//         }
//         stage('Test') {
//             steps {
//                 echo 'Testing...'
//             }
//         }
//         stage('Deploy') {
//             steps {
//                 echo 'Deploying...'
//             }
//         }
//     }
// }