
/*
The MIT License (MIT)
Copyright (c) 2016 efaysal
*/



package main
import 
(
"fmt"
"bytes"
"encoding/binary"
	)
	

   func main() {
        var word string = "cafés"

 fmt.Println((len(word)))

       
        bt := []byte(word)

      buf := new(bytes.Buffer)
	for _, v := range bt {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%0x",buf.Bytes())
	fmt.Println("")
	
	var wordé string = "cafés"

 fmt.Println((len(wordé)))

       
        bt = []byte(wordé)

      buf = new(bytes.Buffer)
	for _, v := range bt {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%0x",buf.Bytes())
	
}
