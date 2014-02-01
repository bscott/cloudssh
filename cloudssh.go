package main

import (
	"fmt"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/ec2"
	"github.com/spf13/cobra"
)

func main() {

	auth, err := aws.EnvAuth()
	if err != nil {
		fmt.Println(err)
		// panic(err.String())
	}

	e := ec2.New(auth, aws.USEast)

	var cloudsshCmd = &cobra.Command{
		Use:   "ec2",
		Short: "cloudssh lists cloud instances and allows you to ssh the target node",
		Long:  "cloudssh lists cloud instances and allows you to ssh the target node",
		Run: func(c *cobra.Command, arg []string) {
			fmt.Println("Listing EC2 Instances...")
			//var instIds []string
			filter := ec2.NewFilter()
			instances, err := e.DescribeInstances(nil, filter)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#v", instances)
		},
	}

	var rootCmd = &cobra.Command{Use: "cloudssh"}
	rootCmd.AddCommand(cloudsshCmd)
	rootCmd.Execute()
}
