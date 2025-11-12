package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Endpoint     string `yaml:"endpoint"`
	SecretID     string `yaml:"secretID"`
	SecretKey    string `yaml:"secretKey"`
	ContactEmail string `yaml:"contactEmail"`
	ContactPhone string `yaml:"contactPhone"`
}

var (
	commands map[string]func(args []string) error
	config   Config
	api      *Api
)

func init() {
	initConfig()
	initCommands()

	var err error
	api, err = New(config.Endpoint, config.SecretID, config.SecretKey)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func initConfig() {
	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func initCommands() {
	commands = map[string]func(args []string) error{
		"help":     func(args []string) error { HelpInfo(); return nil },
		"apply":    ApplyCertificate,
		"list":     ListCertificates,
		"detail":   DescribeCertificate,
		"keys":     CertificateKeyPair,
		"download": DownloadCertificate,
		"url":      DescribeDownloadCertificateUrl,
		"revoke":   RevokeCertificate,
		"delete":   DeleteCertificate,
	}
}

func main() {
	parts := os.Args

	if len(parts) == 1 {
		HelpInfo()
		return
	}

	cmd := parts[1]
	args := parts[2:]

	if f, ok := commands[cmd]; ok {
		err := f(args)
		if err != nil {
			fmt.Printf("发生错误：%+v", err)
		}
	} else {
		fmt.Printf("Unknown command: %s", cmd)
	}

	fmt.Print("\n")
}
