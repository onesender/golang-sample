package main

import (
	"app/onesender"
	"fmt"
)

func main() {
	fmt.Println("Test kirim pesan ke WhatsApp")

	onesender.ApiUrl = "https://onesender.my.id/api/v1/messages"
	onesender.ApiKey = "YOUR_ONESENDER_API_KEY"

	res, err := onesender.SendTextMessage("6281227494445", "Hello ini pesan text")
	if err != nil {
		fmt.Println("Error send text:", err)
	}
	fmt.Println(res)

	res, err = onesender.SendImageMessage("6281227494445", "https://media.geeksforgeeks.org/wp-content/uploads/20200319202059/remove-example-1.jpg", "Hello ini pesan text")
	if err != nil {
		fmt.Println("Error send image:", err)
	}
	fmt.Println(res)

	res, err = onesender.SendDocumentMessage("6281227494445", "https://media.geeksforgeeks.org/wp-content/uploads/20200319202059/remove-example-1.jpg")
	if err != nil {
		fmt.Println("Error send document:", err)
	}
	fmt.Println(res)
}
