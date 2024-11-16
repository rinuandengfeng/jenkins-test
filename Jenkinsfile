pipeline {
    agent any

    // 阶段
    stages{
        // 阶段
        stage('拉取代码'){
            // 步骤
            steps{
                sh 'go version'
            }
        }

        stage('构建'){
            steps{
                sh """
                    make mod
                    make build
                   """
            }

        }

        stage('镜像'){
            steps{
                echo "构建docker镜像"
                echo "推送到仓库"
            }
        }
    }

}