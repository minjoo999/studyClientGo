package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"studyClientGo/ready"
	"sync"
)

// labelName부터 deployment까지 한번에 만들어주는 함수 만들기
// 그 함수를 통해 여러 개의 deployment 객체를 쉽게 만들고, goroutine을 간단히 돌릴 수 있게 하기
func oneStopDeploy(contName string, imageName string, portName string, portNum int32,
	labelKey string, labelValue string, metaName string, repNum int32) *appsv1.Deployment {
	// replica 갯수, labelName map
	ln := ready.MakeLableName(labelKey, labelValue)
	num := ready.Int32Ptr(repNum)

	// 컨테이너, 템플릿, deployment 요소채우기
	cont := ready.NewContainer(contName, imageName, portName, portNum)
	temp := ready.NewTemplate(ln, *cont)
	dpm := ready.NewDp(metaName, num, ln, *temp)

	// 완성된 deployment 객체
	dpl := ready.MakeDeployment(dpm)
	return dpl
}

// 위에서 제작한 함수들을 순서대로 실행시키면 deployment가 생성됨
// goroutine, waitgroup 써서 deployment 3개 동시에 띄우기
func main() {
	dpl1 := oneStopDeploy("nginx-container", "nginx:1.14", "http", 80,
		"app", "webui", "deploy-nginx", 3)
	dpl2 := oneStopDeploy("centos-container", "centos:7", "http", 1000,
		"app", "os", "deploy-centos", 2)
	dpl3 := oneStopDeploy("wordpress", "wordpress:3.1", "http", 8080,
		"app", "wordpress", "deploy-wordpress", 3)

	check := ready.CheckKube()

	var wg sync.WaitGroup
	wg.Add(3)

	// ShootDeploy에 wg 넣어야 함
	go ready.ShootDeploy(dpl1, check, &wg)
	go ready.ShootDeploy(dpl2, check, &wg)
	go ready.ShootDeploy(dpl3, check, &wg)
	wg.Wait()

	fmt.Println("주어진 Deployment를 모두 만들었습니다")

}
