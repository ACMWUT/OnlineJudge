package config

type LangConfig struct {
	Lang         string
	BuildSh      string
	RunnerConfig string
}

func GetJudgeConfig() map[string]interface{} {
	judgeConfig := make(map[string]interface{})

	judgeConfig["env"] = "dev"
	judgeConfig["address"] = "127.0.0.1:8800"
	judgeConfig["base_dir"] = "/home/acmwhut/data"
	judgeConfig["tmp_dir"] = "/home/ana_tmpdir"

	return judgeConfig
}

func GetLangConfigs() []LangConfig {
	langBasePath := "/home/env"

	langBuildPath := []string {
		"",
		"",
		"/java.openjdk10/build.sh",
		"/python.cpython3.6/build.sh",
	}

	langRunnerConfig := []string {
		"/c.gcc/runner.toml",
		"/cpp.g++/runner.toml",
		"/java.openjdk10/runner.toml",
		"/python.cpython3.6/runner.toml",
	}
	
	langConfigs := []LangConfig{
		{"c.gcc", "", langBasePath + langRunnerConfig[0]},
		{"cpp.g++", "", langBasePath + langRunnerConfig[1]},
		{"java.openjdk10", langBasePath + langBuildPath[2], langBasePath + langRunnerConfig[2]},
		{"python.cpython3.6", langBasePath + langBuildPath[3], langBasePath + langRunnerConfig[3]},
	}
	return langConfigs
}
