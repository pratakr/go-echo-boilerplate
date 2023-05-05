pipeline {
  agent any
  stages {
    stage('Initial') {
      steps {
        echo 'Init docker dependency'
        sh 'docker image prune -f'
      }
    }

    stage('DockerBuild Prod') {
      when {
        branch 'master'
      }
      steps {
        echo 'Make docker image'
        sh 'docker build -t localhost:32000/kancha-api .'
      }
    }

    stage('DeployECR Prod') {
      when {
        branch 'master'
      }
      steps {
        echo 'Push docker image to ECR'
        sh 'docker push localhost:32000/kancha-api'
      }
    }

    stage('Publish Prod') {
      when {
        branch 'master'
      }
      steps {
        echo 'Publish to server'
        sh '/var/jenkins_home/bin/kubectl -n default rollout restart deployments/kancha-api'
      }
    }

  }
  triggers {
    pollSCM('* * * * *')
  }
}