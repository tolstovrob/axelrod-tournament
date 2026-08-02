package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

func main() {
	exp := flag.String("e", "all", "experiment: all|0|1|2|2s|3|4|5|7")
	rounds := flag.Int("rounds", 200, "rounds per match")
	defectors := flag.Int("defectors", 20, "number of defectors (exp 1, 2)")
	cluster := flag.Int("cluster", 5, "cluster size K (exp 2)")
	flag.Parse()

	cfg := internal.Config{
		Rounds:       *rounds,
		PayoffMatrix: internal.DefaultPayoffMatrix(),
	}

	switch *exp {
	case "all":
		runAll(cfg)
	case "0":
		exp0Classic(cfg)
	case "1":
		exp1LoneNice(cfg, *defectors)
	case "2":
		exp2Cluster(cfg, *cluster, *defectors)
	case "2s":
		exp2ClusterSweep(cfg, 12, *defectors)
	case "3":
		exp3Envy(cfg)
	case "4":
		exp4Clarity(cfg)
	case "5":
		exp5Noise(cfg, []float64{0.0, 0.01, 0.05, 0.10})
	case "7":
		exp7Shadow(cfg, []int{5, 10, 20, 50, 100, 200})
	default:
		fmt.Fprintf(os.Stderr, "unknown experiment: %s\n", *exp)
		fmt.Fprintf(os.Stderr, "usage: experiments -e all|0|1|2|2s|3|4|5|7\n")
		os.Exit(1)
	}
}
