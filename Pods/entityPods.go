package Pods

import (
	"fmt"
	//"strings"

	"gopkg.in/yaml.v2"
)

//"strings"
//"gopkg.in/yaml.v2"

type ConfigMapKeyReff struct {
	Name string `yaml:"name,omitempty"`
	Key  string `yaml:"key,omitempty"`
}
type ValueFromm struct {
	ConfigMapKeyRef ConfigMapKeyReff `yaml:"configMapKeyRef,omitempty"`
}
type VolumeMountss struct {
	Name      string `yaml:"name,omitempty"`
	MountPath string `yaml:"mountPath,omitempty"`
	ReadOnly  bool   `yaml:"readOnly,omitempty"`
}
type Envv struct {
	Name      string     `yaml:"name"`
	ValueFrom ValueFromm `yaml:"valueFrom,omitempty"`
	Value     string     `yaml:"value"`
}
type NodePortss struct {
	NodePort int64 `yaml:"nodePort,omitempty"`
}
type Portssss struct {
	Port          int64      `yaml:"port,omitempty"`
	ContainerPort int64      `yaml:"containerPort,omitempty"`
	TargetPort    int64      `yaml:"targetPort,omitempty"`
	NodePorts     NodePortss `yaml:",inline"`
	Protocol      string     `yaml:"protocol,omitempty"`
	Name          string     `yaml:"name,omitempty"`
}
type Containers struct {
	Name            string          `yaml:name,omitempty`
	Image           string          `yaml:"image,omitempty"`
	Portsss         []Portssss      `yaml:"ports,omitempty"`
	VolumeMount     []VolumeMountss `yaml:"volumeMounts,omitempty"`
	Env             []Envv          `yaml:"env,omitempty"`
	ImagePullPolicy string          `yaml:"imagePullPolicy,omitempty"`
	Command         string          `yaml:"command,omitempty"`
	Args            []string        `yaml:"args,omitempty"`
	Tty             bool            `yaml:"tty,omitempty"`
	Stdin           bool            `yaml:"stdin,omitempty"`
}
type Labelss struct {
	// App     string `yaml:"app,omitempty"`
	// Env     string `yaml:"env,omitempty"`
	// Role    string `yaml:"role,omitempty"`
	Zone    string `yaml:"zone,omitempty"`
	Version string `yaml:"version,omitempty"`
	Topic   string `yaml:"topic,omitempty"`
	Chapter string `yaml:"chapter,omitempty"`
}
type MetaData struct {
	//Namespace   string       `yaml:"namespace,omitempty"`
	Name string `yaml:"name,omitempty"`
	//Annotations Annotationss `yaml:"annotations,inline"`
	Labels Labelss `yaml:"labels,omitempty"`
}
type Containersss struct {
	Containerss []Containers `yaml:"containers,omitempty"`
}
type ConfigMapp struct {
	Name string `yaml:"name,omitempty"`
}
type Secrett struct {
	SecretName string `yaml:"secretName,omitempty"`
}
type Empty struct{}
type PersistentVolumeClaimm struct {
	ClaimName string `yaml:"claimName,omitempty"`
}
type Volumess struct {
	Name                  string                 `yaml:"name,omitempty"`
	ConfigMap             ConfigMapp             `yaml:"configMap,omitempty"`
	MountPath             string                 `yaml:"mountPath,omitempty"`
	EmptyDir              Empty                  `yaml:"emptyDir,omitempty"`
	PersistentVolumeClaim PersistentVolumeClaimm `yaml:"persistentVolumeClaim,omitempty"`
	Secret                Secrett                `yaml:"secret,inline"`
}
type VolumeInline struct {
	Volumes []Volumess `yaml:"volumes,omitempty"`
}
type InitContainers struct {
	Name    string `yaml:"name"`
	Image   string `yaml:"image"`
	Command string `yaml:"command,omitempty"`
}
type InitContainersss struct {
	InitContainerss []InitContainers `yaml:"initContainers,omitempty"`
}
type MatchLabels struct {
	App     string `yaml:"app,omitempty"`
	Chapter string `yaml:"chapter,omitempty"`
}
type Selectorr struct {
	App        string      `yaml:"app,omitempty"`
	MatchLabel MatchLabels `yaml:"matchLabels,omitempty"`
	Envv       string      `yaml:"env,omitempty"`
}

type TemSpecs struct {
	TerminationGracePeriodSeconds int64  `yaml:"terminationGracePeriodSeconds,omitempty"`
	Replicas                      int64  `yaml:"replicas,omitempty"`
	Topic                         string `yaml:"topic,omitempty"`
	//Rules                   Rulessss         `yaml:",inline"`
	InitContainer           InitContainersss `yaml:",inline"`
	Container               Containersss     `yaml:",inline"`
	Selector                Selectorr        `yaml:"selector,inline"`
	RevisionHistoryLimit    int64            `yaml:"revisionHistoryLimit,omitempty"`
	ProgressDeadlineSeconds int64            `yaml:"progressDeadlineSeconds,omitempty"`
	MinReadySeconds         int64            `yaml:"minReadySeconds,omitempty"`
	//Strategy                Strategyy        `yaml:"strategy,inline"`
	Type string `yaml:"type,omitempty"`
}
type Template struct {
	Metadata MetaData `yaml:"metadata,omitempty"`
	Spec     TemSpecs `yaml:"spec,omitempty"`
}

type TemplateInline struct {
	Templates Template `yaml:"template,omitempty"`
}
type Specs struct {
	RestartPolicy string `yaml:"restartPolicy,omitempty"`
	//Csi 					Csii 			`yaml:"csi,omitempty"`
	// PersistentVolumeReclaimPolicy	string `yaml:"persistentVolumeReclaimPolicy,omitempty"`
	// GcePersistentDisk GcePersistentDiskk 	`yaml:"gcePersistentDisk,omitempty"`
	// Capacity 				Capacityy		 `yaml:"capacity,omitempty"`
	// AccessModes 			string         	 `yaml:"accessModes,omitempty"`
	// StorageClassName 		string 			 `yaml:"storageClassName,omitempty"`
	// Resources 				Resourcess 		 `yaml:"resources,inline"`
	// ClusterIP 				string 			 `yaml:"clusterIP,omitempty"`
	TerminationGracePeriodSeconds int64 `yaml:"terminationGracePeriodSeconds,omitempty"`
	// SpecPorts 				SpecPortInline 	 `yaml:",inline"`
	// IpFamilyPolicy 			string 			 `yaml:"ipFamilyPolicy,omitempty"`
	// Controller              string           `yaml:"controller,omitempty"`
	// IngressClassName        string           `yaml:"ingressClassName,omitempty"`
	// Replicas                int              `yaml:"replicas,omitempty"`
	// Topic                   string           `yaml:"topic,omitempty"`
	// Rules                   Rulessss         `yaml:",inline"`
	InitContainer InitContainersss `yaml:",inline"`
	Volumes       VolumeInline     `yaml:",inline"`
	Container     Containersss     `yaml:",inline"`
	// Selector                SelectorInline   `yaml:",inline"`
	// RevisionHistoryLimit    int              `yaml:"revisionHistoryLimit,omitempty"`
	// ProgressDeadlineSeconds int              `yaml:"progressDeadlineSeconds,omitempty"`
	// MinReadySeconds         int              `yaml:"minReadySeconds,omitempty"`
	// Strategy                Strategyy        `yaml:"strategy,inline"`
	// Type                    string           `yaml:"type,omitempty"`
	Templatess TemplateInline `yaml:",inline"`
	// VolumeClaimTemplatess   VolumeClaimTemplateInline `yaml:",inline"`
}
type Main struct {
	//Type string `yaml:"-,omitempty"`
	//Provisioner string `yaml:"provisioner,omitempty"`
	//VolumeBindingMode string `yaml:"volumeBindingMode,omitempty"`
	//AllowVolumeExpansion string `yaml:"allowVolumeExpansion,omitempty"`
	//Parameter Parameters `yaml:"parameters,inline"`
	Apiversion string   `yaml:"apiversion"`
	Kind       string   `yaml:"kind"`
	Metadata   MetaData `yaml:"metadata"`
	Spec       Specs    `yaml:"spec"`
	// Volumes    VolumeInline `yaml:",inline"`
	// ReclaimPolicy string `yaml:"reclaimPolicy,omitempty"`
	// Data Dataa `yaml:"data,inline"`
}

func SendYaml(apiversion string, kind string, topicLabel string, chapterLabel string, metaname string,
	zone string, version string, contname string, contImage string, contPort int64) string {
	fmt.Println("Pods")
	podMain := Main{
		Apiversion: apiversion,
		Kind:       kind,
		Metadata: MetaData{
			Name: metaname,
			Labels: Labelss{
				Zone:    zone,
				Version: version,
				Topic:   topicLabel,
				Chapter: chapterLabel,
			},
		},
		Spec: Specs{
			Container: Containersss{
				[]Containers{
					{
						Name:  contname,
						Image: contImage,
						Portsss: []Portssss{
							{
								ContainerPort: contPort,
							},
						},
					},
				},
			},
		},
	}
	yamlData, err := yaml.Marshal(&podMain)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}
	fmt.Println(string(yamlData))
	return string(yamlData)
}

// func SendYaml(apiversion string, kind string, topicLabel string, chapterLabel string, metaname string,
// 	zone string, version string, contname string, contImage string, contPort int64,
// 	restartPolicy string, volumeName string, mountPath string, readOnly bool, contEnvName string, temSpecContEnvName string,
// 	configMapKeyRefName string, configMapKeyRefKey string, contEnvValue string, temSpecContEnvValue string, imagePullPolicy string, temSpecimagePullPolicy string, contCommand string, temSpecContCommand string,
// 	contArgstext string, temSpecContArgstext string, contTty bool, contStdin bool, SpecTerminationGracePeriodSeconds int64,
// 	TemSpecTerminationGracePeriodSeconds int64, TemSpecReplicas int64, temspecTopic string, initCntName string,
// 	initCntImage string, initCntcommand string, temspecSelectorApp string, temspecSelectorMatchLabelsApp string, temspecSelectorMatchLabelsChapter string,
// 	temspecSelectorEnv string, temspecRevisionHistoryLimit int64, temspecProgressDeadlineSeconds int64, temspecMinReadySeconds int64, temspecType string,
// 	targetPort int64, port int64, nodePort int64, temSpecinitCntName string, temSpecinitCntImage string, temSpecinitCntcommand string, temSpeccontname string,
// 	temSpeccontImage string, temSpeccontPort int64, temSpectargetPort int64, temSpecport int64, temSpecnodePort int64, temSpecvolumeName string, temSpecmountPath string, temSpecreadOnly bool,
// 	temSpecconfigMapKeyRefName string, temSpecconfigMapKeyRefKey string, temSpeccontTty bool, temSpeccontStdin bool) string {
// 	fmt.Println("Pods")
// 	var contArgs []string
// 	var temSpecContArgs []string
// 	if contArgstext == "" {
// 		contArgs = []string{}
// 	} else {
// 		contArgs = append(contArgs, strings.Split(contArgstext, "\n")...)
// 		// for _, v := range strings.Split(contArgstext, "\n") {
// 		// 	contArgs = append(contArgs, v)
// 		// }
// 	}
// 	if temSpecContArgstext == "" {
// 		temSpecContArgs = []string{}
// 	} else {
// 		temSpecContArgs = append(temSpecContArgs, strings.Split(temSpecContArgstext, "\n")...)
// 		// for _, v := range strings.Split(temSpecContArgstext, "\n") {
// 		// 	contArgs = append(temSpecContArgs, v)
// 		// }
// 	}

// 	podMain := Main{
// 		Apiversion: apiversion,
// 		Kind:       kind,
// 		Metadata: MetaData{
// 			Name: metaname,
// 			Labels: Labelss{
// 				Zone:    zone,
// 				Version: version,
// 				Topic:   topicLabel,
// 				Chapter: chapterLabel,
// 			},
// 		},
// 		Spec: Specs{
// 			RestartPolicy:                 restartPolicy,
// 			TerminationGracePeriodSeconds: SpecTerminationGracePeriodSeconds,
// 			InitContainer: InitContainersss{
// 				InitContainerss: []InitContainers{
// 					{
// 						Name:    initCntName,
// 						Image:   initCntImage,
// 						Command: initCntcommand,
// 					},
// 				},
// 			},
// 			Container: Containersss{
// 				[]Containers{
// 					{
// 						Name:  contname,
// 						Image: contImage,
// 						Portsss: []Portssss{
// 							{
// 								ContainerPort: contPort,
// 								TargetPort:    targetPort,
// 								Port:          port,
// 								NodePorts: NodePortss{
// 									NodePort: nodePort,
// 								},
// 							},
// 						},
// 						VolumeMount: []VolumeMountss{
// 							{
// 								Name:      volumeName,
// 								MountPath: mountPath,
// 								ReadOnly:  readOnly,
// 							},
// 						},
// 						Env: []Envv{
// 							{
// 								Name: contEnvName,
// 								ValueFrom: ValueFromm{
// 									ConfigMapKeyReff{
// 										Name: configMapKeyRefName,
// 										Key:  configMapKeyRefKey,
// 									},
// 								},
// 								Value: contEnvValue,
// 							},
// 						},
// 						ImagePullPolicy: imagePullPolicy,
// 						Command:         contCommand,
// 						Args:            contArgs,
// 						Tty:             contTty,
// 						Stdin:           contStdin,
// 					},
// 				},
// 			},
// 			Templatess: TemplateInline{
// 				Template{
// 					Metadata: MetaData{
// 						Name: metaname,
// 						Labels: Labelss{
// 							Zone:    zone,
// 							Version: version,
// 							Topic:   topicLabel,
// 							Chapter: chapterLabel,
// 						},
// 					},
// 					Spec: TemSpecs{
// 						TerminationGracePeriodSeconds: TemSpecTerminationGracePeriodSeconds,
// 						Replicas:                      TemSpecReplicas,
// 						Topic:                         temspecTopic,
// 						InitContainer: InitContainersss{
// 							InitContainerss: []InitContainers{
// 								{
// 									Name:    temSpecinitCntName,
// 									Image:   temSpecinitCntImage,
// 									Command: temSpecinitCntcommand,
// 								},
// 							},
// 						},
// 						Container: Containersss{
// 							[]Containers{
// 								{
// 									Name:  temSpeccontname,
// 									Image: temSpeccontImage,
// 									Portsss: []Portssss{
// 										{
// 											ContainerPort: temSpeccontPort,
// 											TargetPort:    temSpectargetPort,
// 											Port:          temSpecport,
// 											NodePorts: NodePortss{
// 												NodePort: temSpecnodePort,
// 											},
// 										},
// 									},
// 									VolumeMount: []VolumeMountss{
// 										{
// 											Name:      temSpecvolumeName,
// 											MountPath: temSpecmountPath,
// 											ReadOnly:  temSpecreadOnly,
// 										},
// 									},
// 									Env: []Envv{
// 										{
// 											Name: temSpecContEnvName,
// 											ValueFrom: ValueFromm{
// 												ConfigMapKeyReff{
// 													Name: temSpecconfigMapKeyRefName,
// 													Key:  temSpecconfigMapKeyRefKey,
// 												},
// 											},
// 											Value: temSpecContEnvValue,
// 										},
// 									},
// 									ImagePullPolicy: temSpecimagePullPolicy,
// 									Command:         temSpecContCommand,
// 									Args:            temSpecContArgs,
// 									Tty:             temSpeccontTty,
// 									Stdin:           temSpeccontStdin,
// 								},
// 							},
// 						},
// 						Selector: Selectorr{
// 							App: temspecSelectorApp,
// 							MatchLabel: MatchLabels{
// 								App:     temspecSelectorMatchLabelsApp,
// 								Chapter: temspecSelectorMatchLabelsChapter,
// 							},
// 							Envv: temspecSelectorEnv,
// 						},
// 						RevisionHistoryLimit:    temspecRevisionHistoryLimit,
// 						ProgressDeadlineSeconds: temspecProgressDeadlineSeconds,
// 						MinReadySeconds:         temspecMinReadySeconds,
// 						Type:                    temspecType,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	yamlData, err := yaml.Marshal(&podMain)
// 	if err != nil {
// 		fmt.Printf("Error while Marshaling. %v", err)
// 	}
// 	fmt.Println(string(yamlData))
// 	return string(yamlData)
// }
