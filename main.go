package main

import (
	"context"
	"flag"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// deployment 만들 때 내용으로 들어가는 것들
type Dp struct {
	metaName     string
	countReplica *int32
	labelName    map[string]string
	template     Template
}

type Template struct {
	labelName map[string]string
	container Container
}

type Container struct {
	contName  string
	imageName string
	portName  string
	portNum   int32
}

// Container struct에 내용 넣어 객체 만드는 함수
func NewContainer(contName string, imageName string, portName string, portNum int32) *Container {
	container := Container{contName: contName, imageName: imageName, portName: portName, portNum: portNum}
	return &container
}

// template struct에 내용 넣어 객체 만드는 함수
func NewTemplate(labelName map[string]string, container Container) *Template {
	template := Template{labelName: labelName, container: container}
	return &template
}

// Dp struct에 내용 넣어 객체 만드는 함수
func NewDp(metaName string, countReplica *int32, labelName map[string]string, template Template) *Dp {
	dp := Dp{metaName: metaName, countReplica: countReplica, labelName: labelName, template: template}
	return &dp
}

// 내용 채워진 struct를 진짜 deployment로 만드는 함수
func makeDeployment(d *Dp) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.metaName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: d.countReplica,
			Selector: &metav1.LabelSelector{
				MatchLabels: d.labelName,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: d.template.labelName,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  d.template.container.contName,
							Image: d.template.container.imageName,
							Ports: []apiv1.ContainerPort{
								{
									Name:          d.template.container.portName,
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: d.template.container.portNum,
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}

// Homedir 써서 kubeconfig 잘 살아있는지 확인하는 함수
var kubeconfig *string

func checkKube() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

}

func int32Ptr(i int32) *int32 { return &i }

// 위에서 제작한 함수들을 순서대로 실행시키면 deployment가 생성됨
// + goroutine 써서 deployment 3개 동시에 띄우기
// waitgroup 등을 활용하여 동시실행 결과 알림 받고 싶음
func main() {
	// labelName map 만들기
	var ln map[string]string
	ln = make(map[string]string)
	ln["app"] = "webui"

	// replica 갯수 변환
	num := int32Ptr(3)

	// 일단 deployment 1개 만들어 보기
	checkKube()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	cont := NewContainer("nginx-container", "nginx:1.14", "http", 80)
	temp := NewTemplate(ln, *cont)
	dp := NewDp("deploy-nginx", num, ln, *temp)

	deployment := makeDeployment(dp)
	fmt.Println("deployment 만들기 시작")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("deployment %q를 만들었습니다", result.GetObjectMeta().GetName())
}
