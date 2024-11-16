pipeline {
    agent any

    // 阶段
    stages{
        // 阶段
        steps('拉取代码'){
            // 步骤
            steps{
                sh git 'git@github.com:rinuandengfeng/jenkins-test.git'
            }
        }

        stage('构建'){
            steps{
                sh make
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