package ready

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
	"sync"
)

var (
	kubeconfig *string
)

// CheckKube로 config 설정 한번 체크하고, 그 결과를 deployment 만들때 적용시키기
func CheckKube() *string {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	return kubeconfig
}

func ShootDeploy(d *appsv1.Deployment, kubeconfig *string, wg *sync.WaitGroup) {
	fmt.Println("deployment 만들기 시작")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("deployment 만들기 시작")
	result, err := deploymentsClient.Create(context.TODO(), d, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("deployment %q를 만들었습니다\n", result.GetObjectMeta().GetName())
	wg.Done()
}
