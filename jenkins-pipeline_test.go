package pipeline

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var cf *configFile

func ExampleJenkinsPipeline_UnmarshalJSON() {
	var fname = "Pipeline"
	var ppl JenkinsPipeline

	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file %s: %s", fname, err.Error())
		os.Exit(1)
	}

	// Decode will call UnmarshalJSON
	err = json.NewDecoder(f).Decode(&ppl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse %s: %s\n", fname, err.Error())
		os.Exit(1)
	}
}

func ExampleJenkinsPipeline() {
	f, err := os.Open("Pipeline.example")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file Pipeline.example: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	// read config from file Pipeline:
	ppl, err := NewJenkinsPipeline(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	// optionally: overwrite JenkinsServer
	if js := os.Getenv("JENKINS_SERVER_OVERWRITE"); js != "" {
		ppl.JenkinsServer = JenkinsServer(js)
	}
}

func TestBasicConfigParsing(t *testing.T) {
	testConfigFile()
}

func TestStagesParsing(t *testing.T) {
	expectations{
		{"number of stages", 6, len(testConfigFile().Stages)},

		{"stage0 name", "stage0", testConfigFile().Stages[0].Name},
		{"stage0 num of jobs", 2, len(testConfigFile().Stages[0].Jobs)},
		{"stage0 num of next-stages", 1, len(testConfigFile().Stages[0].NextStages)},
		{"stage0 num of next-manual-stages", 0, len(testConfigFile().Stages[0].NextManualStages)},

		{"stage2 name", "stage2", testConfigFile().Stages[2].Name},
		{"stage2 has one job", 1, len(testConfigFile().Stages[2].Jobs)},
		{"stage2's job is a multijob", true, testConfigFile().Stages[2].Jobs[0].isMultiJob()},
		{"stage2's multijob has two sub jobs", 2, len(testConfigFile().Stages[2].Jobs[0].SubJobs)},

		{"stage3 num of next-manual-stages", 2, len(testConfigFile().Stages[3].NextManualStages)},
	}.Run(t)
}

func TestSettingsValidations(t *testing.T) {
	_, noSettingsErr := jenkinsConfigFromString(`{}`)

	_, noJenkinsServer := jenkinsConfigFromString(`{"settings":{"git-url": "http://github.com/soundcloud/pipeline-generator"}}`)
	_, emptyJenkinsServer := jenkinsConfigFromString(`{"settings":{"git-url": "http://github.com/soundcloud/pipeline-generator", "jenkins-server": ""}}`)

	_, noGitURL := jenkinsConfigFromString(`{"settings":{"jenkins-server": "http://jenkins:8080"}}`)
	_, emptyGitURL := jenkinsConfigFromString(`{"settings":{"jenkins-server": "http://jenkins:8080", "git-url": ""}}`)

	expectations{
		{"raises error when settings is missing", ErrSettingsMissing, noSettingsErr},

		{"raises error when jenkins-server is missing", ErrJenkinsServerMissing, noJenkinsServer},
		{"raises error when jenkins-server is empty", ErrJenkinsServerMissing, emptyJenkinsServer},

		{"raises error when git-url is missing", ErrGitURLMissing, noGitURL},
		{"raises error when git-url is empty", ErrGitURLMissing, emptyGitURL},
	}.Run(t)
}

func TestPipelineCreation(t *testing.T) {
	jp, err := jenkinsConfigFromFile("tests-fixtures/test_config.json")

	job1EndPoint, job1EndPointError := jp.resources[1].projectName("test-project")
	pipelineViewEndPoint, pipelineViewEndPointError := jp.resources[16].projectName("test-project")
	expectations{
		{"no parsing error happened", nil, err},
		{"endpoint rendering for job1 worked", nil, job1EndPointError},
		{"endpoint rendering for pipeline view worked", nil, pipelineViewEndPointError},
		{"jenkins server", JenkinsServer("http://jenkins:8080"), jp.JenkinsServer},

		{"parsed all resources", 17, len(jp.resources)},
		{"job0 name", "job0", jp.resources[0].(jenkinsSingleJob).TaskName},
		{"job0 IsInitialJob", true, jp.resources[0].(jenkinsSingleJob).IsInitialJob},
		{"job0 projectNameFmt", "{{ .PipelineName }}", jp.resources[0].(jenkinsSingleJob).ProjectNameTempl},
		{"job0 stage name", "stage0", jp.resources[0].(jenkinsSingleJob).StageName},
		{"job0 cmd", "\n# change to working dir:\ncd subdir\n\n\n# job setup\nexport VAR=foobar\n\n# job\necho 'job0'; mv spec spec2", jp.resources[0].(jenkinsSingleJob).Command},
		{"job0 git url", "http://github.com/kesselborn/tuev", jp.resources[0].(jenkinsSingleJob).GitURL},
		{"job0 git branch is set to master", "master", jp.resources[0].(jenkinsSingleJob).BranchSpecifier},
		{"job0 slave label", "master", jp.resources[0].(jenkinsSingleJob).SlaveLabel},
		{"job0 nextJobs", []string{"~{{ .PipelineName }}.01.stage0.job1"}, jp.resources[0].(jenkinsSingleJob).NextJobs},
		{"job0 manualNextJobs", "", jp.resources[0].(jenkinsSingleJob).NextManualJobs},
		{"job0 workingDir", "subdir/.*", jp.resources[0].(jenkinsSingleJob).WorkingDir},

		{"job1 name", "job1", jp.resources[1].(jenkinsSingleJob).TaskName},
		{"job1 IsInitialJob", false, jp.resources[1].(jenkinsSingleJob).IsInitialJob},
		{"job1 projectNameFmt", "~{{ .PipelineName }}.01.stage0.job1", jp.resources[1].(jenkinsSingleJob).ProjectNameTempl},
		{"job1 git url", "http://github.com/kesselborn/tuev", jp.resources[1].(jenkinsSingleJob).GitURL},
		{"job1 jenkins projectName", "~test-project.01.stage0.job1", job1EndPoint},

		{"job2 next job", []string{"~{{ .PipelineName }}.03.stage1.multi__job4_job5"}, jp.resources[2].(jenkinsSingleJob).NextJobs},

		{"job3 projectNameFmt", "~{{ .PipelineName }}.03.stage1.multi__job4_job5", jp.resources[3].(jenkinsMultiJob).ProjectNameTempl},
		{"job3 nextJobs", []string{"~{{ .PipelineName }}.06.stage2.multi__job7_job8"}, jp.resources[3].(jenkinsMultiJob).NextJobs},

		{"job4 projectName", "~{{ .PipelineName }}.04.stage1.job4", jp.resources[4].(jenkinsSingleJob).ProjectNameTempl},
		{"job4 name has whitespace prefix", "---- job4", jp.resources[4].(jenkinsSingleJob).TaskName},

		{"job7 does not clean workspace", false, jp.resources[7].(jenkinsSingleJob).CleanWorkspace},
		{"job8 does clean workspace", true, jp.resources[8].(jenkinsSingleJob).CleanWorkspace},

		{"job9 multi job has successor", 1, len(jp.resources[9].(jenkinsMultiJob).NextJobs)},

		{"job12 nextManualJobs", "~{{ .PipelineName }}.13.stage4.job13,~{{ .PipelineName }}.14.stage5.job14,~{{ .PipelineName }}.15.stage5.job15", jp.resources[12].(jenkinsSingleJob).NextManualJobs},

		{"job12 stage name", "|>| stage4", jp.resources[13].(jenkinsSingleJob).StageName},
		{"job12 does not have next job as it is in manual stage", 0, len(jp.resources[13].(jenkinsSingleJob).NextJobs)},

		{"pipeline projectName", "test-project", pipelineViewEndPoint},
	}.Run(t)
}

func TestPipelineArtifactLogic(t *testing.T) {
	jp, _ := jenkinsConfigFromFile("tests-fixtures/test_config.json")

	expectations{
		{"job0 artifact", "spec2/", jp.resources[0].(jenkinsSingleJob).Artifact},
		{"job1 artifact dependency", []artifactDep{{"{{ .PipelineName }}", "spec2/"}}, jp.resources[1].(jenkinsSingleJob).ArtifactDep},

		{"job2 artifact", "spec1/,spec2/", jp.resources[2].(jenkinsSingleJob).Artifact},
		{"job4 artifact dependency", []artifactDep{{"~{{ .PipelineName }}.02.stage1.job2", "spec1/,spec2/"}}, jp.resources[4].(jenkinsSingleJob).ArtifactDep},
		{"job5 artifact dependency", []artifactDep{{"~{{ .PipelineName }}.02.stage1.job2", "spec1/,spec2/"}}, jp.resources[5].(jenkinsSingleJob).ArtifactDep},
		{"job7 artifact dependency", []artifactDep{{"~{{ .PipelineName }}.04.stage1.job4", "spec3/"}, {"~{{ .PipelineName }}.05.stage1.job5", "spec4/"}}, jp.resources[7].(jenkinsSingleJob).ArtifactDep},
		{"job8 artifact dependency", []artifactDep{{"~{{ .PipelineName }}.04.stage1.job4", "spec3/"}, {"~{{ .PipelineName }}.05.stage1.job5", "spec4/"}}, jp.resources[8].(jenkinsSingleJob).ArtifactDep},

		{"job12 artifact dependency", []artifactDep{{"~{{ .PipelineName }}.10.stage3.job10", "art1"}, {"~{{ .PipelineName }}.11.stage3.job11", "art2"}}, jp.resources[12].(jenkinsSingleJob).ArtifactDep},
	}.Run(t)
}

func TestRendering(t *testing.T) {
	jp, _ := jenkinsConfigFromFile("tests-fixtures/test_config.json")

	multiJob, multiJobErr := jp.resources[9].renderResource("multi-job-test")
	singleInitialJob, singleInitialJobErr := jp.resources[0].renderResource("initial-job-with-artifact")
	singleJob, singleJobErr := jp.resources[1].renderResource("normal-job-with-artifact-dep")
	multiArtifDeps, multiArtifDepsErr := jp.resources[7].renderResource("multiple-artifact-deps")
	pipelineView, pipelineViewError := jp.resources[15].renderResource("pipeline-view")

	expectations{
		{"rendering multi job should not throw error", nil, multiJobErr},
		{"rendering normal job should not throw error", nil, singleJobErr},
		{"rendering normal initial job should not throw error", nil, singleInitialJobErr},
		{"rendering normal job with multiple artifact deps correctly", nil, multiArtifDepsErr},
		{"rendering pipeline view shoud not throw error", nil, pipelineViewError},
	}.Run(t)

	if testing.Verbose() {
		fmt.Fprintf(os.Stderr, "Verbose mode: storing rendering results in temporary files:\n")
		for name, resource := range map[string]io.Reader{"multijob": multiJob, "singleInitialJob": singleInitialJob, "singleJob": singleJob, "pipelineView": pipelineView, "multipleArtifDeps": multiArtifDeps} {
			f, err := ioutil.TempFile("", "__"+name+".xml__")
			if err != nil {
				t.Errorf(err.Error())
				continue
			}
			fmt.Fprintf(os.Stderr, "\t%s\n", f.Name())
			fmt.Fprintf(f, "%s", resource)
		}
	}
}

func TestJenkinsJobsParsing(t *testing.T) {
	f := configFileDescriptor("tests-fixtures/jenkins_jobs.json")
	var jobList jenkinsJobList
	err := json.NewDecoder(f).Decode(&jobList)

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	expectations{
		{"reading jenkins job list should work", nil, err},
		{"reading jenkins job list should work (msg)", "", errMsg},
	}.Run(t)
}

func TestJenkinsPluginsParsing(t *testing.T) {
	f := configFileDescriptor("tests-fixtures/jenkins_plugins.json")
	var plugins jenkinsPlugins

	err := json.NewDecoder(f).Decode(&plugins)

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	validCheck := plugins.installed([]string{"matrix-auth", "git", "translation"})
	invalidCheck := plugins.installed([]string{"invalid", "plugin", "list"})

	expectations{
		{"reading jenkins plugins list should work", nil, err},
		{"reading jenkins plugins list should work (msg)", "", errMsg},
		{"parses all plugins", 32, len(plugins.List)},
		{"plugin 1 name", "matrix-auth", plugins.List[0].ShortName},
		{"plugin 31 version", "1.32.1", plugins.List[31].Version},
		{"check should work with valid plugins list", nil, validCheck},
		{"check should fail if plugins are missing", true, invalidCheck != nil},
		{"check should fail if plugins are missing (msg)", "jenkins server not setup correctly: required plugin 'invalid' not installed", invalidCheck.Error()},
	}.Run(t)
}
