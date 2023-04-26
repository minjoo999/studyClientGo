package main

import (
	"studyClientGo/ready"
)

// 변수 정리: ready 패키지에서 만들어 놓은 struct와 함수를 갖고, 단계적으로 변수를 만든다
var (
	// replica 갯수, labelName map
	ln  = ready.MakeLableName("app", "webui")
	num = ready.Int32Ptr(3)

	// 컨테이너, 템플릿, deployment요소
	cont = ready.NewContainer("nginx-container", "nginx:1.14", "http", 80)
	temp = ready.NewTemplate(ln, *cont)
	dpm  = ready.NewDp("deploy-nginx", num, ln, *temp)

	// 완성된 deployment 객체
	dpl = ready.MakeDeployment(dpm)
)

// 위에서 제작한 함수들을 순서대로 실행시키면 deployment가 생성됨
// + goroutine 써서 deployment 3개 동시에 띄우기
// waitgroup 등을 활용하여 동시실행 결과 알림 받고 싶음
func main() {

	ready.ShootDeploy(dpl)

}
