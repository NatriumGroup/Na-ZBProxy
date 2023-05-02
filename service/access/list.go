package access

import (
	"fmt"
	//"net/http"
	//"io/ioutil"
	"github.com/layou233/ZBProxy/common/set"
	"github.com/layou233/ZBProxy/config"
)

type Users struct {
	Code string `json:"code"`
	Data string `json:"data"`
}
func GetTargetList(listName string) (set.StringSet, error) {
	//set, ok := config.Config.Lists[listName]
	/*res, err := http.Get("http://124.223.44.130/api/v2/ZBP.php?ign="+listName)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//bodystr := string(body)
	
	if string(body)==string(`"200"`) {
		return listName, nil
	}
	return nil, fmt.Errorf("list %q not found", listName)
	*/
	set, ok := config.Config.Lists[listName]
	if ok {
		return set, nil
	}
	return nil, fmt.Errorf("list %q not found", listName)
	
}
