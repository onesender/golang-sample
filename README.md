# golang-sample
Contoh kirim pesan ke WhatsApp menggunakan Golang dan OneSender


```
package main

import (
	"app/onesender"
	"fmt"
)

func main() {
	fmt.Println("Test kirim pesan ke WhatsApp")

    // Masukkan alamat api dan key
	onesender.ApiUrl = "https://onesender.my.id/api/v1/messages"
	onesender.ApiKey = "YOUR_ONESENDER_API_KEY"

    // Kirim pesan text
	res, err := onesender.SendTextMessage("6281227494445", "Hello ini pesan text")
	if err != nil {
		fmt.Println("Error send text:", err)
	}
	fmt.Println(res)


    // Kirim pesan gambar
	res, err = onesender.SendImageMessage("6281227494445", "https://media.geeksforgeeks.org/wp-content/uploads/20200319202059/remove-example-1.jpg", "Hello ini pesan text")
	if err != nil {
		fmt.Println("Error send image:", err)
	}
	fmt.Println(res)

    // Kirim pesan dokumen
	res, err = onesender.SendDocumentMessage("6281227494445", "https://media.geeksforgeeks.org/wp-content/uploads/20200319202059/remove-example-1.jpg")
	if err != nil {
		fmt.Println("Error send document:", err)
	}
	fmt.Println(res)
}

```