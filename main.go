package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"

	messagediff "gopkg.in/d4l3k/messagediff.v1"
)

func main() {
	pName1 := flag.String("p1", "", "Cache Parameter Group Name")
	pName2 := flag.String("p2", "", "Cache Parameter Group Name")

	flag.Parse()

	if *pName1 == "" || *pName2 == "" {
		fmt.Fprintln(os.Stderr, "Cache Parameter Group Name is missing.")
		os.Exit(1)
	}

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := elasticache.New(session)

	pInput1 := &elasticache.DescribeCacheParametersInput{
		CacheParameterGroupName: aws.String(*pName1),
	}

	pOutput1, err := client.DescribeCacheParameters(pInput1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	pInput2 := &elasticache.DescribeCacheParametersInput{
		CacheParameterGroupName: aws.String(*pName2),
	}

	pOutput2, err := client.DescribeCacheParameters(pInput2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	diff, _ := messagediff.PrettyDiff(pOutput1, pOutput2)

	var diffIds []int

	for _, d := range strings.Split(diff, "\n") {
		if strings.Contains(d, "modified") && !strings.Contains(d, ".Source") {
			id, _ := strconv.Atoi(d[22:24])
			diffIds = append(diffIds, id)
		}
	}

	fmt.Printf("===== Cache Parameter Group Name : %s =====\n\n", *pName1)

	for _, diffID := range diffIds {
		fmt.Println(pOutput1.Parameters[diffID])
	}

	fmt.Println("")
	fmt.Printf("===== Cache Parameter Group Name : %s =====\n\n", *pName2)

	for _, diffID := range diffIds {
		fmt.Println(pOutput2.Parameters[diffID])
	}
}
