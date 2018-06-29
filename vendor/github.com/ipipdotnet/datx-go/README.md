# datx-go
IPIP.net官方支持datx格式的解析代码

## 安装 Installation
<pre><code>go get github.com/ipipdotnet/datx-go</code></pre>

## 示例 Usage
<pre>
<code>
package main

import "github.com/ipipdotnet/datx-go"
import "fmt"

func main(){

    city, err := datx.NewCity("/path/to/17monipdb.datx") // For City Level IP Database
    if err == nil {
        fmt.Println(city.Find("8.8.8.8"))
        fmt.Println(city.Find("128.8.8.8"))
        fmt.Println(city.Find("255.255.255.255"))
        loc, err := city.FindLocation("27.190.252.103")
        if err == nil {
            fmt.Println(string(loc.ToJSON()))
            // Output:
            /*
            {
                "Country": "China",
                "Province": "Hebei",
                "City": "Tangshan",
                "Organization": "",
                "ISP": "ChinaTelecom",
                "Latitude": "39.635113",
                "Longitude": "118.175393",
                "TimeZone": "Asia/Shanghai",
                "TimeZone2": "UTC+8",
                "CityCode": "130200",
                "PhonePrefix": "86",
                "CountryCode": "CN",
                "ContinentCode": "AP",
                "IDC": "",
                "BaseStation": "",
                "Anycast": false
            }
            */
        }
    }
    dis, err := datx.NewDistrict("/path/to/quxian.datx") // Only China District IP Database
    if err == nil {
        fmt.Println(dis.Find("1.12.46.0"))
        fmt.Println(dis.Find("223.255.127.0"))
    }
    bst, err := datx.NewBaseStation("/path/to/station_ip.datx") // Only China Base Station IP Database
    if err == nil {
        fmt.Println(bst.Find("1.30.6.0"))
        fmt.Println(bst.Find("223.221.121.0"))
        fmt.Println(bst.Find("223.221.121.255"))
    }    
}
</code>
</pre>
