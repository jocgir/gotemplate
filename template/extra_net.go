package template

import (
	"math"
	"net"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/coveooss/gotemplate/v3/collections"
)

const (
	netBase = "Net"
)

var netFuncs = dictionary{
	"httpGet":   httpGet,
	"httpDoc":   httpDocument,
	"cidr":      func(network interface{}) *net.IPNet { return asCIDR(network) },
	"maskBits":  func(mask interface{}) int { return maskSize(mask, true) },
	"maskSize":  func(mask interface{}) int { return maskSize(mask, false) },
	"cidrMask":  func(network interface{}) net.IPMask { return asCIDR(network).Mask },
	"cidrCount": func(network interface{}) uint64 { return cidr.AddressCount(asCIDR(network)) },
	"cidrRange": func(network interface{}) iList {
		result := collections.CreateList(2)
		r := result.AsArray()
		r[0], r[1] = cidr.AddressRange(asCIDR(network))
		return result
	},
	"cidrHost": func(network interface{}, num interface{}) (net.IP, error) {
		return cidr.Host(asCIDR(network), toInt(num))
	},
	// "cidrMerge": func() {},
	"cidrSubnet": func(network, newBits, num interface{}) (*net.IPNet, error) {
		return cidr.Subnet(asCIDR(network), toInt(newBits), toInt(num))
	},
	"cidrSubnets": func(network interface{}, divider interface{}) (iList, error) {
		main := asCIDR(network)
		quantity := int(toUnsignedInteger(divider))
		result := collections.CreateList(0, int(quantity))
		mask := int(math.Ceil(math.Log2(float64(quantity))))
		for i := 0; i < quantity; i++ {
			if subnet, err := cidr.Subnet(main, mask, i); err == nil {
				result = result.AppendRaw(subnet)
			} else {
				return result, err
			}
		}
		return result, nil
	},
}

var netFuncsArgs = arguments{
	"httpGet": {"url"},
	"httpDoc": {"url"},
}

var netFuncsAliases = aliases{
	"httpDoc":     {"httpDocument", "curl"},
	"maskSize":    {"maskOnes"},
	"cidrCount":   {"cidrSize"},
	"cidrHost":    {"host"},
	"cidrSubnet":  {"subnet"},
	"cidrSubnets": {"subnets"},
}

var netFuncsHelp = descriptions{
	"httpGet": "Returns http get response from supplied URL.",
	"httpDoc": "Returns http document returned by supplied URL.",
}

func (t *Template) addNetFuncs() {
	t.AddFunctions(netFuncs, netBase, FuncOptions{
		FuncHelp:    netFuncsHelp,
		FuncArgs:    netFuncsArgs,
		FuncAliases: netFuncsAliases,
	})
}

func httpGet(url interface{}) (*http.Response, error) {
	return http.Get(toString(url))
}

func httpDocument(url interface{}) (interface{}, error) {
	response, err := httpGet(url)
	if err != nil {
		return response, err
	}
	return goquery.NewDocumentFromResponse(response)
}

func asCIDR(value interface{}) *net.IPNet {
	_, cidr, err := net.ParseCIDR(toString(value))
	must(err)
	return cidr
}

func maskSize(mask interface{}, getBits bool) int {
	switch value := mask.(type) {
	case net.IPMask:
	default:
		mask = asCIDR(value).Mask
	}
	ones, bits := mask.(net.IPMask).Size()
	return iif(getBits, bits, ones).(int)
}
