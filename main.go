package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Rules is the data structure for storing iptables rules
type Rules struct {
	Rules []string
}

var rules Rules

var port, timeInterval int

var (
	rulesInactive = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "inactive_rules",
			Help: "Number of inactive rules in iptables",
		})
)

/*
Init function parse the commandline arguments and load the file with
iptables rules to check.
*/
func init() {
	rulesFile := flag.String(
		"r",
		"/usr/share/iptables-exporter/rules.json",
		"Location of the file containing rules to check",
	)

	flag.IntVar(&port, "p", 9455, "Listening port")

	flag.IntVar(&timeInterval, "i", 60, "Time period in seconds between metric collector run")

	flag.Parse()

	data, err := ioutil.ReadFile(*rulesFile)
	if err != nil {
		fmt.Print(err)
		fmt.Println("Error with reading the configuration file. Exiting.")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &rules)
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("Error with loading the configuration from file. Exiting.")
		os.Exit(2)
	}

	prometheus.MustRegister(rulesInactive)

	fmt.Println("Starting iptables exporter")
	fmt.Printf("Listening on port: %v\n", port)
	fmt.Printf("Metric collection interval: %v seconds\n", timeInterval)
	fmt.Println("Checking for following iptables rules:")
	for _, rule := range rules.Rules {
		fmt.Printf("\t%v\n", rule)
	}
	fmt.Println("Iptables exporter started")
}

/*
RecordMetrics function is a go routine to periodically check if the given
list of rules exist and update the metric value in the metric endpoint
*/
func recordMetrics() {
	go func() {
		for {
			var inactiveRules []string
			for _, rule := range rules.Rules {
				splitedRule := strings.Split(rule, " ")
				args := []string{"-C"}
				args = append(args, splitedRule...)

				err := exec.Command("/usr/sbin/iptables", args...).Run()
				if err != nil {
					fmt.Printf("Rule \"%v\" doesn't exist\n", rule)
					inactiveRules = append(inactiveRules, rule)
				}
			}
			rulesInactive.Set(float64(len(inactiveRules)))
			fmt.Println("Metric updated")
			time.Sleep(time.Duration(timeInterval) * time.Second)
		}
	}()
}

func main() {
	prometheus.Unregister(prometheus.NewGoCollector())
	recordMetrics()
	port := fmt.Sprintf(":%d", port)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(port, nil)
}
