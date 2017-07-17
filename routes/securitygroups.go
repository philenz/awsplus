package routes

/*
 * Run /securitygroups to get a JSON array of Security Group Information
 */

import (
	"encoding/json"
	"net/http"
	"awsplus/aws/securitygroups"
	"awsplus/aws/vpcs"
)

type Ingress struct {
	Protocol string
	FromPort int64
	ToPort   int64
	Source   []string
}

type SecurityGroup struct {
	Id          string
	Name        string
	Vpc         string
	Description string
	Ingress     []Ingress
}

// lookups for vpc name and group name
var vmap = make(map[string]string)
var gmap = make(map[string]string)

func init() {

	// lookup vpc names
	for _, v := range vpcs.List() {
		vmap[*v.VpcId] = vpcs.GetName(*v.VpcId)
	}

	// lookup security group names
	for _, g := range securitygroups.List("") {
		gmap[*g.GroupId] = *securitygroups.GetName(*g.GroupId)
	}

}

func ListSecurityGroups(writer http.ResponseWriter, request *http.Request) {

	var sgs = make([]SecurityGroup, len(securitygroups.List("")))

	for sgcount, g := range securitygroups.List("") {

		if len(g.IpPermissions) > 0 {

			ips := make([]Ingress, len(g.IpPermissions))

			for ipcount, ip := range g.IpPermissions {

				if *ip.IpProtocol != "-1" {

					sources := make([]string, len(ip.IpRanges)+len(ip.UserIdGroupPairs))

					sourcecount := 0
					for _, source := range ip.IpRanges {
						sources[sourcecount] = *source.CidrIp
						sourcecount++
					}
					for _, source := range ip.UserIdGroupPairs {
						sources[sourcecount] = gmap[*source.GroupId]
						sourcecount++
					}

					ips[ipcount] = Ingress{*ip.IpProtocol, *ip.FromPort, *ip.ToPort, sources}

				} else { // default security groups... don't really care about these... maybe filter out in securitygroups.List ???
					ips[ipcount] = Ingress{*ip.IpProtocol, 0, 65535, nil}

				}
			}

			sgs[sgcount] = SecurityGroup{*g.GroupId, *g.GroupName, vmap[*g.VpcId], *g.Description, ips}

		} else { // security group with no ingress rules...
			sgs[sgcount] = SecurityGroup{*g.GroupId, *g.GroupName, vmap[*g.VpcId], *g.Description, nil}
		}

	}

	json.NewEncoder(writer).Encode(sgs)

}
