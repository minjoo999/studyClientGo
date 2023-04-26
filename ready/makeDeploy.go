package ready

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 컨테이너, 템플릿, dp를 변수 선언하고 만든 뒤
// 변수에다가 함수 먹여서 객체 만들어준 다음에
// 그 객체로 realDeploy 함수 만들어주기

func Int32Ptr(i int32) *int32 { return &i }

func MakeDeployment(d *dp) *appsv1.Deployment {
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
