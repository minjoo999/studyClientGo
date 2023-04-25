module studyClientGo

go 1.16

replace (
	k8s.io/api => k8s.io/api v0.0.0-20230424214213-e93fc0e43e75
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20230424213653-0db4b4996746
)

require (
	k8s.io/client-go v0.27.1
	k8s.io/api v0.27.1
	k8s.io/apimachinery v0.27.1
)
