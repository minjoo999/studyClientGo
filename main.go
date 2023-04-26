package main

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// deployment 만들 때 내용으로 들어가는 것들
type Dp struct {
	metaName     string
	countReplica int32
	labelName    map[string]string
	template     Template
}

type Template struct {
	labelName     map[string]string
	container Container
}

type Container struct {
	contName string
	imageName string
	portName string
	portNum int
}

// Container struct에 내용 넣어 객체 만드는 함수
func NewContainer(contName string, imageName string, portName string, portNum int) *Container {
	container := Container{contName: contName, imageName: imageName, portName: portName, portNum: portNum}
	return &container
}

// template struct에 내용 넣어 객체 만드는 함수
func NewTemplate(labelName map[string]string, container Container) *Template {
	template := Template{labelName: labelName, container: Container{}}
	return &template
}

// Dp struct에 내용 넣어 객체 만드는 함수
func NewDp(metaName string, countReplica int, labelName map[string]string, template Template) *Dp {
	dp := Dp{metaName: metaName, countReplica: countReplica, labelName: labelName, template: Template{}}
	return &dp
}

// 내용 채워진 struct를 진짜 deployment로 만드는 함수
func (d Dp) makeDeployment() {
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

				}
			}
		}
	}
}

// Homedir 써서 kubeconfig 잘 살아있는지 확인하는 함수

// 위에서 제작한 함수들을 순서대로 실행시키면 deployment가 생성됨
// + goroutine 써서 deployment 3개 동시에 띄우기
// waitgroup 등을 활용하여 동시실행 결과 알림 받고 싶음
func main() {
}
