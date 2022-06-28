package main

import (
	//"strconv"
	"html/template"
	"net/http"
	"strconv"
	"github.com/26tanishabanik/yamlFileGenerator/Pods"
	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

// func main(){
// 	var metaname, zone, version, contName, contImage,topicLabel,chapterLabel,kind, apiversion,restartPolicy, volumeName, mountPath ,  contEnvName , temSpecContEnvName ,
// 	configMapKeyRefName , configMapKeyRefKey , contEnvValue , temSpecContEnvValue , imagePullPolicy , temSpecimagePullPolicy , contCommand , temSpecContCommand ,
// 	contArgstext , temSpecContArgstext, temspecTopic , initCntName , initCntImage , initCntcommand , temspecSelectorApp , temspecSelectorMatchLabelsApp , temspecSelectorMatchLabelsChapter ,
// 	temspecSelectorEnv , temspecType ,temSpecinitCntName , temSpecinitCntImage , temSpecinitCntcommand , temSpeccontname ,
// 	temSpeccontImage , temSpecvolumeName , temSpecmountPath ,
// 	temSpecconfigMapKeyRefName , temSpecconfigMapKeyRefKey  string
// 	var contPort, SpecTerminationGracePeriodSeconds,TemSpecTerminationGracePeriodSeconds, TemSpecReplicas, temspecRevisionHistoryLimit,
// 		temspecProgressDeadlineSeconds, temspecMinReadySeconds, targetPort,port, nodePort,temSpeccontPort, temSpectargetPort,temSpecport, temSpecnodePort int64
// 	var readOnly, temSpecreadOnly, contTty, contStdin, temSpeccontTty, temSpeccontStdin bool

// 	fmt.Println("Enter apiversion...")
// 	fmt.Scanln(&apiversion)
// 	kind := widget.NewEntry()
// 	fmt.Println("Enter kind...")
// 	metaname := widget.NewEntry()
// 	metaname.SetPlaceHolder("Enter metaname...")
// 	zone := widget.NewEntry()
// 	zone.SetPlaceHolder("Enter zone...")
// 	version := widget.NewEntry()
// 	version.SetPlaceHolder("Enter version...")
// 	contName := widget.NewEntry()
// 	contName.SetPlaceHolder("Enter container name...")
// 	contImage := widget.NewEntry()
// 	contImage.SetPlaceHolder("Enter container image...")
// 	topicLabel := widget.NewEntry()
// 	topicLabel.SetPlaceHolder("Enter topic...")
// 	chapterLabel := widget.NewEntry()
// 	chapterLabel.SetPlaceHolder("Enter chapter.....")
// 	contPort := widget.NewEntry()
// 	contPort.SetPlaceHolder("Enter container port.....")
// 	targetPort := widget.NewEntry()
// 	targetPort.SetPlaceHolder("Enter target port.....")
// 	nodePort := widget.NewEntry()
// 	nodePort.SetPlaceHolder("Enter node port.....")
// 	restartPolicy := widget.NewEntry()
// 	restartPolicy.SetPlaceHolder("Enter restart policy for spec.....")
// 	SpecTerminationGracePeriodSeconds := widget.NewEntry()
// 	SpecTerminationGracePeriodSeconds.SetPlaceHolder("Enter restart policy for spec.....")
// 	volumeName := widget.NewEntry()
// 	volumeName.SetPlaceHolder("Enter volume name for container.....")
// 	mounthPath := widget.NewEntry()
// 	mounthPath.SetPlaceHolder("Enter mount path of the volume for container")
// 	readOnly := widget.NewEntry()
// 	readOnly.SetPlaceHolder("Enter readonly of the volume for container")
// 	contEnvName := widget.NewEntry()
// 	contEnvName.SetPlaceHolder("Enter container environment name....")
// 	temSpecContEnvName := widget.NewEntry()
// 	temSpecContEnvName.SetPlaceHolder("Enter container environment name for template...")
// 	configMapKeyRefName := widget.NewEntry()
// 	configMapKeyRefName.SetPlaceHolder("Enter spec.container.env.valuefrom.configMapKeyRef.name....")
// 	configMapKeyRefKey := widget.NewEntry()
// 	configMapKeyRefKey.SetPlaceHolder("Enter spec.container.env.valuefrom.configMapKeyRef.key....")
// 	contEnvValue := widget.NewEntry()
// 	contEnvValue.SetPlaceHolder("Enter spec.container.Env.value....")
// 	temSpecContEnvValue := widget.NewEntry()
// 	temSpecContEnvValue.SetPlaceHolder("Enter spec.template.container.env.value....")
// 	imagePullPolicy := widget.NewEntry()
// 	imagePullPolicy.SetPlaceHolder("Enter spec.container.imagePullPolicy....")
// 	temSpecimagePullPolicy := widget.NewEntry()
// 	temSpecimagePullPolicy.SetPlaceHolder("Enter spec.template.container.imagePullPolicy....")
// 	contCommand := widget.NewEntry()
// 	contCommand.SetPlaceHolder("Enter spec.container.Command....")
// 	temSpecContCommand := widget.NewEntry()
// 	temSpecContCommand.SetPlaceHolder("Enter spec.template.container.command...")
// 	contArgs := widget.NewEntry()
// 	contArgs.SetPlaceHolder("Enter spec.container.args....")
// 	temSpecContArgs := widget.NewEntry()
// 	temSpecContArgs.SetPlaceHolder("Enter spec.template.container.args.....")
// 	contTty := widget.NewEntry()
// 	contTty.SetPlaceHolder("Enter spec.container.tty.....")
// 	contStdin := widget.NewEntry()
// 	contStdin.SetPlaceHolder("Enter spec.container.stdin...")
// 	TemSpecTerminationGracePeriodSeconds := widget.NewEntry()
// 	TemSpecTerminationGracePeriodSeconds.SetPlaceHolder("Enter spec.templates.spec.temSpecTerminationGracePeriodSeconds........")
// 	TemSpecReplicas := widget.NewEntry()
// 	TemSpecReplicas.SetPlaceHolder("Enter spec.templates.replicas....")
// 	temspecTopic := widget.NewEntry()
// 	temspecTopic.SetPlaceHolder("Enter spec.templates.spec.topic.....")
// 	initCntName := widget.NewEntry()
// 	initCntName.SetPlaceHolder("Enter spec.initcontainers.name.....")
// 	initCntImage := widget.NewEntry()
// 	initCntImage.SetPlaceHolder("Enter spec.initcontainers.image...")
// 	initCntcommand := widget.NewEntry()
// 	initCntcommand.SetPlaceHolder("Enter spec.initcontainers.commands....")
// 	temspecSelectorApp := widget.NewEntry()
// 	temspecSelectorApp.SetPlaceHolder("Enter spec.templates.selector.app....")
// 	temspecSelectorMatchLabelsApp := widget.NewEntry()
// 	temspecSelectorMatchLabelsApp.SetPlaceHolder("Enter spec.templates.selector.matchLabels.app....")
// 	temspecSelectorMatchLabelsChapter := widget.NewEntry()
// 	temspecSelectorMatchLabelsChapter.SetPlaceHolder("Enter spec.templates.selector.matchLabels.chapter....")
// 	temspecSelectorEnv := widget.NewEntry()
// 	temspecSelectorEnv.SetPlaceHolder("Enter spec.templates.Selector.env....")
// 	temspecRevisionHistoryLimit := widget.NewEntry()
// 	temspecRevisionHistoryLimit.SetPlaceHolder("Enter spec.templates.spec.revisionHistoryLimit....")
// 	temspecProgressDeadlineSeconds := widget.NewEntry()
// 	temspecProgressDeadlineSeconds.SetPlaceHolder("Enter spec.templates.spec.deadlineSeconds....")
// 	temspecMinReadySeconds := widget.NewEntry()
// 	temspecMinReadySeconds.SetPlaceHolder("Enter spec.templates.spec.minReadySeconds....")
// 	temspecType := widget.NewEntry()
// 	temspecType.SetPlaceHolder("Enter spec.templates.spec.type")
// 	port := widget.NewEntry()
// 	port.SetPlaceHolder("Enter spec.container.ports.port....")
// 	temSpecinitCntName := widget.NewEntry()
// 	temSpecinitCntName.SetPlaceHolder("Enter spec.templates.spec.initcontainers.name....")
// 	temSpecinitCntImage := widget.NewEntry()
// 	temSpecinitCntImage.SetPlaceHolder("Enter spec.templates.spec.initcontainers.image....")
// 	temSpecinitCntcommand := widget.NewEntry()
// 	temSpecinitCntcommand.SetPlaceHolder("Enter spec.templates.initcontainers.commands....")
// 	temSpeccontname := widget.NewEntry()
// 	temSpeccontname.SetPlaceHolder("Enter spec.templates.spec.containers.name....")
// 	temSpeccontImage := widget.NewEntry()
// 	temSpeccontImage.SetPlaceHolder("Enter spec.templates.spec.containers.image....")
// 	temSpeccontPort := widget.NewEntry()
// 	temSpeccontPort.SetPlaceHolder("Enter spec.templates.spec.containers.ports.containerPort.....")
// 	temSpectargetPort := widget.NewEntry()
// 	temSpectargetPort.SetPlaceHolder("Enter spec.templates.spec.containers.ports.targetPort.....")
// 	temSpecport := widget.NewEntry()
// 	temSpecport.SetPlaceHolder("Enter spec.templates.spec.container.ports.port....")
// 	temSpecnodePort := widget.NewEntry()
// 	temSpecnodePort.SetPlaceHolder("Enter spec.templates.spec.container.ports.nodePort....")
// 	temSpecvolumeName := widget.NewEntry()
// 	temSpecvolumeName.SetPlaceHolder("Enter spec.templates.spec.containers.volumeMount.name.....")
// 	temSpecmountPath := widget.NewEntry()
// 	temSpecmountPath.SetPlaceHolder("Enter spec.templates.spec.containers.volumeMount.mountPath.....")
// 	temSpecreadOnly := widget.NewEntry()
// 	temSpecreadOnly.SetPlaceHolder("Enter spec.templates.spec.containers.volumeMount.readOnly.....")
// 	temSpecconfigMapKeyRefName := widget.NewEntry()
// 	temSpecconfigMapKeyRefName.SetPlaceHolder("Enter spec.templates.spec.containers.env.configmapKeyRef.name....")
// 	temSpecconfigMapKeyRefKey := widget.NewEntry()
// 	temSpecconfigMapKeyRefKey.SetPlaceHolder("Enter spec.templates.spec.containers.env.configmapKeyRef.key....")
// 	temSpeccontTty := widget.NewEntry()
// 	temSpeccontTty.SetPlaceHolder("Enter spec.templates.spec.containers.tty....")
// 	temSpeccontStdin := widget.NewEntry()
// 	temSpeccontStdin.SetPlaceHolder("Enter spec.templates.spec.containers.stdin....")
// }

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("KUBERNETES YAML FILES")
// 	myWindow.Resize(fyne.NewSize(800, 1000))
// 	title := canvas.NewText("Generate yaml file", color.White)
// 	title.TextStyle = fyne.TextStyle{
// 		Bold: true,
// 	}
// 	title.Alignment = fyne.TextAlignCenter
// 	title.TextSize = 24
// 	apiversion := widget.NewEntry()
// 	apiversion.SetPlaceHolder("Enter apiversion...")
// 	kind := widget.NewEntry()
// 	kind.SetPlaceHolder("Enter kind...")
// 	metaname := widget.NewEntry()
// 	metaname.SetPlaceHolder("Enter metaname...")
// 	zone := widget.NewEntry()
// 	zone.SetPlaceHolder("Enter zone...")
// 	version := widget.NewEntry()
// 	version.SetPlaceHolder("Enter version...")
// 	contName := widget.NewEntry()
// 	contName.SetPlaceHolder("Enter container name...")
// 	contImage := widget.NewEntry()
// 	contImage.SetPlaceHolder("Enter container image...")
// 	topicLabel := widget.NewEntry()
// 	topicLabel.SetPlaceHolder("Enter topic...")
// 	chapterLabel := widget.NewEntry()
// 	chapterLabel.SetPlaceHolder("Enter chapter.....")
// 	contPort := widget.NewEntry()
// 	contPort.SetPlaceHolder("Enter container port.....")
// 	targetPort := widget.NewEntry()
// 	targetPort.SetPlaceHolder("Enter target port.....")
// 	nodePort := widget.NewEntry()
// 	nodePort.SetPlaceHolder("Enter node port.....")
// 	restartPolicy := widget.NewEntry()
// 	restartPolicy.SetPlaceHolder("Enter restart policy for spec.....")
// 	SpecTerminationGracePeriodSeconds := widget.NewEntry()
// 	SpecTerminationGracePeriodSeconds.SetPlaceHolder("Enter restart policy for spec.....")
// 	volumeName := widget.NewEntry()
// 	volumeName.SetPlaceHolder("Enter volume name for container.....")
// 	mounthPath := widget.NewEntry()
// 	mounthPath.SetPlaceHolder("Enter mount path of the volume for container")
// 	readOnly := widget.NewEntry()
// 	readOnly.SetPlaceHolder("Enter readonly of the volume for container")
// 	contEnvName := widget.NewEntry()
// 	contEnvName.SetPlaceHolder("Enter container environment name....")
// 	temSpecContEnvName := widget.NewEntry()
// 	temSpecContEnvName.SetPlaceHolder("Enter container environment name for template...")
// 	configMapKeyRefName := widget.NewEntry()
// 	configMapKeyRefName.SetPlaceHolder("Enter spec.container.env.valuefrom.configMapKeyRef.name....")
// 	configMapKeyRefKey := widget.NewEntry()
// 	configMapKeyRefKey.SetPlaceHolder("Enter spec.container.env.valuefrom.configMapKeyRef.key....")
// 	contEnvValue := widget.NewEntry()
// 	contEnvValue.SetPlaceHolder("Enter spec.container.Env.value....")
// 	temSpecContEnvValue := widget.NewEntry()
// 	temSpecContEnvValue.SetPlaceHolder("Enter spec.template.container.env.value....")
// 	imagePullPolicy := widget.NewEntry()
// 	imagePullPolicy.SetPlaceHolder("Enter spec.container.imagePullPolicy....")
// 	temSpecimagePullPolicy := widget.NewEntry()
// 	temSpecimagePullPolicy.SetPlaceHolder("Enter spec.template.container.imagePullPolicy....")
// 	contCommand := widget.NewEntry()
// 	contCommand.SetPlaceHolder("Enter spec.container.Command....")
// 	temSpecContCommand := widget.NewEntry()
// 	temSpecContCommand.SetPlaceHolder("Enter spec.template.container.command...")
// 	contArgs := widget.NewEntry()
// 	contArgs.SetPlaceHolder("Enter spec.container.args....")
// 	temSpecContArgs := widget.NewEntry()
// 	temSpecContArgs.SetPlaceHolder("Enter spec.template.container.args.....")
// 	contTty := widget.NewEntry()
// 	contTty.SetPlaceHolder("Enter spec.container.tty.....")
// 	contStdin := widget.NewEntry()
// 	contStdin.SetPlaceHolder("Enter spec.container.stdin...")
// 	TemSpecTerminationGracePeriodSeconds := widget.NewEntry()
// 	TemSpecTerminationGracePeriodSeconds.SetPlaceHolder("Enter spec.templates.spec.temSpecTerminationGracePeriodSeconds........")
// 	TemSpecReplicas := widget.NewEntry()
// 	TemSpecReplicas.SetPlaceHolder("Enter spec.templates.replicas....")
// 	temspecTopic := widget.NewEntry()
// 	temspecTopic.SetPlaceHolder("Enter spec.templates.spec.topic.....")
// 	initCntName := widget.NewEntry()
// 	initCntName.SetPlaceHolder("Enter spec.initcontainers.name.....")
// 	initCntImage := widget.NewEntry()
// 	initCntImage.SetPlaceHolder("Enter spec.initcontainers.image...")
// 	initCntcommand := widget.NewEntry()
// 	initCntcommand.SetPlaceHolder("Enter spec.initcontainers.commands....")
// 	temspecSelectorApp := widget.NewEntry()
// 	temspecSelectorApp.SetPlaceHolder("Enter spec.templates.selector.app....")
// 	temspecSelectorMatchLabelsApp := widget.NewEntry()
// 	temspecSelectorMatchLabelsApp.SetPlaceHolder("Enter spec.templates.selector.matchLabels.app....")
// 	temspecSelectorMatchLabelsChapter := widget.NewEntry()
// 	temspecSelectorMatchLabelsChapter.SetPlaceHolder("Enter spec.templates.selector.matchLabels.chapter....")
// 	temspecSelectorEnv := widget.NewEntry()
// 	temspecSelectorEnv.SetPlaceHolder("Enter spec.templates.Selector.env....")
// 	temspecRevisionHistoryLimit := widget.NewEntry()
// 	temspecRevisionHistoryLimit.SetPlaceHolder("Enter spec.templates.spec.revisionHistoryLimit....")
// 	temspecProgressDeadlineSeconds := widget.NewEntry()
// 	temspecProgressDeadlineSeconds.SetPlaceHolder("Enter spec.templates.spec.deadlineSeconds....")
// 	temspecMinReadySeconds := widget.NewEntry()
// 	temspecMinReadySeconds.SetPlaceHolder("Enter spec.templates.spec.minReadySeconds....")
// 	temspecType := widget.NewEntry()
// 	temspecType.SetPlaceHolder("Enter spec.templates.spec.type")
// 	port := widget.NewEntry()
// 	port.SetPlaceHolder("Enter spec.container.ports.port....")
// 	temSpecinitCntName := widget.NewEntry()
// 	temSpecinitCntName.SetPlaceHolder("Enter spec.templates.spec.initcontainers.name....")
// 	temSpecinitCntImage := widget.NewEntry()
// 	temSpecinitCntImage.SetPlaceHolder("Enter spec.templates.spec.initcontainers.image....")
// 	temSpecinitCntcommand := widget.NewEntry()
// 	temSpecinitCntcommand.SetPlaceHolder("Enter spec.templates.initcontainers.commands....")
// 	temSpeccontname := widget.NewEntry()
// 	temSpeccontname.SetPlaceHolder("Enter spec.templates.spec.containers.name....")
// 	temSpeccontImage := widget.NewEntry()
// 	temSpeccontImage.SetPlaceHolder("Enter spec.templates.spec.containers.image....")
// 	temSpeccontPort := widget.NewEntry()
// 	temSpeccontPort.SetPlaceHolder("Enter spec.templates.spec.containers.ports.containerPort.....")
// 	temSpectargetPort := widget.NewEntry()
// 	temSpectargetPort.SetPlaceHolder("Enter spec.templates.spec.containers.ports.targetPort.....")
// 	temSpecport := widget.NewEntry()
// 	temSpecport.SetPlaceHolder("Enter spec.templates.spec.container.ports.port....")
// 	temSpecnodePort := widget.NewEntry()
// 	temSpecnodePort.SetPlaceHolder("Enter spec.templates.spec.container.ports.nodePort....")
// 	temSpecvolumeName := widget.NewEntry()
// 	temSpecvolumeName.SetPlaceHolder("Enter spec.templates.spec.containers.volumeMount.name.....")
// 	temSpecmountPath := widget.NewEntry()
// 	temSpecmountPath.SetPlaceHolder("Enter spec.templates.spec.containers.volumeMount.mountPath.....")
// 	temSpecreadOnly := widget.NewEntry()
// 	temSpecreadOnly.SetPlaceHolder("Enter spec.templates.spec.containers.volumeMount.readOnly.....")
// 	temSpecconfigMapKeyRefName := widget.NewEntry()
// 	temSpecconfigMapKeyRefName.SetPlaceHolder("Enter spec.templates.spec.containers.env.configmapKeyRef.name....")
// 	temSpecconfigMapKeyRefKey := widget.NewEntry()
// 	temSpecconfigMapKeyRefKey.SetPlaceHolder("Enter spec.templates.spec.containers.env.configmapKeyRef.key....")
// 	temSpeccontTty := widget.NewEntry()
// 	temSpeccontTty.SetPlaceHolder("Enter spec.templates.spec.containers.tty....")
// 	temSpeccontStdin := widget.NewEntry()
// 	temSpeccontStdin.SetPlaceHolder("Enter spec.templates.spec.containers.stdin....")

// 	factText := widget.NewLabel("Yaml File")
// 	factText.Wrapping = fyne.TextWrapWord
// 	inputs := container.NewVBox(title, apiversion, kind, metaname, zone, version,
// 		contName, contImage, contPort, targetPort, nodePort, topicLabel, chapterLabel, contPort,
// 		targetPort,nodePort,restartPolicy,SpecTerminationGracePeriodSeconds)
// 	content := container.NewVScroll(container.NewVBox(inputs, widget.NewButton("Show Yaml", func() {

// 			log.Println("Content")
// 			contPort, _ := strconv.ParseInt(contPort.Text, 10, 64)
// 			SpecTerminationGracePeriodSeconds, _ := strconv.ParseInt(SpecTerminationGracePeriodSeconds.Text, 10, 64)
// 			readOnly, _ := strconv.ParseBool(readOnly.Text)
// 			TemSpecTerminationGracePeriodSeconds, _ := strconv.ParseInt(TemSpecTerminationGracePeriodSeconds.Text, 10, 64)
// 			TemSpecReplicas, _ := strconv.ParseInt(TemSpecReplicas.Text, 10, 64)
// 			temspecRevisionHistoryLimit, _ := strconv.ParseInt(temspecRevisionHistoryLimit.Text, 10, 64)
// 			temspecProgressDeadlineSeconds, _ := strconv.ParseInt(temspecProgressDeadlineSeconds.Text, 10, 64)
// 			temspecMinReadySeconds, _ := strconv.ParseInt(temspecMinReadySeconds.Text, 10, 64)
// 			targetPort, _ := strconv.ParseInt(targetPort.Text, 10, 64)
// 			port, _ := strconv.ParseInt(port.Text, 10, 64)
// 			nodePort, _ := strconv.ParseInt(nodePort.Text, 10, 64)
// 			contTty, _ := strconv.ParseBool(contTty.Text)
// 			contStdin, _ := strconv.ParseBool(contStdin.Text)
// 			temSpeccontPort, _ := strconv.ParseInt(temSpeccontPort.Text, 10, 64)
// 			temSpectargetPort, _ := strconv.ParseInt(temSpectargetPort.Text, 10, 64)
// 			temSpecport, _ := strconv.ParseInt(temSpecport.Text, 10, 64)
// 			temSpecnodePort, _ := strconv.ParseInt(temSpecnodePort.Text, 10, 64)
// 			temSpecreadOnly, _ := strconv.ParseBool(temSpecreadOnly.Text)
// 			temSpeccontTty, _ := strconv.ParseBool(temSpeccontTty.Text)
// 			temSpeccontStdin, _ := strconv.ParseBool(temSpeccontStdin.Text)

// 			yamlData := Pods.SendYaml(apiversion.Text, kind.Text, topicLabel.Text, chapterLabel.Text, metaname.Text, zone.Text, version.Text, contName.Text, contImage.Text, contPort,
// 				restartPolicy.Text,volumeName.Text , mounthPath.Text, readOnly, contEnvName.Text, temSpecContEnvName.Text,configMapKeyRefName.Text, configMapKeyRefKey.Text,contEnvValue.Text, temSpecContEnvValue.Text,
// 				imagePullPolicy.Text, temSpecimagePullPolicy.Text,contCommand.Text , temSpecContCommand.Text,contArgs.Text,temSpecContArgs.Text,contTty,contStdin, SpecTerminationGracePeriodSeconds,TemSpecTerminationGracePeriodSeconds,
// 				TemSpecReplicas, temspecTopic.Text, initCntName.Text, initCntImage.Text, initCntcommand.Text, temspecSelectorApp.Text, temspecSelectorMatchLabelsApp.Text, temspecSelectorMatchLabelsChapter.Text,
// 				temspecSelectorEnv.Text, temspecRevisionHistoryLimit, temspecProgressDeadlineSeconds, temspecMinReadySeconds, temspecType.Text, targetPort, port, nodePort,
// 				temSpecinitCntName.Text, temSpecinitCntImage.Text, temSpecinitCntcommand.Text, temSpeccontname.Text, temSpeccontImage.Text,
// 				temSpeccontPort, temSpectargetPort, temSpecport, temSpecnodePort, temSpecvolumeName.Text, temSpecmountPath.Text, temSpecreadOnly, temSpecconfigMapKeyRefName.Text, temSpecconfigMapKeyRefKey.Text,
// 				temSpeccontTty, temSpeccontStdin)

// 			factText.SetText(string(yamlData))
// 		})))
// 	content.Direction = container.ScrollHorizontalOnly
// 	//vBox := container.New(layout.NewVBoxLayout(), content, factText)
// 	myWindow.SetContent(container.NewHSplit(content, factText))

// 	myWindow.ShowAndRun()
// }

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/home/tanisha/Desktop/Projects/yamlFileGenerator/templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)

}
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	//io.WriteString(w, "Hello, I am Tanisha")
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	apiversion := r.FormValue("apiversion")
	kind := r.FormValue("kind")
	topicLabel := r.FormValue("topicLabel")
	chapterLabel := r.FormValue("chapterLabel")
	metaname := r.FormValue("metaname")
	zone := r.FormValue("zone")
	version := r.FormValue("version")
	contImage := r.FormValue("contImage")
	contName := r.FormValue("contName")
	contPort := r.FormValue("contPort")

	// d := struct {
	// 	ApiVersion string
	// 	Kind string
	// }{
	// 	ApiVersion: apiversion,
	// 	Kind: kind,
	// }
	contPortint, _ := strconv.ParseInt(contPort, 10, 64)
	if kind == "Pod" {
		yamlData := Pods.SendYaml(apiversion, kind, topicLabel, chapterLabel, metaname, zone, version, contName, contImage, contPortint)
		//yamlData = strings.Replace(yamlData, "\n", `\<br/>\`, -1)
		d := struct {
			YamlData string
		}{
			YamlData: yamlData,
		}
		tpl.ExecuteTemplate(w, "processor.gohtml", d)
	}

}
