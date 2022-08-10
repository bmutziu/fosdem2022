package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/gocode/gocodec"
)

type ab struct {
	A int
	B int
}

var cnstrnt = ">0 & <4"

func main() {
	//abba := &ab{
	//	A: 100,
	//	B: 200,
	//}

	ctx := cuecontext.New()
	i := ctx.Encode(cnstrnt)
	fmt.Println(i)

	g := gocodec.New((*cue.Runtime)(ctx), nil)
	v, err := g.Decode(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("v: %v\n", v)

	// errbis := g.Validate(i, []byte(`{A:100,B:200}`))
	// errbis := g.Validate(i, ab{100, 200})
	errbis := g.Validate(i, ">0 & <4")

	if errbis != nil {
		panic(errbis)
	}

	instances := load.Instances([]string{}, &load.Config{Dir: "."})
	for _, i := range instances {
		ctx := cuecontext.New()
		v := ctx.BuildInstance(i)
		if v.Err() != nil {
			log.Fatalf("%s", v.Err())
		}
		fmt.Printf("%s\n", v)
	}
}
