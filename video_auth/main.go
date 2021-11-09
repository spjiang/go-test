package main

import (
	"fmt"
	"github.com/spjiang/go-test/video_auth/util"
)

func main() {

	c := &util.CreatorAuth{
		Body: "camera/8.flv",
		Sign: "",
	}
	_, err := c.CreateExpireSign()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Sign)
	b := &util.CreatorAuth{
		Body: "camera/8.flv",
		Sign: "310c502da0c50ae0654ab6f6baf1e7721696a987e44131ec29d6f00c4776af581619489147",
	}
	vr, err := b.VerifyExpireSign()
	fmt.Println(vr)
	fmt.Println(err)

}
