pipeline {
    agent any

    // 阶段
    stages{
        // 步骤
        stage('拉取代码'){
            git 'git@github.com:rinuandengfeng/jenkins-test.git'
        }

        stage('构建'){
            make
        }

        stage('镜像'){
            echo "构建docker镜像"
            echo "推送到仓库"
        }
    }

}