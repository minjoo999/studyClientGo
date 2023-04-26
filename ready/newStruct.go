package ready

type dp struct {
	metaName     string
	countReplica *int32
	labelName    map[string]string
	template     template
}

type template struct {
	labelName map[string]string
	container container
}

type container struct {
	contName  string
	imageName string
	portName  string
	portNum   int32
}

// Container struct에 내용 넣어 객체 만드는 함수
func NewContainer(contName string, imageName string, portName string, portNum int32) *container {
	container := container{contName: contName, imageName: imageName, portName: portName, portNum: portNum}
	return &container
}

// template struct에 내용 넣어 객체 만드는 함수
func NewTemplate(labelName map[string]string, container container) *template {
	template := template{labelName: labelName, container: container}
	return &template
}

// Dp struct에 내용 넣어 객체 만드는 함수
func NewDp(metaName string, countReplica *int32, labelName map[string]string, template template) *dp {
	dp := dp{metaName: metaName, countReplica: countReplica, labelName: labelName, template: template}
	return &dp
}

// labelname 만들기
func MakeLableName(key string, value string) map[string]string {
	ln := make(map[string]string)
	ln[key] = value
	return ln
}
