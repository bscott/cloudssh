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
			resp, err := e.DescribeInstances(nil, filter)
			if err != nil {
				panic(err)
			}
			for _, instance := range resp.Reservations {
				for _, reservation := range instance.Instances {
					fmt.Printf("Instance ID: %s\n", reservation.InstanceId)
					fmt.Printf("IP Address: %s\n", reservation.IPAddress)
					fmt.Printf("DNS Name: %s\n", reservation.DNSName)
					fmt.Printf("State: %v\n", reservation.State.Name)
					fmt.Printf("Key Pair: %s\n", reservation.KeyName)
					fmt.Printf("Tag Name: %s\n", reservation.Tags)
				}
			}
			//fmt.Printf("%#v", resp)
		},
	}

	var rootCmd = &cobra.Command{Use: "cloudssh"}
	rootCmd.AddCommand(cloudsshCmd)
	rootCmd.Execute()
}
