package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig_Default(t *testing.T) {
	config := InitConfig()
	assert.Equal(t, 8080, config.Port)
}

func TestInitConfig_WithEnv(t *testing.T) {
	err := os.Setenv(EnvPrefix+"_PORT", "3000")
	assert.NoError(t, err, "unable to setEnv")
	err = os.Setenv(EnvPrefix+"_MONITORABLE_JENKINS_URL", "test")
	assert.NoError(t, err, "unable to setEnv")
	err = os.Setenv(EnvPrefix+"_MONITORABLE_AZUREDEVOPS_URL", "test")
	assert.NoError(t, err, "unable to setEnv")

	config := InitConfig()

	assert.Equal(t, "production", config.Env)
	assert.Equal(t, 3000, config.Port)
}

////nolint:dupl
//func TestPingdom_IsValid(t *testing.T) {
//	config := InitConfig()
//
//	config.Monitorable.Pingdom[DefaultVariant].URL = ""
//	config.Monitorable.Pingdom[DefaultVariant].Token = ""
//	assert.False(t, config.Monitorable.Pingdom[DefaultVariant].IsValid())
//
//	config.Monitorable.Pingdom[DefaultVariant].URL = ""
//	config.Monitorable.Pingdom[DefaultVariant].Token = "abcde"
//	assert.True(t, config.Monitorable.Pingdom[DefaultVariant].IsValid())
//
//	config.Monitorable.Pingdom[DefaultVariant].URL = "url%url"
//	config.Monitorable.Pingdom[DefaultVariant].Token = "abcde"
//	assert.False(t, config.Monitorable.Pingdom[DefaultVariant].IsValid())
//
//	config.Monitorable.Pingdom[DefaultVariant].URL = "http://pingdom.com/api"
//	config.Monitorable.Pingdom[DefaultVariant].Token = "abcde"
//	assert.True(t, config.Monitorable.Pingdom[DefaultVariant].IsValid())
//}
//
//func TestTravisCI_IsValid(t *testing.T) {
//	config := InitConfig()
//	assert.True(t, config.Monitorable.TravisCI[DefaultVariant].IsValid())
//
//	config.Monitorable.TravisCI[DefaultVariant].URL = "url%url"
//	assert.False(t, config.Monitorable.TravisCI[DefaultVariant].IsValid())
//
//	config.Monitorable.TravisCI[DefaultVariant].URL = ""
//	assert.False(t, config.Monitorable.TravisCI[DefaultVariant].IsValid())
//}
//
//func TestJenkins_IsValid(t *testing.T) {
//	config := InitConfig()
//
//	config.Monitorable.Jenkins[DefaultVariant].URL = ""
//	assert.False(t, config.Monitorable.Jenkins[DefaultVariant].IsValid())
//
//	config.Monitorable.Jenkins[DefaultVariant].URL = "url%url"
//	assert.False(t, config.Monitorable.Jenkins[DefaultVariant].IsValid())
//
//	config.Monitorable.Jenkins[DefaultVariant].URL = "http://jenkins.example.com"
//	assert.True(t, config.Monitorable.Jenkins[DefaultVariant].IsValid())
//}
//
////nolint:dupl
//func TestAzureDevOps_IsValid(t *testing.T) {
//	config := InitConfig()
//
//	config.Monitorable.AzureDevOps[DefaultVariant].URL = ""
//	config.Monitorable.AzureDevOps[DefaultVariant].Token = ""
//	assert.False(t, config.Monitorable.AzureDevOps[DefaultVariant].IsValid())
//
//	config.Monitorable.AzureDevOps[DefaultVariant].URL = ""
//	config.Monitorable.AzureDevOps[DefaultVariant].Token = "abcde"
//	assert.False(t, config.Monitorable.AzureDevOps[DefaultVariant].IsValid())
//
//	config.Monitorable.AzureDevOps[DefaultVariant].URL = "url%url"
//	config.Monitorable.AzureDevOps[DefaultVariant].Token = "abcde"
//	assert.False(t, config.Monitorable.AzureDevOps[DefaultVariant].IsValid())
//
//	config.Monitorable.AzureDevOps[DefaultVariant].URL = "http://pingdom.com/api"
//	config.Monitorable.AzureDevOps[DefaultVariant].Token = "abcde"
//	assert.True(t, config.Monitorable.AzureDevOps[DefaultVariant].IsValid())
//}
//
//func TestGithub_IsValid(t *testing.T) {
//	config := InitConfig()
//
//	config.Monitorable.Github[DefaultVariant].Token = ""
//	assert.False(t, config.Monitorable.Github[DefaultVariant].IsValid())
//
//	config.Monitorable.Github[DefaultVariant].Token = "abcde"
//	assert.True(t, config.Monitorable.Github[DefaultVariant].IsValid())
//}
