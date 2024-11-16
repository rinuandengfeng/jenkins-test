pipeline {
    agent any

    // 配置go的版本
//     tools{
//         go 'go1.21.0'
//     }

    // 配置go的环境变量
    environment{
        GOPROXY= 'https://goproxy.cn,direct'
        PATH = "${env.PATH}:/usr/local/go/bin"
    }
    // 阶段
    stages{
        // 阶段
        stage('拉取代码'){
            // 步骤
            steps{
                sh '''
                go version
                '''
            }
        }

        stage('构建'){
            steps{
                sh '''
                    make
                   '''
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