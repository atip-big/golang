package main
 
import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// set proxy
	//os.Setenv("HTTP_PROXY", "10.138.250.10:9119")

	http.HandleFunc("/fb", facebookHandler)
	
	http.ListenAndServe(":8080", nil)
}

func facebookHandler(writer http.ResponseWriter, _ *http.Request) {
	token := "[add token here]"
	resp, _ := http.Get("https://graph.facebook.com/v1.0/me?access_token="+token)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	fmt.Fprintf(writer, string(body))
}
