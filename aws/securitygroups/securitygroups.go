package securitygroups

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"awsplus/config"
)

var region string
var sess *session.Session
var sessionErr error

func init() {

	cfg := config.Config
	region = cfg.Region

	sess, sessionErr = session.NewSession(&aws.Config{Region: aws.String(region)})

	if sessionErr != nil {
		log.Fatalln(sessionErr.Error())
	}

}

// list all security groups
// if vpcid passed, list all security groups for that vpc
func List(vpcid string) []*ec2.SecurityGroup {

	svc := ec2.New(sess)
	params := &ec2.DescribeSecurityGroupsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(vpcid),
				},
			},
		},
	}

	if len(vpcid)==0 {
		params = nil
	}

	resp, err := svc.DescribeSecurityGroups(params)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return resp.SecurityGroups
}

func Get(id string) *ec2.SecurityGroup {

	svc := ec2.New(sess)
	params := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{
			aws.String(id),
		},
	}
	resp, err := svc.DescribeSecurityGroups(params)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return resp.SecurityGroups[0]
}

func GetName(id string) *string {

	svc := ec2.New(sess)
	params := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{
			aws.String(id),
		},
	}
	resp, err := svc.DescribeSecurityGroups(params)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return resp.SecurityGroups[0].GroupName
}