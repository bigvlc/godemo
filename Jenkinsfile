podTemplate(label: 'mypod',
    containers: [
        containerTemplate(name: 'docker', image: 'docker', ttyEnabled: true, command: 'cat'),
        containerTemplate(name: 'helm', image: 'lachlanevenson/k8s-helm', ttyEnabled: true, command: 'cat'),
        containerTemplate(name: 'kubectl', image: 'lachlanevenson/k8s-kubectl', ttyEnabled: true, command: 'cat'),
    ],
    volumes: [
        hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
    ]
)
{
    node('mypod') {
        stage('Publish Docker Image') {
            container('docker') {
                checkout scm
            }
        }

        stage('Deploy to Prod') {
            container('helm') {

            }
            container('kubectl') {

            }
        }
    }
}