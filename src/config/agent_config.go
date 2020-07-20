package config

import (
	"flag"
	"os"

	"github.com/ori-amizur/introspector/generated/bm-inventory/client"
)

var GlobalAgentConfig struct {
	IsText             bool
	TargetHost         string
	TargetPort         int
	ClusterID          string
	IntervalSecs       int
	ConnectivityParams string
	InventoryImage     string
	JournalLogging     bool
	TextLogging        bool
	AgentVersion       string
}

func printHelpAndExit() {
	flag.CommandLine.Usage()
	os.Exit(0)
}

func ProcessArgs() {
	ret := &GlobalAgentConfig
	flag.BoolVar(&ret.IsText, "text", false, "Output only as text")
	flag.StringVar(&ret.TargetHost, "host", client.DefaultHost, "The target host")
	flag.IntVar(&ret.TargetPort, "port", 80, "The target port")
	flag.StringVar(&ret.ClusterID, "cluster-id", "default-cluster", "The value of the cluster-id")
	flag.StringVar(&ret.AgentVersion, "agent-version", "", "Discovery agent version")
	flag.IntVar(&ret.IntervalSecs, "interval", 60, "Interval between steps polling in seconds")
	flag.StringVar(&ret.ConnectivityParams, "connectivity", "", "Test connectivity as output string")
	flag.StringVar(&ret.InventoryImage, "inventory-image", "quay.io/ocpmetal/inventory:latest", "The image of inventory")
	flag.BoolVar(&ret.JournalLogging, "with-journal-logging", true, "Use journal logging")
	flag.BoolVar(&ret.TextLogging, "with-text-logging", true, "Output log to file")
	h := flag.Bool("help", false, "Help message")
	flag.Parse()
	if h != nil && *h {
		printHelpAndExit()
	}
}
